package com.oys.bigdata.spark.core.acc

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Acc_01 {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("ACC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(
      List(1,2,3,4,5,6), 2
    )

    val sumAcc = sc.longAccumulator("sum")

    //sc.doubleAccumulator("s")
    //sc.collectionAccumulator("s")

    rdd.foreach(
      num => {
        sumAcc.add(num)
      }
    )

    println(sumAcc.value)

    //关闭
    sc.stop()
  }

}
