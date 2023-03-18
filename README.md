# 用Thrift通信的Go项目

从创建一个新的文件夹开始，我们将按照以下步骤为你的Go项目建立一个基本的项目结构：

1. 创建项目文件夹：
```sh
mkdir your_project
cd your_project
```
2. 初始化Go模块：
```sh
go mod init your_project
```

这将创建一个go.mod文件，其中包含你的项目名称和依赖项。

3. 创建项目结构：
创建以下目录和文件结构：

```sh
your_project/
├── cmd/
│   ├── client/
│   │   └── main.go
│   └── server/
│       └── main.go
├── idl/
│   └── example.thrift
├── pkg/
│   ├── thrift/
│   └── consts/
└── service/
    └── user_service.go
```

你可以使用以下命令创建这些目录和文件：

```sh
mkdir -p cmd/client cmd/server idl pkg/thrift pkg/consts service
touch cmd/client/main.go cmd/server/main.go idl/example.thrift service/user_service.go
```

4. 在idl/example.thrift文件中定义你的Thrift接口。

5. 生成Thrift代码并将其放入pkg/thrift目录中：

```sh
thrift --gen go -out pkg/thrift idl/example.thrift
```
6. 编写service/user_service.go以实现Thrift接口。

7. 编写cmd/server/main.go以设置和启动Thrift服务端。

8. 编写cmd/client/main.go以设置和使用Thrift客户端。

现在你的项目已经具有基本的项目结构，并且可以使用Thrift进行通信。你可以根据需要修改这个结构，添加新的服务、数据结构以及业务逻辑。



---

你已经创建了一个使用Thrift进行通信的Go项目。服务端和客户端通过自动生成的Thrift代码进行交互，使用预先定义的IDL文件中的服务和数据结构。在实际项目中，你可以根据需要添加更多的服务和数据结构，以满足不同的业务需求。
