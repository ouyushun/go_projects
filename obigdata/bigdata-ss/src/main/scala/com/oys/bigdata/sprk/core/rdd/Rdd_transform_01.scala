package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.{SparkConf, SparkContext}
import org.apache.spark.rdd.RDD
/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_01 {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    //操作
    //RDD的并行度 he 分区
    val rdd:RDD[Int] = sc.makeRDD(List(1, 2, 3, 4), 3)

    def mapFunc(num : Int) =  num * num
    rdd.map(mapFunc).collect().foreach(println)
    //匿名函数
    rdd.map((num:Int) =>{num * 2}).collect().foreach(println)
    rdd.map((num:Int) =>num * 2).collect().foreach(println)
    rdd.map(num =>{num * 2}).collect().foreach(println)
    rdd.map(_ * 100).collect().foreach(println)

    val rdd2 : RDD[String] = sc.textFile("apache.log")
    val mapRdd2 = rdd2.map(
      line => {
        val dates = line.split(" ")
        dates(1)
      }
    )
    mapRdd2.foreach(println)

    //关闭
    sc.stop()
  }

}
