### 简介
> go gin脚手架， 一些简单的封装

### 安装
```shell
# 拉取依赖
go mod tidy

# 修改相关配置
cat .env.example > .env
```

### 启动
```shell
go run main.go

# or

go build main.go && ./main
```

### 目录说明
~~~
gin-skeleton
├─app                  
│  ├─api        接口
│  ├─cron       定时任务
│  ├─form       接收表单数据
│  ├─library    包
│  ├─middleware 中间件
│  ├─models     模型：自定义模型不进行数据操作
│  ├─servers    服务：封装curd具体操作
│  
├─config        项目配置文件
├─router  
│  ├─router.go  路由定义
├─main.go       入口
~~~