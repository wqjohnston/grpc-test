#### 执行生成proto文件命令  
```
protoc --go_out=. --go-grpc_out=. hello.proto  
```

#### 启动server  
```
go run server/server.go
```

#### 启动client  
另起一个窗口  
```
go run client/client.go
```

#### 安装gprc环境参考  
https://blog.csdn.net/Johnston_man/article/details/140466843  
