package com.oys.bigdata.ss

import org.apache.spark.SparkConf
import org.apache.spark.storage.StorageLevel
import org.apache.spark.streaming.receiver.Receiver
import org.apache.spark.streaming.{Seconds, StreamingContext}

import java.io.{BufferedReader, InputStreamReader}
import java.net.Socket
import java.nio.charset.StandardCharsets


object KafkaReceiver {
	def main(args: Array[String]): Unit = {
		// 入口点

		val sparkConf = new SparkConf()
		  .setAppName(this.getClass.getSimpleName)
		  .setMaster("local[2]")

		// 指定间隔5秒为一个批次
		var ssc: StreamingContext = new StreamingContext(sparkConf, Seconds(5))





		ssc.start()
		// 2. 等待采集器的关闭
		ssc.awaitTermination()
	}

	class MyReceiver(host: String, port: Int) extends Receiver[String](StorageLevel.MEMORY_ONLY) {

		var socket: Socket = _

		override def onStart(): Unit = {
			new Thread(new Runnable {
				override def run(): Unit = {
					receive()
				}
			}).start()
		}

		override def onStop(): Unit = {
			if(socket != null) {
				socket.close()
				socket = null
			}
		}

		def receive() {
			socket = new Socket(host, port)
			val reader = new BufferedReader(new InputStreamReader(socket.getInputStream, StandardCharsets.UTF_8))
			var line: String = null
			while ((line = reader.readLine()) != null) {
				this.store(line)
			}
		}

	}
}
