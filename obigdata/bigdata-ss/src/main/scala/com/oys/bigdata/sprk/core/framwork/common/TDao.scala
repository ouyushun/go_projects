package com.oys.bigdata.sprk.core.framwork.common

import com.oys.bigdata.sprk.core.framwork.util.EnvUtil

trait TDao {

    def readFile(path:String) = {
        EnvUtil.take().textFile(path)
    }
}
