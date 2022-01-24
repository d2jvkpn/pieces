# httpfileserver
简单的 http 静态文件服务

1. 运行
```bash
go run main.go -address ":8080"
```

2. 参数说明
- localPath, 本地文件路径, 默认当前路径;
- prefix, http 服务地址前缀, 默认 /, 如果设置为 /static, 则使用 /static/path/to/file 路径访问;
- address, http 服务地址, 默认 :8080;

