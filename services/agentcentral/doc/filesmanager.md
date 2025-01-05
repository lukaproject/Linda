# Files Manager

## Overview

需要对上传 / 下载的文件进行管理。上传文件保存在一个本地文件夹内。

文件夹架构如下:
```
{rootdir}/{block}/{file name}
```
## APIs

### upload
```
[post] /api/files/upload

form:
fileName: string
block: string
multiform:
file: stream
```
### download
```
[get] /api/files/download/{block}/{filename}
```