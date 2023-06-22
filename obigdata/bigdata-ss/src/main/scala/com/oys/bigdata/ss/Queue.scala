package com.oys.bigdata.ss

import org.apache.spark.SparkConf
import org.apache.spark.rdd.RDD
import org.apache.spark.streaming.{Seconds, StreamingContext}

import scala.collection.mutable


object Queue {
	def main(args: Array[String]): Unit = {
		// 入口点

		val sparkConf = new SparkConf()
		  .setAppName(this.getClass.getSimpleName)
		  .setMaster("local[2]")

		// 指定间隔5秒为一个批次
		var ssc: StreamingContext = new StreamingContext(sparkConf, Seconds(5))

		val rddQueue = new mutable.Queue[RDD[Int]]()

		val inputStream = ssc.queueStream(rddQueue,oneAtATime = false)
		val mappedStream = inputStream.map((_,1))
		val reducedStream = mappedStream.reduceByKey(_ + _)
		reducedStream.print()

		// 启动采集器
		ssc.start()

		for (i <- 1 to 5) {
			rddQueue += ssc.sparkContext.makeRDD(1 to 300, 10)
			Thread.sleep(2000)
		}





		// 由于SparkStreaming采集器是长期执行的任务，所以不能直接关闭
		// 如果main方法执行完毕，应用程序也会自动结束。所以不能让main执行完毕
		//ssc.stop()

		// 2. 等待采集器的关闭
		ssc.awaitTermination()
	}
}
