package com.oys.ds.collection

/**
 * @Author ouyushun
 * @Date 2023/1/10
 * @Version 1.0
 */
object ImmutableColl {
	def main(args: Array[String]): Unit = {
		//immutable.Map
		val m1 = Map("a" -> 10, "b" -> 14)
		println(m1)
		var seq1 = Seq(1,2,3,4)
		println(seq1)
	}

}
