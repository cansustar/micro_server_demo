package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	student_service "micro_service/idl/my_proto"
	"testing"
)

func TestServer(t *testing.T) {
	// 连接服务端
	conn, err := grpc.Dial("127.0.0.1:2346", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		t.Fail()
	}
	// 连接成功后，记得关闭
	defer conn.Close()
	// 生成的go文件里有一个NewStudentServiceClient方法，用来创建一个客户端对象
	client := student_service.NewStudentServiceClient(conn)
	// 可以用这个client对象调用服务端的接口
	// context.TODO() 用来创建一个空的context, 在这里的作用是用来传递参数的。为什么是TODO呢？因为这里的参数还没有确定
	resp, err := client.GetStudentInfo(context.TODO(), &student_service.Request{StudentId: "stu1"})
	if err != nil {
		fmt.Printf("call GetStudentInfo failed, err:%v\n", err)
		t.Fail()
	} else {
		fmt.Println(resp.Name)
		fmt.Println(resp.Age)
		fmt.Println(resp.Height)
	}
}
