package com.oys.bigdata.spark.sparkOpt

import org.apache.spark.SparkConf
import org.apache.spark.sql.{Row, SparkSession}

/**
 * @Author ouyushun
 * @Date 2023/3/5
 * @Version 1.0
 */
object Opt01 extends App {
	val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-opt01")
	val spark = SparkSession.builder().config(sc).getOrCreate()

	val dataExtractFields: Iterator[Row] => Iterator[(String, Long)] = {
		(rows: Iterator[Row]) => {
			var fields = Seq[(String, Long)]()
			rows.foreach(row => {
				if (!row.isNullAt(row.fieldIndex("user_id")) && !row.isNullAt(row.fieldIndex("server_id"))) {
					fields = fields :+ (row.getAs[String]("user_id"), row.getAs[Long]("server_id"))
				}
				else {
					fields = fields :+ ("", 999999L)
				}
			})
			fields.toIterator
		}
	}


	//优化后
	val dataExtractFieldsOpt: Iterator[Row] => Iterator[(String, Long)] = {
		(rows: Iterator[Row]) =>
			rows.map(
				row => {
					if (!row.isNullAt(row.fieldIndex("user_id")) && !row.isNullAt(row.fieldIndex("server_id"))) {
						(row.getAs("user_id"), row.getAs("server_id"))
					} else {
						("", 999999)
					}
				})
	}
}
