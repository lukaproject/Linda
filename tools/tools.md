# Tools

目前需要安装的依赖以及环境准备都在这个文件夹

## 下载新版本golang
```bash
go install golang.org/dl/go1.21.12@latest
go1.21.12 download
```

## ENV
### build ENV

- go version 1.21+, 目前是 go1.21.12
- 使用installswag.sh 安装 swag
- run
```bash
source .profile
```

### external dependencies

需要docker环境，docker中安装pgsql / redis. 目前使用的pgsql版本为15，redis版本为7.0