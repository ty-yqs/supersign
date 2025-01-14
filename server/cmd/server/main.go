package main

import (
	"supersign/internal/server"
	"supersign/pkg"
	"supersign/pkg/ali"
	"supersign/pkg/conf"
	"supersign/pkg/log"
	"supersign/pkg/sign"
	"supersign/pkg/validatorer"

	"github.com/spf13/pflag"
)

var (
	confPath = pflag.StringP("conf", "c", "conf/default.yaml", "指定配置文件路径")
)

func setup() {
	conf.Path = *confPath
	conf.Setup()
	log.Setup(conf.Log.Level)
	validatorer.Setup()
	sign.Setup(log.New("sign").L(), conf.Server.MaxJob)
	ali.Setup()
	conf.OnChange(func() {
		if err := pkg.Reset(); err != nil {
			return
		}
		server.Reset()
		log.New("conf").L().Info("OnChange")
	})
}

// @title supersign 后端服务接口文档
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	pflag.Parse()
	setup()
	server.Start()
}
