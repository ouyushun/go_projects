package com.oys.bigdata.spark.core.CasePractice.top3

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_action_countByValue_countByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd  = sc.makeRDD(
      List(("a", 11), ("a", 11), ("a", 31), ("b", 66), ("a", 66)), 2
    )

    println(rdd.countByValue())
    println(rdd.countByKey())

    //关闭
    sc.stop()
  }

}
