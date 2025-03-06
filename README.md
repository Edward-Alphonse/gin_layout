### 以 gin为基础搭建的 DDD 项目架构
以下为单个 main函数入口的工程结构示例，具体代码在 single_cmd分支，
切换到 multi_cmd分支可查看多个 main函数入口示例
``` 
.
├── api                      #  前后端接口参数，DTO 对象
├── config                   # 系统配置
│   └── config.yaml
├── idl                      #  api 中 DTO 对象的 PB 定义
├── internal
│   ├── biz                  # 路由和业务逻辑
│   │   ├── handler         # DDD 架构中的 application 层，处理主要的业务逻辑
│   │   └── router          # 路由，处理参数的解析和返回结果
│   ├── config               # 系统配置初始化逻辑 
│   ├── db                   # 数据库的初始化逻辑 
│   ├── infra                # DDD 架构 中 infra 层，存放数据 model 和持久化操作
│   │   ├── loader          # 对数据访问层接口的并发封装，非必需
│   │   ├── model           # 数据表 持久化对象，PO 对象
│   │   └── repo            # 数据库访问层
│   ├── logger               # 日志包初始化逻辑   
│   ├── consts               # 公共常量、变量 
│   ├── utils                # 常用工具函数
│   ├── redis                # redis 的初始化逻辑
│   └── service              # DDD 架构中的 domain 层
│   │   ├── domain1          # domain领域名称
│   │   │   ├── entity      # domain 接口的入参和出参，DO 对象
│   │   │   ├── handler     # domain 实际的业务处理
│   │   │   ├── loader      # 对 service 接口的并发封装，非必需
│   │   │   └── service     # domain 的对外接口
├── Dockerfile
├── docker-compose.yaml
├── README.md
├── go.mod
├── go.sum
└── main.go
```