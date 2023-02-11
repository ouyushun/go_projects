package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/31
 * @Version 1.0
 */
object MatchCaseLearn {
	def main(args: Array[String]): Unit = {
		def patternType(obj2: Any)  = obj2 match {
			case  x:  Int  =>  x
			case  s:  String  =>  s
			case  _:  BigInt  =>  Int.MaxValue
			case  _  =>  0
		}
		println(patternType(10))
		println(patternType("13"))
		println(patternType(BigInt.apply(10)))
		println(patternType(BigDecimal.apply(10)))

		println("-----------------------")

		def patternArray(obj: Array[Any]) = obj match {
			// 匹配包含 0 的数组
			case Array(0) => "0"
			case Array(999) => "999"
			// 匹配任意两个元素的数组，并将这两个元素分别绑定到 x 和 y 变量
			case Array(x, y) => x + " <-> " + y
			// 匹配以 0 开始的数组
			case Array(0, _*) => "0 ..."
			case _ => "something else"
		}
		println(patternArray(Array(0)))
		println(patternArray(Array(999,1,1,2)))
		println(patternArray(Array(1, 2)))
		println(patternArray(Array(0, 2, 4)))
		println(patternArray(Array(2, 4, 6)))
	}

}
