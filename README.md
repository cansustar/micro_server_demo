# micro_server_demo
 proto文件编写完成后，在生成代码前还需要：
 // 安装protoc  https://github.com/protocolbuffers/protobuf/releases 添加到path环境变量
 protoc是Protocol Buffers的编译器，用于将.proto文件编译成可在不同编程语言中使用的代码文件。

 
// 安装gofast  在终端执行go install github.com/gogo/protobuf/protoc-gen-gofast@latest  会安装在GOPATH下
protoc-gen-gogofaster 是 Google 开源的一款 Protocol Buffer 编译器插件，可以生成 Golang 的代码。

// protoc --gofast_out=plugins=grpc:. --proto_path=idl -I=idl/third_proto idl/student_service.proto

这是使用 Protocol Buffers 编译器 protoc 生成 gRPC 代码的命令。下面解释各个参数的意义：

--gofast_out=plugins=grpc:.：表示使用 gofast 插件生成 gRPC 相关的代码。gofast 是 gogo/protobuf 的一个高性能插件，提供了比标准插件更高效的代码生成。

--proto_path=idl：指定搜索 .proto 文件的根目录。在这个例子中，.proto 文件应该位于 idl 目录下。

idl/student_service.proto：指定要编译的 .proto 文件的路径和文件名。在这个例子中，是 idl 目录下的 student_service.proto 文件。

输出目录：由于命令中的 . 表示输出到当前目录，生成的代码将会放在当前目录下。这个输出目录可以根据需要进行调整。
注意，输出目录这里的.表示是当前目录，子目录，以及生成的go文件名由proto文件中的option_go_package 参数决定
option go_package="./idl/my_proto;student_service";
