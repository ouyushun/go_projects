package com.oys.bd.flowStat;

import com.oys.bd.flowStat.writable.FlowBean;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/5
 * @Version 1.0
 */
public class FlowMapper extends Mapper<LongWritable, Text, Text, FlowBean> {

    private Text outK = new Text();
    private FlowBean outV = new FlowBean();

    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, FlowBean>.Context context) throws IOException, InterruptedException {
        //获取一行
        String line = value.toString();
        String[] splite = line.split("\t");

        String phone = splite[0];
        String upFlow = splite[1];
        String downFlow = splite[2];

        //封装
        outK.set(phone);
        outV.setUpFlow(Long.parseLong(upFlow));
        outV.setDownFlow(Long.parseLong(downFlow));
        outV.setSumFlow();

        //写出
        context.write(outK, outV);
    }
}
