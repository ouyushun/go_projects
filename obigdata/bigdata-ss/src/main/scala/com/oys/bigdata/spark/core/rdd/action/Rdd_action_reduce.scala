package com.oys.bigdata.spark.core.rdd.action

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_action_reduce {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(List(1,2,3,4), 3)
    //res == 50
    //初始值既参与分区内的计算， 也参与分区间的计算。
    val res = rdd.aggregate(10)(
      _+_,
      _+_
    )

    // fold 当分区内外计算规则相同， fold简化即可。
    val res2 = rdd.fold(10)(
      _+_
    )

    println(res)

    //关闭
    sc.stop()
  }

}
