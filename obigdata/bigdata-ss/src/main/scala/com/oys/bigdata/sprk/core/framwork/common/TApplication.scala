package com.oys.bigdata.sprk.core.framwork.common


import com.oys.bigdata.sprk.core.framwork.util.EnvUtil
import org.apache.spark.{SparkConf, SparkContext}

trait TApplication {

    def start(master:String ="local[*]", app:String = "Application")( op : => Unit ): Unit = {
        val sparConf = new SparkConf().setMaster(master).setAppName(app)
        val sc = new SparkContext(sparConf)
        EnvUtil.put(sc)

        try {
            op
        } catch {
            case ex: Throwable => println(ex.getMessage)
        }

        // TODO 关闭连接
        sc.stop()
        EnvUtil.clear()
    }
}
