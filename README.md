# taskmaster
分布式任务调度中心

## 模块图
![](doc/taskmaster.png)

+ jobmaster

    + 1.控制中心负责存储任务
    + 2.控制中心负责分配、终止任务
    + 3.jobmaster是长链接,有状态服务.一旦jobmaster停止,则所有节点没有任务可以执行.
    + 4.jobmaster重启后,从数据库、文件系统载入任务列表,重新开启服务,接入长链接,并分配任务
+ worker
    + 1.worker负责具体执行任务
+ register
    + 1.register模块负责向jobmaster注册新任务,jobmaster可以控制register的数量,通常为1,防止重复添加任务
 