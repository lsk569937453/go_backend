package service

import (
	"context"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"strings"

	"github.com/jhump/protoreflect/grpcreflect"
	"go_backend/log"
	"google.golang.org/grpc"
	"time"
)

func GrpcGetServiceList(ipAndPortString string) []string {
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
	result := make([]string, 0)
	//filter the default service name:grpc.reflection.v1alpha.ServerReflection
	for _, item := range service {
		if !strings.Contains(item, "grpc.reflection") {
			result = append(result, item)
		}

	}
	return result
}
