package com.oys.bigdata.spark.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_filter {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(List(1,2,3,4,5,6,7,8,9), 3)
    val filterRdd = rdd.filter(
      num => num % 2 == 0
    )
    println(filterRdd.collect().mkString(","))

    //关闭
    sc.stop()
  }

}
