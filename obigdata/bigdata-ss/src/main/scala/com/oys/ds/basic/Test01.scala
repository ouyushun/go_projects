package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/24
 * @Version 1.0
 */
object Test01 {
	def main(args: Array[String]): Unit = {
		var bbb:BBB = new BBB
		var res = bbb.cal(3,5)
		println(res)
	}
}

class BBB extends AAA {

}

class AAA {
	def cal(x:Int,y:Int)= {
		(x+y)/(x-y)
	}
}