package com.oys.bd.udf;

import org.apache.commons.lang.StringUtils;
import org.apache.hadoop.hive.ql.exec.UDF;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @Author ouyushun
 * @Date 2022/10/19
 * @Version 1.0
 */
public class EncryptNumber extends UDF {

    public String evaluate(String phone) {
        String encryptNum = "";
        if (StringUtils.isNotEmpty(phone) && phone.trim().length() == 11) {
            String regex = "^(1[3-9]\\d{9}$)";
            Pattern p = Pattern.compile(regex);
            Matcher m = p.matcher(phone);
            if (m.matches()) {
                encryptNum = phone.trim().replaceAll("(\\d{3})\\d{4}(\\d{4})", "$1****$2");
            } else {
                encryptNum = phone;
            }
        }
        return encryptNum;
    }
}
