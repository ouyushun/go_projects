package com.oys.bigdata.log.utils;

import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;

/**
 * 上传日志数据工具类
 */
public class UploadUtils {

    public static void upload(String path, String log) {

        try {
            URL url = new URL(path);
            System.out.println(url);
            HttpURLConnection conn = (HttpURLConnection)url.openConnection();

            // 设置请求方式为post
            conn.setRequestMethod("POST");

            conn.setDoOutput(true);

            conn.setRequestProperty("Content-Type", "application/json");

            OutputStream out = conn.getOutputStream();
            out.write(log.getBytes());
            out.flush();
            out.close();

            int code = conn.getResponseCode();
            System.out.println(code);

        } catch (Exception e) {
            e.printStackTrace();
        }


    }
}
