# Tools

目前需要安装的依赖以及环境准备都在这个文件夹

## 下载新版本golang
```bash
go install golang.org/dl/go1.23.4@latest
go1.23.4 download
```

## ENV
### build ENV

- go version 1.22+, 目前是 go1.23.4
- 使用installswag.sh 安装 swag
- run
```bash
source .profile
```

### external dependencies

[给docker配置网络代理](https://www.cnblogs.com/Chary/p/18096678)

需要docker环境，docker中安装pgsql / redis. 目前使用的pgsql版本为15，redis版本为7.0

### swagger codegen
```bash
docker pull swaggerapi/swagger-codegen-cli-v3
```
在根目录下使用下面这个命令生成swagger client
```bash
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3:latest generate \
    -i /local/services/agentcentral/docs/swagger.yaml \
    -l go \
    -o /local/out/go
```