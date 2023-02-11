package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_groupBy {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(List("Hello", "Spark", "Hello", "Hadoop", "Scala"), 3)

    val groupRdd = rdd.groupBy(
      _.charAt(0)
    )

    println(groupRdd.collect().mkString)
    //关闭
    sc.stop()
  }

}
