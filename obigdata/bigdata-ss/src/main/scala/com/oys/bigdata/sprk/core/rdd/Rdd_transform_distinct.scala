package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_distinct {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(List(11, 22, 33, 55, 66, 11,22,33), 3)

    println(rdd.distinct().collect().mkString)

    //关闭
    sc.stop()
  }

}
