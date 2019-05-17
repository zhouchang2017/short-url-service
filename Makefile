go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build

build:
# 一定要注意 Makefile 中的缩进，否则执行 make build 可能报错：make: Nothing to be done for `build'.
# protoc 命令前边是一个 Tab，不是四个空格
	# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
	GOOS=linux GOARCH=amd64 go build
	# 根据当前目录下的 Dockerfile 生成名为 t.wewee 的镜像
	docker build -t zhouchang2018/shor-url-service .


run:
	docker run -p 8000:8000 -e DB_USERNAME=root -e DB_PASSWORD=12345678 -e DB_HOST=127.0.0.1 -e DB_PORT=3306 -e DB_NAME=micro_book_mall zhouchang2018/shor-url-service
