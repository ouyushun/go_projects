package com.oys.bigdata.spark.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_intersection_union_subtract_zip {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    //分区内数据的执行是有序的
    //不同分区内数据的执行是无序的


    //zip 拉链
    //拉链操作的2个RDD, 分区数和分区内元素个数必须相等
    val rdd1 : RDD[Int] = sc.makeRDD(List(1, 2, 3, 4), 8)
    val rdd2 : RDD[Int] = sc.makeRDD(List(3, 4, 5, 6), 8)

    val value = rdd1.zip(rdd2)
    println(value.collect().mkString(","))

    //关闭
    sc.stop()
  }

}
