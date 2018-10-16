channel:

子为生产者,父为消费者
结论1:
子做生产者发送完数据确定一定要close channel,否则无论消费者如何弄都会死锁
结论2:
子做生产者close channel,父消费者要用for+if或者for+range来读取channel,才能识别出close channel,否则读出大量空数据

子为消费者,父为生产者
结论1:
父可以不close channel
结论2:
父做生产者close channel,子消费者要用for+if或者for+range来读取channel,才能识别出close channel,否则读出大量空数据
