package com.oys.bd.flowStat;

import com.oys.bd.flowStat.writable.FlowBean;
import com.oys.bd.wc.MyWCMapper;
import com.oys.bd.wc.MyWCReducer;
import com.oys.bd.wc.WCDriver;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

import java.io.IOException;

/**
 * /Users/ouyushun/work/code/goProject/obigdata/bigdata-flume
 * @Author ouyushun
 * @Date 2022/10/5
 * @Version 1.0
 */
public class FlowDriver {
    public static void main(String[] args) throws IOException, InterruptedException, ClassNotFoundException {
        //1. 获取job
        Configuration conf = new Configuration();
        Job job = Job.getInstance(conf);

        //2. 获取jar包路径
        job.setJarByClass(FlowDriver.class);

        //3. 关联 mapper reducer
        job.setMapperClass(FlowMapper.class);
        job.setReducerClass(FlowReducer.class);

        //4. 设置map输出 k v 类型
        job.setMapOutputKeyClass(Text.class);
        job.setMapOutputValueClass(FlowBean.class);

        //5.设置最终输出的k v 类型
        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(FlowBean.class);

        //6. 设置输入路径和数出路径
        FileInputFormat.setInputPaths(job, new Path("/Users/ouyushun/mr/flow/in.txt"));
        FileOutputFormat.setOutputPath(job, new Path("/Users/ouyushun/mr/flow/out"));

        //7. 提交job
        boolean res = job.waitForCompletion(true);
        System.exit(res ? 0 : 1);
    }
}
