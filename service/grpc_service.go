package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"io"
	"strings"

	"github.com/jhump/protoreflect/grpcreflect"
	"go_backend/log"
	"google.golang.org/grpc"
	"time"
)

//
//Get all the Service Name
func GrpcGetServiceList(ipAndPortString string) map[string]int {
	refClient, _ := getRefClient(ipAndPortString)
	defer refClient.Reset()
	service, err := refClient.ListServices()
	log.Info("service List is:%v", service)
	if err != nil {
		log.Error("ListServices error%s", err.Error())
	}
	resultMap := make(map[string]int)
	//filter the default service name:grpc.reflection.v1alpha.ServerReflection
	for _, item := range service {
		if !strings.Contains(item, "grpc.reflection") {
			findAllServiceAndMethod(item, refClient, resultMap)
		}
	}
	return resultMap
}

//find all the service and method name
func findAllServiceAndMethod(realServiceName string, refClient *grpcreflect.Client, mapRes map[string]int) {
	log.Info("realServiceName:%s", realServiceName)
	fileDesc, err := refClient.FileContainingSymbol(realServiceName)
	if err != nil {
		log.Error("FileContainingSymbol error:%s", err)
		return
	}

	//get the service name
	serviceList := fileDesc.GetServices()

	for _, item := range serviceList {
		//get all the mthod description
		methodDescriptions := item.GetMethods()
		//package+service
		serviceName := item.GetFullyQualifiedName()
		for _, methodDescItem := range methodDescriptions {
			methodName := methodDescItem.GetName()
			fmt.Println(methodDescItem.GetInputType().String() + "~~~~~~~~~" + methodDescItem.GetOutputType().String())
			key := serviceName + "$$" + methodName
			vaule := 0
			if methodDescItem.IsServerStreaming() {
				vaule = 0
			} else if methodDescItem.IsServerStreaming() && methodDescItem.IsClientStreaming() {
				vaule = 1
			} else if !methodDescItem.IsClientStreaming() {
				vaule = 2
			}

			mapRes[key] = vaule
		}
	}

}

/**
 *
 * @Description invokeGrpc from ipPort,serviceName,methodName,reqBody
 * @Date 8:11 上午 2020/9/9
 **/
func GrpcRemoteInvoke(ipAndPortString string, serviceName string, methodName string, reqBody string) string {
	refClient, ccReflect := getRefClient(ipAndPortString)
	defer ccReflect.Close()
	defer refClient.Reset()
	service, err := refClient.ListServices()
	log.Info("service List is:%v", service)
	if err != nil {
		log.Error("ListServices error%s", err.Error())
	}
	var methodDesc *desc.MethodDescriptor
	//filter the default service name:grpc.reflection.v1alpha.ServerReflection
	for _, item := range service {
		if !strings.Contains(item, "grpc.reflection") {
			methodDesc = findMethodDesc(item, refClient, serviceName, methodName)
			if methodDesc != nil {
				break
			}
		}
	}
	var result string
	if methodDesc != nil {
		result = callGrpc(ccReflect, methodDesc, reqBody)
	} else {
		log.Error("methodDesc is null")

	}
	return result

}

/**
 *
 * @Description  call the grpc_server with three mode
 * @Date 8:31 上午 2020/9/9
 **/
func callGrpc(ccReflect *grpc.ClientConn, methodDesc *desc.MethodDescriptor, reqBody string) string {
	var result string
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	var ext dynamic.ExtensionRegistry
	msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(&ext)
	stub := grpcdynamic.NewStubWithMessageFactory(ccReflect, msgFactory)
	req := msgFactory.NewMessage(methodDesc.GetInputType())
	jsonpb.Unmarshal(bytes.NewReader([]byte(reqBody)), req)

	if methodDesc.IsServerStreaming() && methodDesc.IsClientStreaming() {

	} else if methodDesc.IsServerStreaming() {
		tr, err := stub.InvokeRpcServerStream(ctx, methodDesc, req)
		//	tr, err := stub.InvokeRpcServerStream(ctx, methodDescItem, data)
		if err != nil {
			log.Error("%s", err.Error())
		}
		for {
			var resp proto.Message

			// stream 有一个最重要的方法，就是 Recv()，Recv 的返回值就是 *pb.StringMessage，这里面包含了多个 Ss []*StringSingle
			resp, err = tr.RecvMsg()
			if err == io.EOF {
				log.Info("find eof and exit")
				// Note: If `maxResults` are returned this will never be reached.
				break
			}
			if err != nil {
				fmt.Printf("error %v", err)
				break
			}
			result = resp.String()
			log.Info("callGrpc  serverStream:%s", result)
		}

	} else if !methodDesc.IsServerStreaming() {

		mes, err := stub.InvokeRpc(ctx, methodDesc, req)
		if err != nil {
			log.Error("%s", err.Error())
		}
		result = mes.String()
		log.Info("callGrpc not serverStream:%s", result)
	} else {
		log.Error("callGrpc could not find the call mode")
		return ""
	}
	return result
}

/**
 *
 * @Description  find the methodDescription
 * @Date 8:21 上午 2020/9/9
 **/
func findMethodDesc(realServiceName string, refClient *grpcreflect.Client, dstServiceName string, dstMethodName string) *desc.MethodDescriptor {
	log.Info("realServiceName:%s", realServiceName)
	fileDesc, err := refClient.FileContainingSymbol(realServiceName)
	if err != nil {
		log.Error("FileContainingSymbol error:%s", err)
		return nil
	}

	//get the service name
	serviceList := fileDesc.GetServices()

	for _, item := range serviceList {
		//get all the mthod description
		methodDescriptions := item.GetMethods()
		//package+service
		srcServiceName := item.GetFullyQualifiedName()
		if srcServiceName != dstServiceName {
			continue
		}

		for _, methodDescItem := range methodDescriptions {
			srcMethodName := methodDescItem.GetName()
			if srcMethodName == dstMethodName {
				return methodDescItem
			}
		}
	}
	return nil

}

/**
 *
 * @Description  get grpcReflect.client from ipandport
 * @Date 8:14 上午 2020/9/9
 **/
func getRefClient(ipAndPortString string) (*grpcreflect.Client, *grpc.ClientConn) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	ipAndPort := ipAndPortString
	ccReflect, err := grpc.DialContext(ctx, ipAndPort, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Error("", err.Error())
	}
	//defer ccReflect.Close()
	refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(ccReflect))
	defer refClient.Reset()

	if err != nil {
		log.Error("%s", err.Error())
	}
	return refClient, ccReflect

}

//
//Get all the Service Name
//func GrpcGetMethodList(ipAndPortString string, serviceMethodName string) []string {
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
//	defer cancel()
//	ipAndPort := ipAndPortString
//	ccReflect, err := grpc.DialContext(ctx, ipAndPort, grpc.WithInsecure(), grpc.WithBlock())
//
//	if err != nil {
//		log.Error("", err.Error())
//	}
//	defer ccReflect.Close()
//	refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(ccReflect))
//	defer refClient.Reset()
//
//	if err != nil {
//		log.Error("%s", err.Error())
//	}
//	service, err := refClient.ListServices()
//	log.Info("service List is:%v", service)
//	if err != nil {
//		log.Error("ListServices error%s", err.Error())
//	}
//	fileDesc, err := refClient.FileContainingSymbol(serviceName)
//	if err != nil {
//		log.Error("find fileDesc error%s", err.Error())
//	}
//	fileDesc.getm
//
//	return result
//}
