# kafka

## kafka的性能

JVM性能调优使用kafka-server-start.sh
kafka自带的测试工具：针对生产者的kafka-producer-perf-test.sh和针对消费者的kafka-consumer-perf-test.sh

kafka使用zookeeper来做分布式同步，zookeeper使用fast paxos的共识算法

