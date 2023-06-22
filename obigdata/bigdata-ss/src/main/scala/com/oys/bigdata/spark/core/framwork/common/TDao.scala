package com.oys.bigdata.spark.core.framwork.common

import com.oys.bigdata.spark.core.framwork.util.EnvUtil

trait TDao {

    def readFile(path:String) = {
        EnvUtil.take().textFile(path)
    }
}
