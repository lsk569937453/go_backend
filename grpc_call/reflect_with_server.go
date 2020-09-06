package grpc_call

import (
	"context"
	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"go_backend/log"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"strings"
	"time"
)

func TestGrpc() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ccReflect, err := grpc.DialContext(ctx, "127.0.0.1:8028", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error("", err.Error())
	}
	defer ccReflect.Close()
	refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(ccReflect))
	defer refClient.Reset()

	desc := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	//requestFunc := func(m proto.Message) error {
	//
	//	reqStats.Sent++
	//	req := input.Data[0]
	//	input.Data = input.Data[1:]
	//	if err := jsonpb.Unmarshal(bytes.NewReader([]byte(req)), m); err != nil {
	//		return status.Errorf(codes.InvalidArgument, err.Error())
	//	}
	//
	//	return nil
	//}
	err = grpcurl.InvokeRPC(ctx, desc, ccReflect, "test.MaxSize.Echo", nil, grpcurl.NewDefaultEventHandler(log.BaseGinLog(), desc, grpcurl.NewTextFormatter(false), false), nil)

	if err != nil {
		log.Error("%s", err.Error())
	}
	service, err := refClient.ListServices()
	log.Info("service List is:%v", service)
	if err != nil {
		log.Error("ListServices error%s", err.Error())
	}
	var realServiceName string
	//filter the reflection service
	for _, item := range service {
		if strings.Contains(item, "reflection") {
			log.Info("service: %s, has been filter", item)
		} else {
			realServiceName = item
		}
	}

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
		for _, methodDescItem := range methodDescriptions {
			methodDescItem.GetName()
			methodDescItem.GetService()
		}
	}
	var res string
	err = ccReflect.Invoke(ctx, "test.MaxSize.Echo", "{}", res)
	if err != nil {
		log.Error("%s", err.Error())
	}

}
