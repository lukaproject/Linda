# Tools

目前需要安装的依赖以及环境准备都在这个文件夹

## 下载新版本golang
```bash
go install golang.org/dl/go1.23.4@latest
go1.23.4 download
```

## Python 3.10+

请保证在Linux中开发并且Python Version >= 3.10

## ENV
### build ENV

- go version 1.23+, 目前是 go1.23.4
- 使用installswag.sh 安装 swag
- run
```bash
source .profile
```

### external dependencies

[给docker配置网络代理](https://www.cnblogs.com/Chary/p/18096678)

需要docker环境，docker中安装pgsql / redis. 目前使用的pgsql版本为15，redis版本为7.0


### Setup pqsql
```bash
docker pull postgres:15
```

```bash

docker run -d \
  --name pgsql15 \
  -e POSTGRES_USER=dxyinme \
  -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=linda \
  -p 5432:5432 \
  postgres:15

```

### swagger codegen
```bash
docker pull swaggerapi/swagger-codegen-cli-v3
```
在根目录下使用下面这个命令生成swagger client
```bash
./baselibs/apiscall/renew-swagger.sh
```

### docker buildx install

- [buildx release](https://github.com/docker/buildx/releases)

1. 需要将下载下来的二进制重命名为docker-buildx后放入~/.docker/cli-plugins/
2. 需要修改docker配置文件, 增加experimental: "enabled"这条
```bash
cat .docker/config.json
{
    "experimental": "enabled"
}

# 重启docker reload配置文件
systemctl restart docker

# 判断是否安装成功
docker buildx version
```

###  Local Test
build image
```
如果没有buildx: docker pull docker/buildx-bin
python3 tools/builder/builder.py --agent --agentcentral --fileservicefe

```

```
source .profile
./tools/testscripts/devstart.sh
```