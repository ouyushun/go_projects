package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_glom {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(List(11, 22, 33, 55, 66), 3)
    val glomRdd : RDD[Array[Int]] = rdd.glom()
    //需要注意的是，匿名函数的返回值的数据类型依赖于函数体的最后一行表达式的值，
    // 这个由程序自己判断，匿名函数的返回值类型不能手工指定！
    glomRdd.collect().foreach(
      (data : Array[Int] ) =>{
        println(data.mkString)
      }
    )

    //分区内最大值求和
    val rdd3 = sc.makeRDD(List(1,2,3,4), 2)
    val rdd5 = rdd3.glom()
    val maxRdd = rdd5.map(
      array => {
        array.max
      }
    )
    println(maxRdd.collect().sum)

    //关闭
    sc.stop()
  }

}
