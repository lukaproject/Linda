# Agent

## 需要提供的功能
- 上报机器资源状态
- 运行命令行
- 控制杀死 / 启动 / 中断进程
## 技术选型

- Golang
- 使用Websocket与agentcentral的通信, 使用json用来序列化包体。

## 通信方式

- 计算节点主动请求agentcentral, 上报信息，获取任务以及对进程的控制请求

## 需要共享的数据
- 计算结果
- 机器资源的状态（剩余内存 显存 CPU核数 正在运行的任务数等）

## 任务下发方式
- 任务id通过agentcentral发送给agent, agent在启动任务的时候，通过call api的方式从agentcentral获取