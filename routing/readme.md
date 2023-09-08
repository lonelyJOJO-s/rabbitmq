### core: msg can deliver to routing queue(queue can subscribe the msg)
#### we use direct type of exchanger in this demo
![image](https://www.rabbitmq.com/img/tutorials/direct-exchange.png)

in this section I itimate a app with two kind of users: normal and VIP
VIP will get all the info from exchanger and normal will noly get the normal msg
