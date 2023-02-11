package com.oys.bd.commonfriends;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

import java.io.File;
import java.io.IOException;

public class CFStep2Driver {
    public static void main(String[] args) throws IOException, InterruptedException, ClassNotFoundException {

        //1. 获取job
        Configuration conf = new Configuration();
        Job job = Job.getInstance(conf);
        
        //2. 获取jar包路径
        job.setJarByClass(CFStep2Driver.class);

        //3. 关联 mapper reducer
        job.setMapperClass(CFStep2Mapper.class);
        job.setReducerClass(CFStep2Reducer.class);

        //4. 设置map输出 k v 类型
       job.setMapOutputKeyClass(Text.class);
       job.setMapOutputValueClass(Text.class);

        //5.设置最终输出的k v 类型
        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(Text.class);

        //6. 设置输入路径和数出路径

        FileInputFormat.setInputPaths(job, new Path("/Users/ouyushun/mr/common_friends/out1/part-r-00000"));
        FileOutputFormat.setOutputPath(job, new Path("/Users/ouyushun/mr/common_friends/out2/"));

        //7. 提交job
        boolean res = job.waitForCompletion(true);
        System.exit(res ? 0 : 1);
    }


    /**
     * 删除文件或文件夹
     */
    public static void deleteIfExists(File file) throws IOException {
        if (file.exists()) {
            if (file.isFile()) {
                if (!file.delete()) {
                    throw new IOException("Delete file failure,path:" + file.getAbsolutePath());
                }
            } else {
                File[] files = file.listFiles();
                if (files != null && files.length > 0) {
                    for (File temp : files) {
                        deleteIfExists(temp);
                    }
                }
                if (!file.delete()) {
                    throw new IOException("Delete file failure,path:" + file.getAbsolutePath());
                }
            }
        }
    }

}
