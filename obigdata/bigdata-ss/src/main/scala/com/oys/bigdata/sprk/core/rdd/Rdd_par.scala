package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_par {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    //操作
    //RDD的并行度 he 分区
    val rdd = sc.makeRDD(List(1, 2, 3, 4), 3)
    rdd.saveAsTextFile("")

    //关闭
    sc.stop()
  }

}
