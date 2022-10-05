package com.oys.bd.wc;

import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.io.Writable;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/4
 * @Version 1.0
 */
/*
* KEYIN,  mapper的输入, 行首的偏移量。 longWriteble
* VALUEIN,
* KEYOUT,
* VALUEOUT
* */
public class MyWCMapper extends Mapper<LongWritable, Text, Text, IntWritable> {

    private Text outK = new Text();
    private IntWritable outV = new IntWritable();


    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, IntWritable>.Context context) throws IOException, InterruptedException {


        String line = value.toString();
        String[] words = line.split(" ");
        for  (String word:words) {
            outK.set(word);
            outV.set(1);
            context.write(outK, outV);
        }
    }
}
