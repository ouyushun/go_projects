package com.oys.bigdata.spark.core.CasePractice.user_analysis

import org.apache.spark.rdd.RDD
import org.apache.spark.util.AccumulatorV2
import org.apache.spark.{SparkConf, SparkContext}

import scala.collection.mutable

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object HotCategoryTop10Analysis_Accumulator {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("Hot Category Top10").setMaster("local")
    val sc = new SparkContext(sparkConf)

    //1. 读取全部数据
    val lines: RDD[String] = sc.textFile("tmpdata/user_action.csv")

    val acc = new HotCategoryAccumulator
    sc.register(acc, "top10")

    lines.foreach(
      action => {
        val datas = action.split("_")
        if (datas(6) != "-1") {
          // 点击的场合
          acc.add((datas(6), "click"))
        } else if (datas(8) != "null") {
          // 下单的场合
          val ids = datas(8).split(",")
          ids.foreach(
            id => {
              acc.add( (id, "order") )
            }
          )
        } else if (datas(10) != "null") {
          // 支付的场合
          val ids = datas(10).split(",")
          ids.foreach(
            id => {
              acc.add( (id, "pay") )
            }
          )
        }
      }
    )

    val accVal: mutable.Map[String, HotCategory] = acc.value
    val categories: mutable.Iterable[HotCategory] = accVal.map(_._2)
    val sort = categories.toList.sortWith(
      (left, right) => {
        if ( left.clickCount > right.clickCount ) {
          true
        } else if (left.clickCount == right.clickCount) {
          if ( left.orderCount > right.orderCount ) {
            true
          } else if (left.orderCount == right.orderCount) {
            left.payCount > right.payCount
          } else {
            false
          }
        } else {
          false
        }
      }
    )

    // 5. 将结果采集到控制台打印出来
    sort.take(10).foreach(println)
    sc.stop()
  }


}


case class HotCategory(categoryId: String, var clickCount: Int, var orderCount: Int, var payCount: Int)

/**
 * 自定义累加器
 * 1. 继承AccumulatorV2，定义泛型
 *    IN : ( 品类ID, 行为类型 )
 *    OUT : mutable.Map[String, HotCategory]
 * 2. 重写方法（6）
 */
class HotCategoryAccumulator extends AccumulatorV2[(String, String), mutable.Map[String, HotCategory]] {
  private val hotCategoryResultMap = mutable.Map[String, HotCategory]()

  override def isZero: Boolean = {
    this.hotCategoryResultMap.isEmpty
  }

  override def copy(): HotCategoryAccumulator = {
    new HotCategoryAccumulator
  }

  override def reset(): Unit = {
    this.hotCategoryResultMap.clear()
  }

  override def add(v: (String, String)): Unit = {
    val category = v._1
    val actionType = v._2
    val hotCategory: HotCategory = this.hotCategoryResultMap.getOrElse(category, new HotCategory(category, 0,0,0))

    if (actionType == "click") {
      hotCategory.clickCount += 1
    } else if (actionType == "order") {
      hotCategory.orderCount += 1
    } else if (actionType == "pay") {
      hotCategory.payCount += 1
    }

    this.hotCategoryResultMap.update(category, hotCategory)
  }

  override def merge(other: AccumulatorV2[(String, String), mutable.Map[String, HotCategory]]): Unit = {
    val map1 = this.hotCategoryResultMap
    val map2 = other.value
    map2.foreach{
      case ( cid, hc ) => {
        val category: HotCategory = map1.getOrElse(cid, HotCategory(cid, 0, 0, 0))
        category.clickCount += hc.clickCount
        category.orderCount += hc.orderCount
        category.payCount += hc.payCount
        map1.update(cid, category)
      }
    }
  }

  override def value: mutable.Map[String, HotCategory] = this.hotCategoryResultMap
}
