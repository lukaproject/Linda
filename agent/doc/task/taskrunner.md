# task runner

这是agent的一个模块，用于执行脚本

## 需要提供的功能
- 启动一个线程运行脚本，监控脚本运行的状态以及运行结果。
- 对脚本进行权限控制
- 需要支持创建 / 删除 / 修改 本机的用户

## 技术

- 目前期望是使用 gopsutil 和 exec 模块, 当然其他的模块也可以，只要用的上的都行

## Task 参数
- pathToScript 脚本路径
- workingDir 脚本运行路径
- Resource 占用多少资源
- process 进程