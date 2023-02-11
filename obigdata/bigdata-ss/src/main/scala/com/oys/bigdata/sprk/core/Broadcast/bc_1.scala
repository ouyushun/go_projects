package com.oys.bigdata.sprk.core.Broadcast

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object bc_1 {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd1 = sc.makeRDD(
     List(("a", 1), ("b", 2), ("c", 3), ("c", 5))
    )
    val rdd2 = sc.makeRDD(
      List(("a", 11), ("b", 21), ("c", 33), ("c", 44))
    )

    val value: RDD[(String, (Int, Int))] = rdd1.join(rdd2)

    value.collect().foreach(println)

    //关闭
    sc.stop()
  }

}
