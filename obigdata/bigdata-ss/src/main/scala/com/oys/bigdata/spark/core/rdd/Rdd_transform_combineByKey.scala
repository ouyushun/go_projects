package com.oys.bigdata.spark.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_combineByKey {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(
      List(("a", 3), ("a", 6), ("a", 3), ("b", 6), ("b", 4)), 2
    )

    //计算平均值
    //需要三个参数
    //第一个参数， 将相同key对应的数据进行结构的转换

    val aggRdd = rdd.combineByKey(
      //第一个参数， 将(a, 3) 转化为 (a, (3, 1))
      v => (v, 1),
      //第二个参数， 分区被计算
      (tuple:(Int, Int), v) => {
        //将v相加， 次数累加， 并保存到tuple。
        (tuple._1 + v, tuple._2 + 1)
      },
      //第三个参数， 分区间计算
      (t1:(Int, Int), t2:(Int, Int)) => {
        (t1._1 + t2._1, t1._2 + t2._2)
      }
    )

    aggRdd.collect().foreach(println)



    //关闭
    sc.stop()
  }

}
