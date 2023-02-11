package com.oys.ds.real.common.until

object test extends App {
	def testExceotionF(num: Int): Unit = {
		if (num < 0) {
			throw TestException("num is below zero hhh")
		}
	}

	testExceotionF(-1)
}
