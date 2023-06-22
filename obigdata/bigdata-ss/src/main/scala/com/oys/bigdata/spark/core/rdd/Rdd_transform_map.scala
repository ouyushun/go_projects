package com.oys.bigdata.spark.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_map {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    //分区内数据的执行是有序的
    //不同分区内数据的执行是无序的
    val rdd : RDD[Int] = sc.makeRDD(List(1, 2, 3, 4, 5, 6), 8)
    val mapRdd1 = rdd.map(
      num => {
        println(">>>>>>>>", num)
        num * 2
      }
    )
    val mapRdd2 = mapRdd1.map(
      num => {
        println("########", num.getClass.getSimpleName)
        num * 2
      }
    )


    val rdd3 = mapRdd2.mapPartitions(
      iter => {
        iter.map(_*2)
      }
    )

    val rdd4 = mapRdd2.mapPartitions(
      num => {
        num.map(_*2)
      }
    )

    //按照分区过滤
    val rdd5 = rdd4.mapPartitionsWithIndex(
      (index, iter) => {
        //index: 索引编号
        // 只保留第二个分区的数据。
        if (index == 1) {
          iter
        } else {
          //空的迭代器
          Nil.iterator
        }
      }
    )

    //显示数据所在分区
    val rdd6 = rdd4.mapPartitionsWithIndex(
      (index, iter) => {
        iter.map(
          num => {
            (index, num)
          }
        )
      }
    )

    rdd6.collect().foreach(println)
    //关闭
    sc.stop()
  }

}
