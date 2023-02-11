package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_flodByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd  = sc.makeRDD(
      List(("a", 1), ("a", 2), ("a", 3), ("b", 6)), 2
    )

    //存在函数柯里化， 两个参数列表
    //第一个参数列表： 初始值
    //第二个参数：
    //   第一个参数： 分区内计算规则
    //   第二个参数： 分区间计算规则
    val aggRdd = rdd.aggregateByKey(0)(
      (x, y) => math.max(x, y),
      (x, y) => x + y
    )
    aggRdd.collect().foreach(println)


    val aggRdd2 = rdd.aggregateByKey(0)(
      (x, y) => x + y, // _+_
      (x, y) => x + y // _+_
    )

    //如果分区内和分区间计算规则相同，可以简化方法 flodByKey
    val aggRdd3 = rdd.foldByKey(0)(_+_)




    //关闭
    sc.stop()
  }

}
