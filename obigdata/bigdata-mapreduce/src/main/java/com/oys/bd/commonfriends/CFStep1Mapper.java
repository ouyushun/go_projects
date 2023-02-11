package com.oys.bd.commonfriends;

import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/30
 * @Version 1.0
 */
/*
* KEYIN,  mapper的输入, 行首的偏移量。 longWriteble
* VALUEIN,
* KEYOUT,
* VALUEOUT
* */
public class CFStep1Mapper extends Mapper<LongWritable, Text, Text, Text> {
    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        String line = value.toString();
        String user = line.split(":")[0];
        String friendsStr = line.split(":")[1];
        String[] friends = friendsStr.split(",");
        //输入： A: B,C,D,E
        //输出: <B, A> <C, A> <D, A> <E, A>
        for  (String friend : friends) {
            context.write(new Text(friend), new Text(user));
        }
    }
}
