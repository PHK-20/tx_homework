# tx_homework

go实现访问mysql缓存到redis的web服务

## Installation  

* 安装go依赖
```shell
go mod tidy
go mod vendor
```
* 启动mysql和redis本地服务

* 配置mysql,修改本地对应的用户名、密码、端口和ip

db/conf.ini
```ini
[mysql]
db_user=root
db_pw=123456
db_name=go_demo
```

install.sh  
```shell
# mysql conf
host="127.0.0.1"
port="3306"
db_user="root"
db_pw="123456"
```

* 终端运行intall.sh
```shell
chmod u+x install.sh
./install.sh
```

* 编译启动二进制
```shell
go build
./go_redis  
```

* 打开浏览器测试web服务  
get请求：localhost:9000/GetUser  
第一次请求会从mysql读取数据写入redis缓存，输出到终端  
第二次请求会从redis读出数据，输出到终端  
redis过期时间为15s

