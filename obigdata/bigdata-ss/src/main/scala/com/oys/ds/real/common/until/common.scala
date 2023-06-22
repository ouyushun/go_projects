package com.oys.ds.real.common.until

//包对象
package object common {
	abstract class BaseException(message: String) extends Exception(message) {
		def this(message: String, underlying: Exception) {
			//直接调用主构造方法
			this(s"$message . reason : $underlying")
		}
	}

	case class TestException(message: String) extends BaseException(message: String)
}
