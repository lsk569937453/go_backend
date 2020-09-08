package service

import (
	"context"
	"fmt"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"strings"

	"github.com/jhump/protoreflect/grpcreflect"
	"go_backend/log"
	"google.golang.org/grpc"
	"time"
)

//
//Get all the Service Name
func GrpcGetServiceList(ipAndPortString string) map[string]int {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ipAndPort := ipAndPortString
	ccReflect, err := grpc.DialContext(ctx, ipAndPort, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Error("", err.Error())
	}
	defer ccReflect.Close()
	refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(ccReflect))
	defer refClient.Reset()

	if err != nil {
		log.Error("%s", err.Error())
	}
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
		serviceName := item.GetName()
		for _, methodDescItem := range methodDescriptions {
			methodName := methodDescItem.GetName()
			fmt.Println(methodDescItem.GetInputType().String() + ":::" + methodDescItem.GetOutputType().String())
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
