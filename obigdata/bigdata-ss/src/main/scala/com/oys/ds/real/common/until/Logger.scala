package com.oys.ds.real.common.until

import java.time.LocalDateTime
import java.time.format.DateTimeFormatter
object Logger {
	object Level extends EnumList[Int, Level] {
		val info: Level = add(Level(0, "info"))
		val warn: Level = add(Level(1, "WARN"))
		val error: Level = add(Level(2, "ERROR"))
	}

	case class Level(get: Int, name: String) extends Enum[Int]

	/** Log level. Only logs with level equal or above this will be printed. */
	var logLevel: Level = Level.info

	def info(any: Any*): Unit = log(Level.info, any: _*)

	def warn(any: Any*): Unit = log(Level.warn, any: _*)

	def error(any: Any*): Unit = log(Level.error, any: _*)

	private def log(level: Level, any: Any*): Unit = {
		if (level.get >= logLevel.get) {
			val datetimeFormatter = DateTimeFormatter.ofPattern("YYYY-MM-dd HH:mm:ss")
			val datetimeString = LocalDateTime.now().format(datetimeFormatter)

			val message = any.map(Option(_).map(_.toString).orNull).mkString(", ")
			println(s"$datetimeString | ${level.name} | $message")
		}
	}
}
