package com.oys.bigdata.spark.sql.join

import org.apache.spark.sql.SparkSession

/**
 * @Author ouyushun
 * @Date 2023/2/23
 * @Version 1.0
 */
object JoinDemo01 extends App {
	private val spark = SparkSession.builder().master("local[1]").getOrCreate()
	import spark.implicits._

	// 因为我们下面测试数据都很小，所以我们先把 BroadcastJoin 关闭
	spark.conf.set("spark.sql.autoBroadcastJoinThreshold", 1)

	// 为了启用 Shuffle Hash Join 必须将 spark.sql.join.preferSortMergeJoin 设置为 false
	//spark.conf.set("spark.sql.join.preferSortMergeJoin", false)

	val df1 = Seq(
		(0, "a"),
		(1, "b"),
		(2, "c"),
	).toDF("id", "info")

	val df2 = Seq(
		(0, "zhangsan"),
		(1, "lisi"),
		(2, "wangwu")
	).toDF("id", "name")

	val r2 = df1.join(df2, Seq("id"), "inner")

	r2.explain()

}
