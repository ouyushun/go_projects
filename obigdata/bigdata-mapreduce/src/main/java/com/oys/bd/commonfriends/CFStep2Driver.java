package com.oys.bd.commonfriends;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/30
 * @Version 1.0
 */
public class CFStep1Driver {
    public static void main(String[] args) throws IOException, InterruptedException, ClassNotFoundException {

        //1. 获取job
        Configuration conf = new Configuration();
        Job job = Job.getInstance(conf);
        
        //2. 获取jar包路径
        job.setJarByClass(CFStep1Driver.class);

        //3. 关联 mapper reducer
        job.setMapperClass(CFStep1Mapper.class);
        job.setReducerClass(CFStep1Reducer.class);

        //4. 设置map输出 k v 类型
       job.setMapOutputKeyClass(Text.class);
       job.setMapOutputValueClass(Text.class);

        //5.设置最终输出的k v 类型
        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(Text.class);

        //6. 设置输入路径和数出路径

        FileInputFormat.setInputPaths(job, new Path("/Users/ouyushun/mr/common_friends/in.txt"));
        FileOutputFormat.setOutputPath(job, new Path("/Users/ouyushun/mr/common_friends/out"));

        //7. 提交job
        boolean res = job.waitForCompletion(true);
        System.exit(res ? 0 : 1);
    }
}
