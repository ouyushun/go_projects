package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_groupByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd  = sc.makeRDD(
      List(("a", 1), ("a", 2), ("a", 3), ("b", 6), ("b", 8))
    )

    val groupByKeyRdd  = rdd.groupByKey()
    /*
    * (a,CompactBuffer(1, 2, 3))
      (b,CompactBuffer(6, 8))
    * */
    groupByKeyRdd.collect().foreach(println)


    val groupByRdd  = rdd.groupBy(_._1)
    /*
    * (a,CompactBuffer((a,1), (a,2), (a,3)))
      (b,CompactBuffer((b,6), (b,8)))
    * */
    groupByRdd.collect().foreach(println)
    //关闭
    sc.stop()
  }

}
