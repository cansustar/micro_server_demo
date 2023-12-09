package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	student_service "micro_service/idl/my_proto"
	"net"
	"strconv"
)

// 要实现proto文件中的那个接口

type StudentServer struct {
}

// GetStudentInfo
// func (s *StudentServer) GetStudentInfo() {
//
// }
// 使用指针接收者可以修改接收者的值
// proto中的这个接口只有一个参数request， 但实际上转成go文件后会多一个参数 context. request参数就是proto中定义的那个request，不过是转成go文件后中的
// 返回值除了 proto文件中的student外，还有一个err
func (s *StudentServer) GetStudentInfo(ctx context.Context, request *student_service.Request) (*student_service.Student, error) {
	// 作为一个微服务是支持并发的，每来一个请求都会为这个请求单独开辟一个协程，协程的核心就是来运行这个函数
	// 为了防止某个子协程把整个主协程搞挂，需要对错误即使panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("执行接口时出错：%v\n", err)
		}
	}()
	// 正式的函数逻辑部分
	// 从request中取得参数
	studentId := request.StudentId
	if len(studentId) == 0 {
		return nil, errors.New("缺少必要参数")
	}
	student := GetStudentInfo(studentId)
	return &student, nil
}

func GetStudentInfo(studentId string) student_service.Student {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// 创建连接后，接下来的所有操作都会用到一个context

	ctx := context.TODO()
	stu := student_service.Student{}
	for field, value := range client.HGetAll(ctx, studentId).Val() {
		if field == "Name" {
			stu.Name = value
		} else if field == "Age" {
			age, err := strconv.Atoi(value) // Atoi用来将字符串转换为整型
			if err != nil {
				fmt.Println(err)
			} else {
				stu.Age = int32(age)
			}
		} else if field == "Height" {
			height, err := strconv.ParseFloat(value, 10) // ParseFloat用来将字符串转换成float
			if err == nil {
				stu.Height = float32(height)
			}
		}
	}
	return stu
}

func main() {
	// 需要在main中把整个服务启动起来
	// list 是一个监听
	// 为什么这里不需要指定ip?
	// 因为这个服务是在本机上，所以不需要指定ip
	list, err := net.Listen("tcp", ":2346")
	if err != nil {
		panic(err)
	}
	// 把刚刚的实现 注册到server中。
	//  实例化一个server
	server := grpc.NewServer()
	// 在生成的go文件中，找到注册proto中接口的整个注册方法，传入上面的server和 proto中接口的实现的类 （StudentServer结构体实现了proto中的接口）
	student_service.RegisterStudentServiceServer(server, new(StudentServer))
	/*
		new 函数返回的是指向新分配的零值的指针。对于结构体类型，返回的指针指向一个零值结构体，其中所有字段的值都是类型的零值。例如，new(StudentServer) 返回的是 *StudentServer 类型的指针，指向一个 StudentServer 类型的零值。

		StudentServer{} 是结构体类型的字面量表示法，它创建了一个新的结构体实例，并将所有字段初始化为类型的零值。它直接返回一个 StudentServer 类型的值，而不是指针。
	*/
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}
}
