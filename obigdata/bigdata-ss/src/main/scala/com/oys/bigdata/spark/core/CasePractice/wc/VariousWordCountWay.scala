package com.oys.bigdata.spark.core.CasePractice.wc

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object VariousWordCountWay {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")
    val sc = new SparkContext(sparkConf)

    //操作
    val lines: RDD[String] = sc.textFile("./wc.txt")

    val words: RDD[String] = lines.flatMap(_.split(" "))

    //1.groupBy
    //效率不高， 因为有shuffle过程
    val group = words.groupBy(word => word)
    val wordCount = group.mapValues(iter => iter.size)

    //2. groupByKey, 效率不高， 因为有shuffle过程
    words.map((_, 1)).groupByKey().mapValues(_.size)

    //3. reduceByKey
    words.map((_,1)).reduceByKey(_+_)

    //4. aggregateByKey
     words.map((_, 1)).aggregateByKey(0)(_+_, _+_)

    //5. foldByKey, 分区内和分区间函数相同， 可以简化aggregateByKey
    words.map((_, 1)).foldByKey(0)(_+_)


    //6. combineByKey
    words.map((_, 1)) combineByKey(
      v => v,
      (_: Int) + (_: Int),
      (_: Int) + (_: Int)
    )

    //7.
    words.countByValue()
    //8
    words.map((_, 1)).countByKey()





    //关闭
    sc.stop()
  }

}
