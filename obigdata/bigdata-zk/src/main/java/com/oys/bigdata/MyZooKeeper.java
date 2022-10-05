package com.oys.bigdata;

import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.ZooKeeper;

import java.security.PrivateKey;

/**
 * @Author ouyushun
 * @Date 2022/9/22
 * @Version 1.0
 */
public class MyZooKeeper {

    private static  String connectStr = "127.0.0.1:2181";
    private static ZooKeeper zooKeeper;

    public static void main(String[] args) throws Exception {
        zkInit();
    }

    public static void zkInit() throws Exception{
        zooKeeper = new ZooKeeper(
                connectStr,
                2000,
                new Watcher() {
                    @Override
                    public void process(WatchedEvent watchedEvent) {

                    }
                }
        );

        System.out.println(zooKeeper);
    }
}
