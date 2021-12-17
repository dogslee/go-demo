# golang select for return

> 注意： 无缓冲channel消费端声明一定要早于生产端，否则一直死锁等待

优雅的通过一个channel 的close进而通知接收者结束goroutine

退出select 三种方式：

1. break + tag
2. goto + tag
3. return
