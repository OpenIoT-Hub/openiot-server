# openiot

Four main modules:

- user
- admin
- device
- communication / announcement (mail -> ?)
- utils
  - weather
  - address/location

# Project Structure

## Gateway 网关

这里是网关的项目结构，其负责通过远程过程调用(Remote Process Call, RPC)唤起对应的微服务。同时，网关也负责与鉴权服务进行交互，验证请求可信度。

```shell
├─ biz          // 事务层，业务代码
│  ├─ handler   // 请求处理，控制器层
│  ├─ model     // 调用的模型 DTO
│  ├─ router    // 路由层，路由组中间件管理
│  └─ rpc       // RPC 调用工具函数
├─ script       // 运行脚本
├─ main.go      // 程序入口
└─ Makefile     // 常用指令
```

## Microservice 微服务

对于每个不同的微服务，我们都有类似的文件结构，如下所示，

```shell
user/
├── configs     // 微服务启动配置，包括数据库连接
├── service     // 微服务业务代码，服务层
├── model       // 微服务相关结构体，模型层
├── script      // 编译指令，由 Kitex 生成
├── pack        // 请求封装，序列化层，将 BO -> DTO
├── utils       // 通用函数，在项目根路径中
├── main.go     // 程序入口
├── handler.go  // 请求处理，控制器层
├── kitex.yaml  // Kitex 配置项
└── Makefile    // 常用指令
```
