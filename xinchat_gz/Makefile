
release-test: user-rpc-dev

#编译
#build镜像
#推送到镜像仓库
user-rpc-dev:
	make -f deploy/mk/user-rpc.mk release-test


#从镜像仓库pull镜像
#docker run镜像 开启服务
install-server:
	cd ./deploy/script && chmod +x restart-all-srv.sh && ./restart-all-srv.sh



