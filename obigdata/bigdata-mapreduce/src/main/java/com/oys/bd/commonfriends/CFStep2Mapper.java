package com.oys.bd.commonfriends;

import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;

import java.io.IOException;
import java.util.Arrays;

/*
* KEYIN,  mapper的输入, 行首的偏移量。 longWriteble
* VALUEIN,
* KEYOUT,
* VALUEOUT
* */
public class CFStep2Mapper extends Mapper<LongWritable, Text, Text, Text> {
    @Override
    protected void map(LongWritable key, Text value, Mapper<LongWritable, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        String line = value.toString();
        String[] split = line.split("\t");
        //friend是persons这些人的好友
        String friend = split[0];
        String personsStr = split[1];
        String[] persons = personsStr.split(",");
        Arrays.sort(persons);
        //输入 A    B,C, D, E
        //输出： <B-C, A> <B-D, A> <B-E, A> <C-D, A> <C-E, A> <D-E, A>
        Integer len = persons.length;
        for (int i = 0; i < len - 1; i++) {
            for (int j = i + 1; j < len ; j++) {
                context.write(new Text(persons[i] + "-" + persons[j]), new Text(friend));
            }
        }
    }
}
