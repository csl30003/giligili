package main

import (
	"flag"
	"fmt"

	"giligili/app/video/api/internal/config"
	"giligili/app/video/api/internal/handler"
	"giligili/app/video/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

const maxMemory = 32 << 20 // 32MB

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 因为要传视频，所以调大了
	c.MaxBytes = maxMemory

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
