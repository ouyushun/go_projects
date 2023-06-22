package com.oys.bigdata.spark.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_cogroup {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd1 = sc.makeRDD(
     List(("a", 111), ("b", 211), ("c", 31), ("d", 55))
    )
    val rdd2 = sc.makeRDD(
      List(("a", 11), ("b", 21), ("c", 33), ("c", 44))
    )

    val value: RDD[(String, (Iterable[Int], Iterable[Int]))] = rdd1.cogroup(rdd2)


    //value.collect().foreach(println)
    value.collect().foreach(println)


    //关闭
    sc.stop()
  }

}
