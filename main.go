package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/go-ddd-hertz/boot"
	"github.com/jettjia/go-ddd-hertz/interfaces/event"
	"github.com/jettjia/go-ddd-hertz/interfaces/grpc"
	"github.com/jettjia/go-ddd-hertz/interfaces/http"
	"github.com/jettjia/go-ddd-hertz/interfaces/job"
)

func main() {
	ENV := flag.String("env", "debug", "环境,配置读取")
	flag.Parse()

	// 全局配置
	boot.InitServer(*ENV)

	// 依赖注入
	app := boot.InitApp()

	// start http
	http.InitHttp(app)

	// start grpc
	grpc.InitGrpc(app)

	// start event mq
	event.InitEvent(app)

	// start InitJob
	job.InitJob(app, *ENV)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
