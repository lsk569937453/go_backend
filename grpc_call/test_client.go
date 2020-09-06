package grpc_call

import (
	"context"
	"fmt"
	"go_backend/log"
	"google.golang.org/grpc"
	"io"
)

//test call grpc by stream
func CallGrpc() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Error("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := NewMaxSizeClient(conn)

	// 调用gRPC接口
	tr, err := t.Echo(context.Background(), &Empty{})
	for {
		// stream 有一个最重要的方法，就是 Recv()，Recv 的返回值就是 *pb.StringMessage，这里面包含了多个 Ss []*StringSingle
		data, err := tr.Recv()
		if err == io.EOF {
			// Note: If `maxResults` are returned this will never be reached.
			break
		}
		if err != nil {
			fmt.Printf("error %v", err)
			return
		}
		fmt.Printf("%v", data)
	}

}
