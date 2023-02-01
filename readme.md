项目结构

├── Makefile                        # kitex的生成rpc框架代码，注意生成位置  
├── cmd                             # 业务代码  
│   ├── api                         # hertz暴露接口层  
│   │   ├── handlers                # 给hertz绑定的handler合集  
│   │   │   ├── handler.go          # req和resp的定义  
│   │   │   ├── info.go             # 用户信息handler  
│   │   │   ├── login.go            # 用户登录handler  
│   │   │   └── register.go         # 用户注册handler  
│   │   ├── main.go                 # hertz启动的文件  
│   │   ├── rpc                     # rpc client端代码  
│   │   │   ├── init.go             # rpc client 初始化   
│   │   │   └── user.go             # 实际初始化rpc client并获取rpc user服务   
│   │   └── run.sh                  # **实际运行server脚本**  
│   ├── comment                     # 评论服务  
│   ├── relation                    # 关系服务  
│   ├── thumb-up                    # 点赞服务  
│   ├── user                        # 用户服务  
│   │   ├── Makefile                # hertz生成当前服务框架代码，注意位置  
│   │   ├── build.sh                # **build 用户服务**    
│   │   ├── dal                     # 数据层    
│   │   │   ├── db                  # db代码    
│   │   │   │   ├── init.go  
│   │   │   │   └── user.go   
│   │   │   └── init.go    
│   │   ├── handler.go              # rpc服务端代码  
│   │   ├── idl                     # 复制的一份idl，proto3在生成时与thrift相比有限制，不能用../    
│   │   │   └── user.proto  
│   │   ├── kitex.yaml              # 生成文件，不用管    
│   │   ├── main.go                 # 运行用户服务  
│   │   ├── output                  # 生成文件目录   
│   │   │   ├── bin  
│   │   │   │   └── user  
│   │   │   ├── bootstrap.sh        # **实际运行脚本**  
│   │   │   └── log     
│   │   │       ├── app  
│   │   │       └── rpc  
│   │   ├── pack                    # 抽象resp包的发送  
│   │   │   └── resp.go  
│   │   ├── rpc  
│   │   ├── script                  
│   │   │   ├── bootstrap.sh  
│   │   │   └── log  
│   │   │       ├── app     
│   │   │       └── rpc     
│   │   ├── service                 # 业务层        
│   │   │   ├── user_info.go            
│   │   │   ├── user_login.go       
│   │   │   └── user_register.go        
│   │   └── test        
│   │       └── test.go     
│   └── video                       # video服务     
├── docker-compose.yml              # docker compose up 一下   
├── go.mod                          
├── go.sum          
├── idl         
│   └── user.proto                  # 复制api，改2为3，2感觉有bug
├── kitex_gen                       # 生成目录
│   └── user        
│       ├── user.pb.fast.go     
│       ├── user.pb.go      
│       └── userservice     
│           ├── client.go       
│           ├── invoker.go      
│           ├── server.go       
│           └── userservice.go          
├── middleware                      # JWT鉴权   
│   └── jwt.go                      
├── pkg                             # 工具包    
│   ├── configs                     # 数据库初始化  
│   │   └── sql     
│   │       └── init.sql        
│   ├── constants                   # 一些配置常量      
│   │   └── constant.go         
│   ├── errno       
│   │   └── errno.go                # 自定义的errno     
│   └── tracer                      # Jaeger相关        
│       └── tracer.go       
└── readme.md       

参考easy_note的结构，采用插拔设计

主要技术：
hertz
kitex
etcd
mysql
jwt
七牛云tos

之后可以考虑引入redis和消息队列
