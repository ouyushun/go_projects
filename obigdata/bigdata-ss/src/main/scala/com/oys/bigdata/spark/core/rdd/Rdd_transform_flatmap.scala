package com.oys.bigdata.spark.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_flatmap {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)


    var rdd: RDD[List[Int]] = sc.makeRDD(List(List(1, 2, 3), List(4, 5, 6,7,8,9)))
    var flatRdd = rdd.flatMap(
      list => {
        list
      }
    )
    flatRdd.collect().foreach(println)

    //wordCount
    val rdd2 = sc.makeRDD(List("hello world", "hello spark"))
    val rdd3 = rdd2.flatMap(
      s => {
        s.split(" ")
      }
    )
    rdd3.collect().foreach(println)

    //模式匹配
    val list2 = List(List(1,2), 3, List(4,5))
    val rdd4 = sc.makeRDD(list2)
    val rdd5 = rdd4.flatMap {
      case list: List[_] => list
      case i => List(i)
    }
    rdd5.collect().foreach(println)

    //关闭
    sc.stop()
  }

}
