# kafka 的安装和使用
* 启动 zookeeper 容器：`docker run -d --name zookeeper -p 2181:2181 -v /Users/bamboo/Documents/suhy/svcData/kafka/zookeeper/data:/data -v /usr/local/zookeeper/log:/datalog zookeeper`

* 启动 kafka 容器：`docker run -d --name kafka --publish 9092:9092 \
  --link zookeeper \
  --env KAFKA_ZOOKEEPER_CONNECT=192.168.88.40:2181 \
  --env KAFKA_ADVERTISED_HOST_NAME=192.168.88.40 \
  --env KAFKA_ADVERTISED_PORT=9092  \
  --env KAFKA_LOG_DIRS=/kafka/kafka-logs-1 \
  -v /Users/bamboo/Documents/suhy/svcData/kafka/kafka/logs:/kafka/kafka-logs-1  \
  wurstmeister/kafka`

* 创建 topic：
    * `cd /opt/kafka/bin`
    * `kafka-topics.sh --create --zookeeper 192.168.88.40:2181 --replication-factor 1 --partitions 1 --topic test`
* 查看 topic：
    * `kafka-topics.sh --list --zookeeper 192.168.88.40:2181`

## 案例
公司最近 kafka 队列中堆积很严重，有 10 个分区，每个分区都堆积了 900w 的消息，消费者也一直在消费，大约 20条/秒。
总体来讲，此时的消费速率略大于生产速率，导致消息堆积量一直维持在 900w 左右，很头疼，于是想办法解决，要解决，就需要充分了解现有架构。
了解完，再想方案解决该问题。

## reference
* https://www.hangge.com/blog/cache/detail_2791.html
* go 实现生产者、消费者 https://blog.csdn.net/luslin1711/article/details/105798571