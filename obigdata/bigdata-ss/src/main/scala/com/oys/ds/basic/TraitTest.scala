package com.oys.ds.basic


/**
 * @Author ouyushun
 * @Date 2022/12/30
 * @Version 1.0
 */
object TraitTest {
	def main(args: Array[String]): Unit = {
		val user = new User
		user.addUser()
		user.updateUser1()
	}

	trait MyTrait_1 {
		println("trait start...")
		def updateUser1() : Unit = {
			println("update User")
		}
	}

	trait MyTrait_2 {
		println("trait start...")
		def updateUser2() : Unit = {
			println("update User")
		}
	}

	class Parent {

	}

	class User extends Parent with  MyTrait_1 with MyTrait_2{
		def addUser() : Unit = {
			println("add User")
		}
	}
}
