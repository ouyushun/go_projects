#targets : prerequisites
#	command


x.txt: m.txt c.txt
	cat m.txt c.txt > x.txt

m.txt: a.txt b.txt
	cat a.txt b.txt > m.txt



#clean规则与我们前面编写的规则有所不同，它没有依赖文件，因此，要执行clean，必须用命令make clean：
.PHONY: clean
clean:
	rm x.txt && rm m.txt


cd:
	#pwd ;cd ~ ;pwd
	#另一种执行多条命令的语法是用&&，它的好处是当某条命令失败时，后续命令不会继续执行：
	#cd .. && pwd
	pwd; \
	cd ~ ; \
	pwd ;

echo:
	@echo "aaaaaaaaa";
	echo "bbbbbbbbb";