package grpc_call

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"go_backend/log"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"io"
	"strings"
	"time"
)

type RequestSu struct {
	Data string
}

func (tmp *RequestSu) String() string {
	return tmp.Data
}
func (tmp *RequestSu) Reset() {
}
func (tmp *RequestSu) ProtoMessage() {
}

type RequestSupplier func(proto.Message) error

func (tmp RequestSupplier) String() string {
	return tmp.String()
}
func TestGrpc() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ipAndPort := "127.0.0.1:9000"
	//ipAndPort := "127.0.0.1:50051"
	//	ipAndPort := "45.32.63.93:9000"
	ccReflect, err := grpc.DialContext(ctx, ipAndPort, grpc.WithInsecure(), grpc.WithBlock())
	//data := &RequestSu{Data: "{}"}
	//	dataStr := "{}"
	dataStr := "`{\"name\": \"golang world\"}`"
	var ext dynamic.ExtensionRegistry

	msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(&ext)

	stub := grpcdynamic.NewStubWithMessageFactory(ccReflect, msgFactory)

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
	var realServiceName string
	//filter the reflection service
	for _, item := range service {
		if strings.Contains(item, "reflection") {
			log.Info("service: %s, has been filter", item)
		} else {
			if item == "test.Greeter" {
				realServiceName = item
			}
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
			if methodDescItem.GetName() == "SayGirl" {
				if methodDescItem.IsServerStreaming() && methodDescItem.IsClientStreaming() {

				} else if methodDescItem.IsServerStreaming() {
					req := msgFactory.NewMessage(methodDescItem.GetInputType())
					jsonpb.Unmarshal(bytes.NewReader([]byte(dataStr)), req)

					tr, err := stub.InvokeRpcServerStream(ctx, methodDescItem, req)
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
							return
						}
						fmt.Println(resp.String())
					}

				} else {
					req := msgFactory.NewMessage(methodDescItem.GetInputType())
					jsonpb.Unmarshal(bytes.NewReader([]byte(dataStr)), req)

					mes, err := stub.InvokeRpc(ctx, methodDescItem, req)
					if err != nil {
						log.Error("%s", err.Error())
					}
					log.Info("%s", mes.String())
				}

			}

		}
	}
	//var res string
	//err = ccReflect.Invoke(ctx, "test.MaxSize.Echo", "{}", res)
	//if err != nil {
	//	log.Error("%s", err.Error())
	//}

}
