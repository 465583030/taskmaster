# taskmaster
分布式任务调度中心

## 模块图
![](doc/taskmaster.png)

jobmaster

+ 1.控制中心负责存储任务
+ 2.控制中心负责分配、终止任务
+ 3.jobmaster是长链接,有状态服务.一旦jobmaster停止,则所有节点没有任务可以执行.