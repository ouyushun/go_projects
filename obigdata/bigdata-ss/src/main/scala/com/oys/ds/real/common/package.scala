package com.oys.ds.real

/**
 * @Author ouyushun
 * @Date 2023/2/14
 * @Version 1.0
 */
/*
* 包对象， 提供整个包共享使用的容器
* */
package object common {

	abstract class BaseException(message : String) extends Exception(message) {
		def this(message:String, underlying:Exception) = {
			this(s"message : $message, reason: $underlying")
		}
	}

	case class TestException(message:String) extends BaseException(message:String)
	case class UnknownEnumException(msg: String) extends BaseException(msg)
	case class UnknownHDFSFailureException(msg: String) extends BaseException(msg)


}
