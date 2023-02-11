package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{HashPartitioner, SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_reduceByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd  = sc.makeRDD(
      List(("a", 1), ("a", 2), ("a", 3), ("b", 6))
    )

    val reduceRdd = rdd.reduceByKey(
      (x: Int, y: Int) => {
        x + y
      }
    )
    reduceRdd.collect().foreach(println)
    //关闭
    sc.stop()
  }

}
