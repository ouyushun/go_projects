package com.oys.ds.highfunc

/**
 * @Author ouyushun
 * @Date 2023/1/7
 * @Version 1.0
 */
object Calcutor {
	def main(args: Array[String]): Unit = {
		def calcute(x: Int, y : Int, operate: (Int, Int) => Int) : Int = {
			operate(x, y)
		}
		val res = calcute(10, 20, (x: Int, y: Int)  => {x + y})
		println(res)
	}
}
