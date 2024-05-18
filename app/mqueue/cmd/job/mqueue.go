package main

import (
	"XavierLookLook/app/mqueue/cmd/job/internal/config"
	"XavierLookLook/app/mqueue/cmd/job/internal/logic"
	"XavierLookLook/app/mqueue/cmd/job/internal/svc"
	"context"
	"flag"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	//这里会完成asynq服务端的创建
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	//这里会完成队列和处理函数的绑定
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register() //注册队列和处理函数（或处理接口）

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
