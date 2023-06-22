package com.oys.ds.real.common


package object until {
	abstract class BaseException(message : String) extends Exception(message) {
		def this(message:String, underlying:Exception) = {
			this(s"message : $message, reason: $underlying")
		}
	}

	case class TestException(message:String) extends BaseException(message:String)
	case class UnknownEnumException(msg: String) extends BaseException(msg)
	case class UnknownHDFSFailureException(msg: String) extends BaseException(msg)
}
