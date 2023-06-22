package com.oys.bigdata.spark.sql.join

import org.apache.spark.sql.SparkSession

/**
 * @Author ouyushun
 * @Date 2023/2/23
 * @Version 1.0
 */
object SortMergeJoinDemo01 extends App {
	private val spark = SparkSession.builder().master("local[1]").getOrCreate()
	import spark.implicits._


	/*
	* 仔细分析的话会发现，sort-merge join的代价并不比shuffle hash join小，反而是多了很多。
	* 那为什么SparkSQL还会在两张大表的场景下选择使用sort-merge join算法呢？
	* 这和Spark的shuffle实现有关，目前spark的shuffle实现都适用sort-based shuffle算法，
	* 因此在经过shuffle之后partition数据都是按照key排序的。
	* 因此理论上可以认为数据经过shuffle之后是不需要sort的，可以直接merge。
	* */

	/*
	*
	* 经过上文的分析，可以明确每种Join算法都有自己的适用场景，数据仓库设计时最好避免大表与大表的join查询，SparkSQL也可以根据内存资源、带宽资源适量将参数spark.sql.autoBroadcastJoinThreshold调大，让更多join实际执行为broadcast hash join。
	* */

	/*
	* 条件与特点
仅支持等值连接
支持所有join类型
Join Keys是排序的
参数spark.sql.join.prefersortmergeJoin (默认true)设定为true
	* */

	// 因为我们下面测试数据都很小，所以我们先把 BroadcastJoin 关闭
	spark.conf.set("spark.sql.autoBroadcastJoinThreshold", 1)

	// 为了启用 Shuffle Hash Join 必须将 spark.sql.join.preferSortMergeJoin 设置为 false
	spark.conf.set("spark.sql.join.preferSortMergeJoin", false)

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

/*
* +- Project [id#7, info#8, name#19]
   +- SortMergeJoin [id#7], [id#18], Inner
      :- Sort [id#7 ASC NULLS FIRST], false, 0
      :  +- Exchange hashpartitioning(id#7, 200), ENSURE_REQUIREMENTS, [plan_id=17]
      :     +- LocalTableScan [id#7, info#8]
      +- Sort [id#18 ASC NULLS FIRST], false, 0
         +- Exchange hashpartitioning(id#18, 200), ENSURE_REQUIREMENTS, [plan_id=18]
            +- LocalTableScan [id#18, name#19]
+- Project [id#7, info#8, name#19]
   +- SortMergeJoin [id#7], [id#18], Inner
      :- Sort [id#7 ASC NULLS FIRST], false, 0
      :  +- Exchange hashpartitioning(id#7, 200), ENSURE_REQUIREMENTS, [plan_id=17]
      :     +- LocalTableScan [id#7, info#8]
      +- Sort [id#18 ASC NULLS FIRST], false, 0
         +- Exchange hashpartitioning(id#18, 200), ENSURE_REQUIREMENTS, [plan_id=18]
            +- LocalTableScan [id#18, name#19]


* */