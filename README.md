# OpenIoT-Server

**OpenIoT-Server** 是 OpenIoT-Hub 项目的分布式后台管理平台，基于 Go 语言编写、使用 Kitex 框架和 protobuf IDL 的分布式后台管理系统，旨在为物联网提供服务。

## 特性

- 分布式架构
- 高效的通信协议
- 可扩展性强
- 支持定制化开发

## 服务运行

我们使用 Docker 创建生产环境，故使用前请使用指令生成运行环境：

```shell
docker-compose up -d
```

然后我们就可以启动各个微服务模块，所有微服务模块都被存放在 `/internal` 之中。切入微服务目录，使用指令：

```shell
sh build.sh
sh output/bootstrap.sh
```

最终，你需要运行网关，切入目录 `/api-gateway`，使用指令：

```shell
sh run.sh
```

## 使用说明

1. API 接口

   该系统提供了一系列 API 接口，用于管理分布式系统，包括设备注册、数据采集、命令下发等功能。具体使用方式请参考接口文档。

2. 数据库配置

   该系统使用 MySQL 数据库存储数据，支持使用默认配置，也可以在运行前进行相应的数据库配置。请修改 `config/config.toml` 文件中的数据库连接信息，以及 `deploy/sql/init.sql` 文件中的数据表结构。

3. 日志输出

   该系统使用 [glog](https://github.com/golang/glog) 进行日志输出。您可以在各个微服务最终包装模块下的 `log/` 下查看相关日志信息。


## Project Structure

```shell
├─api-gateway     // 服务器网关模块
├─config          // 公共配置文件
├─deploy          // 部署文件，包括 docker-compose.yaml 等
├─idl             // 接口定义文件，使用的语言为 protobuf
├─internal        // 内部服务，包含所有微服务模块
│  ├─ user        // 用户微服务模块
│  └─ ...         // 其他微服务模块
├─kitex_gen       // RPC 存根，由 Kitex 生成
└─pkg             // 第三方函数
    ├─consts      // 项目公共常量
    ├─errno       // 项目错误处理封装定义
    ├─logging     // 日志工具函数
    └─util        // 其他工具函数，例如雪花算法生成唯一 ID
```

### Gateway 网关

这里是网关的项目结构，其负责通过远程过程调用(Remote Process Call, RPC)唤起对应的微服务。同时，网关也负责与鉴权服务进行交互，验证请求可信度。

```shell
api-gateway/
├─ biz          // 事务层，业务代码
│  ├─ handler   // 请求处理，控制器层
│  ├─ model     // 调用的模型 DTO
│  ├─ router    // 路由层，路由组中间件管理
│  └─ rpc       // RPC 调用工具函数
├─ script       // 运行脚本
├─ main.go      // 程序入口
└─ Makefile     // 常用指令
```

### Microservice 微服务

对于每个不同的微服务，我们都有类似的文件结构，这里我们使用用户模块举例，如下所示，

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

## 贡献指南

欢迎您为 openiot-server 项目作出贡献！您可以通过以下方式参与项目：

- 提出 Issue。

- 编写和改进代码，提交 Pull Request。

## 许可证

该项目基于 [MIT](https://github.com/your_username/openiot-server/blob/main/LICENSE-MIT) 开源协议发布。

