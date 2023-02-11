package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_aggregateByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd  = sc.makeRDD(
      List(("a", 11), ("a", 21), ("a", 31), ("b", 66)), 2
    )

    println(rdd.collect().mkString("--"))

    //存在函数柯里化， 两个参数列表
    //第一个参数列表： 初始值
    //第二个参数：
    //   第一个参数： 分区内计算规则
    //   第二个参数： 分区间计算规则
    val aggRdd = rdd.aggregateByKey(0)(
      //第一个x为给的初始值zeroValue， y为RDD的value
      (x, y) => {
        math.max(x, y)
      },
      (x, y) => x + y
    )

    val rdd2 = sc.makeRDD(
      List(("a", 3), ("a", 6), ("a", 3), ("b", 6), ("b", 4)), 2
    )

    //求平均值
    //初始值为 (0, 0) : (sum值, 次数)
    //第一个参数列表(U, V) => U,
    //第二个参数列表(U, U) => U
    val aggRdd2 = rdd2.aggregateByKey((0, 0))(
      //第一个tuple为给的初始值zeroValue， y为RDD的value
      (tuple, v) => {
        //将v相加， 次数累加， 并保存到tuple。
        (tuple._1 + v, tuple._2 + 1)
      },

      //分区间相加
      (t1, t2) => {
        (t1._1 + t2._1, t1._2 + t2._2)
      }
    )
    println(aggRdd2.collect().mkString("  "))

    val avgRes = aggRdd2.mapValues(
      value => {
        value._1 / value._2
      }
    )
    println(avgRes.collect().mkString("    "))


    //关闭
    sc.stop()
  }

}
