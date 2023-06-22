package com.oys.ds.real.common.until

trait Enum[A] extends Serializable {
	val get: A

	override def equals(obj: Any): Boolean = {
		obj match {
			case e: Enum[A] => get == e.get
			case _          => false
		}
	}

	override def hashCode(): Int = get.hashCode()

	override def toString: String = get.toString
}