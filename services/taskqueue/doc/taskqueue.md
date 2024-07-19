# Task Queue

# 功能
创建队列 / 删除队列 / 任务入队 / 任务出队 

任务具有优先级，出队顺序按照优先级从小到大出队。


# V0
可以暂时使用Redis的ZSet来代替，key值设置为 Priority * 1e10 + orderId