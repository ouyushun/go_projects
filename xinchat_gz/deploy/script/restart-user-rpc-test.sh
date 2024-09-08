#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/gim-ou/userdd-rpc-dev'
tag='latest'

container_name="gim-user-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/userdd-rpc:/userdd/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8090:8080  --name=${container_name} -d ${reso_addr}:${tag}