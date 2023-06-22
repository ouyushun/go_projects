package com.oys.bigdata.spark.core.rdd.other

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_foreach {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)


    val rdd = sc.makeRDD(List(1, 2, 3, 4,5,6), 8)

    rdd.toDebugString

    //在Driver打印
    rdd.collect().foreach(println)

    println("*********")

    //在Excutor执行， 分布式执行， 打印是无序
    rdd.foreach(println)


    //关闭
    sc.stop()
  }

}
