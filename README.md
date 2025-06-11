### 以 gin为基础搭建的 DDD 项目架构
以下为单个 main函数入口的工程结构示例，具体代码在 single_cmd分支，
切换到 multi_cmd分支可查看多个 main函数入口示例
``` 
.
├── api                     #  前后端接口参数，DTO 对象
├── config.yaml             # 系统配置
├── idl                     #  api 中 DTO 对象的 PB 定义
├── internal
│   ├── biz                 # http层路由和业务逻辑，在 http 层不直接读写 db，而是通过domain 进行操作
│   │   ├── handler         # DDD 架构中的 application 层，处理主要的业务逻辑
│   │   └── router          # 路由，处理参数的解析和返回结果
│   ├── libs                # DDD 架构 中 infra 层，存放数据 model 和持久化操作
│   │   ├── db              # 数据库的初始化逻辑 
│   │   ├── config          # 系统配置初始化逻辑 
│   │   ├── logger          # 日志包初始化逻辑   
│   │   └── router          # redis 的初始化逻辑
│   ├── loader              # 对多个service的并发访问需要包转成 loader
│   ├── consts              # 公共常量、变量，不包含 api 层的常量，api层的常量和变量属于接口参数的一部分 
│   ├── utils               # 常用工具函数               
│   └── service             # DDD 架构中的 domain 层，domain包含一个原子化的业务，涉及到多张数据表或者service的操作
│   │   ├── domain1         # domain领域名称
│   │   │   ├── api         # domain 接口的入参和出参，DO 对象
│   │   │   ├── handler     # domain 实际的业务处理
│   │   │   ├── loader      # 对下游多个service 接口的并发封装，非必需
│   │   │   ├── model       # 数据表 持久化对象，PO 对象
│   │   │   ├── repo        # 数据库访问层，service 直接操作 model 时需要           
│   │   │   └── service     # domain 的对外接口
├── Dockerfile
├── docker-compose.yaml
├── README.md
├── go.mod
├── go.sum
└── main.go
```
> 每一个 domain 单独抽出来可以视作一个 rpc 的微服务

## 部署
部署参考 feature/deploy_image 分支

### docker 镜像编译和部署
1. 编译镜像，指定名称和版本
```
docker build -t name:tag ./
docker build -t test:0.0.1 ./
```

2. 部署镜像，将/data/conf与宿主机路径关联，config.yaml下发到/data/conf中即可读取
```
docker run -v 宿主机路径:挂载路径 test
docker run -v /Users/hezhichang/go/src/gin_layout:/data/conf test 
```
