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

    // TODO... 对接网络数据
   val lines = ssc.socketTextStream("localhost", 9528)

    // TODO... 业务逻辑处理
    // 输入数据以逗号分隔
    val result = lines.flatMap(_.split(",")).map((_, 1))
      .reduceByKey(_ + _)

    //    lines.count().print()

    //    val result = lines.flatMap(_.split(","))
    result.print()

    //result.saveAsTextFiles("pk")

    // 启动
    ssc.start()
    ssc.awaitTermination()
  }
}
