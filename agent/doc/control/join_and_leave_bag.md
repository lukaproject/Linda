# Join And Leave Bag

## Join Bag

- AgentCentral 提供 join API, 每个node只能够加入一个bag
- AgentCentral 在收到 join request的时候，需要将内存中的node state改为 joining 状态
- AgentCentral 为joining状态时，每次agent的HB回包都会为带上joinBag包体，内部有对应bagName
- Agent 收到的HB回包中如果存在 joinBag包体，则进入join bag流程
    1. join bag流程完成之后，HB的agent包包体内就会包含带有bagName的nodeInfo
    2. 如果在 join bag流程中，又收到了其他bagName的join bag包，那么拒绝其他的join bag包
- AgentCentral 收到的HB包中如果存在NodeInfo，并且其中的BagName与上一次发到这个Agent的join bag的bagName相同的话，则node state进入 on work状态，此时这个Bag里面的task会被schedule到它上面

## Leave Bag

- AgentCentral 提供 free API，只有on work状态的node才可以处理这个request
- AgentCentral 在收到 free node request的时候，会将内存中的node state改为freeing状态
- Agent收到的HB回包中如果存在freeNode包体，则进入free node流程
- AgentCentral 收到的HB包中的NodeInfo如果不为空，则说明node state已经进入Free状态