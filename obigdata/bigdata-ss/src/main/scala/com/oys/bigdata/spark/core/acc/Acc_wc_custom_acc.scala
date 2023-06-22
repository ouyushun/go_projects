package com.oys.bigdata.spark.core.acc

import org.apache.spark.util.{AccumulatorV2, LongAccumulator}
import org.apache.spark.{SparkConf, SparkContext}

import scala.collection.mutable

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Acc_wc_custom_acc {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("ACC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd = sc.makeRDD(
      List("hello", "spark", "hello", "scala"), 4
    )

    val wcAcc = new CustonAccumulator
    sc.register(wcAcc, "word count")

    rdd.foreach(
      word => {
        //自定义累加器
        wcAcc.add(word)
      }
    )

    println(wcAcc.value)

    //关闭
    sc.stop()
  }

}

/*
* 1. 继承AccumulatorV2， 定义泛型
*   IN:输入类型
*   OUT:输出类型
* */
class CustonAccumulator extends AccumulatorV2[String, mutable.Map[String, Long]] {
  private val wcMap =  mutable.Map[String, Long]()

  //判断是否为初始状态
  override def isZero: Boolean = {
    wcMap.isEmpty
  }

  override def copy(): AccumulatorV2[String,  mutable.Map[String, Long]] = {
    new CustonAccumulator()
  }

  override def reset(): Unit = {
    wcMap.clear()
  }

  override def add(word: String): Unit = {
    val num = wcMap.getOrElse(word, 0L)
    wcMap.update(word, num + 1)
  }

  //Driver合并多个累加器
  override def merge(other: AccumulatorV2[String, mutable.Map[String, Long]]): Unit = {
    val m1 = wcMap
    val m2 = other.value
    m2.foreach {
      case (word, count) => {
        m1.update(word, m1.getOrElse(word, 0L) + count)
      }
    }
  }

  override def value: mutable.Map[String, Long] = {
    wcMap
  }
}
