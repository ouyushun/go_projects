package com.oys.ds.highfunc

import scala.math.ceil

/**
 * @Author ouyushun
 * @Date 2023/1/4
 * @Version 1.0
 */

object fun0  {
	def main(args: Array[String]): Unit = {
		val num = 3.14
		val f1 = ceil(4.3)
		val arr = Array(1.0,2.5,3.3,4,5)
		val b = arr.map(ceil)
		println(b.mkString)


		//声明的函数
		def plus(x:Int) : Int = 666 + x
		val arr2 = Array(1,2,3,4,5,6,7,8)
		val res = arr2.map(plus)
		println(res.mkString(","))
	}
}

object Func2 {
	def main(args: Array[String]): Unit = {
		//函数作为返回值
		def minus(x:Int) = {
			//匿名函数
			(y : Int) => x - y
		}
		println(minus(10)(65))

		def test():Unit = {
			println("aaaaa------")
		}
		
		this.test()
	}

	def test():Unit = {
		println("bbbb-----")

	}

	def test(v:Int):Unit = {
		println("bbbb")

	}
}
