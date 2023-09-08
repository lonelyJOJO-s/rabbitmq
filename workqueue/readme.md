### core
disturbe time-consuming tasks among multi workers
### using scenes
during a short HTTP request window, hard to handle a complex task

### diff from goroutine
we can easily set the counts of consumer(goruntine)

### issues
#### 1. once receiver recives the msg, the msg was marked as deleted which may cause task losing
    solution: ACK mechanism(consumer finish task retrun a ack to rabbitmq), if a consumer dies, rabbitmq will requeue the task ------(default ack waiting time: 30s)
#### 2. once rabbimq server crushs, all the tasks will be lost either
    solutoin:
    1. make sure queue durable
    2. make sure msg durable


