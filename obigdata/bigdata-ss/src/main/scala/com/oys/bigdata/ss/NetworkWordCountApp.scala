package com.oys.bigdata.ss

import org.apache.spark.SparkConf
import org.apache.spark.streaming.{Seconds, StreamingContext}


object NetworkWordCountApp {
	def main(args: Array[String]): Unit = {
		// 入口点

		val sparkConf = new SparkConf()
		  .setAppName(this.getClass.getSimpleName)
		  .setMaster("local[2]")

		// 指定间隔5秒为一个批次
		var ssc: StreamingContext = new StreamingContext(sparkConf, Seconds(5))

		//nc -lk 9999


		// TODO... 对接网络数据
		val lines = ssc.socketTextStream("localhost", 9999)

		// TODO... 业务逻辑处理
		// 输入数据以逗号分隔
		val result = lines.flatMap(_.split(" ")).map((_, 1)).reduceByKey(_ + _)

		result.print()

		// 由于SparkStreaming采集器是长期执行的任务，所以不能直接关闭
		// 如果main方法执行完毕，应用程序也会自动结束。所以不能让main执行完毕
		//ssc.stop()
		// 启动采集器
		ssc.start()
		// 2. 等待采集器的关闭
		ssc.awaitTermination()
	}
}
