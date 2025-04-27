# Tasks

需要支持Bag的增删改查和task的增删改查以及调度

# 任务调度流程

当一个task被调度给某个agent的时候，我们需要重新生成一个AccessKey并提供给agent，agent可以通过这个AccessKey去获取对应的task

# Table Define

使用GORM 和 PGSQL

## Bags

|column name | column type | description |
|:----:|:----:|:----:|
|BagName|string| uuid |
|BagDisplayName|string| customize name |
|CreateTimeMs|int64| Bag创建的时间 |
|UpdateTimeMs|int64| Bag上次修改的时间 |

## Tasks

|column name | column type | description |
|:----:|:----:|:----:|
|TaskName|string| uuid |
|TaskDisplayName|string| customize name |
|BagName|string| uuid |
|ScriptPath| string | 需要执行的脚本的位置 |
|Script|string| 用户直接提供的脚本内容 |
|WorkingDir| string | 脚本的运行位置 |
|Priority| int16 | 任务的优先级（默认为1, 最大限制为10000, Priority越小，优先级越高） |
|OrderId| uint32 | 任务入队的顺序Id |
|AccessKey | string | 访问权限密钥 |
|NodeId | string | 任务执行的node的id |
|CreateTimeMs|int64| Task创建的时间 |
|FinishTimeMs|int64| Task结束的时间 |
|ScheduledTimeMs|int64| Task被调度的时间 |
