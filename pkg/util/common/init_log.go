package common

import (
	log "github.com/cihub/seelog"
	"github.com/yin-zt/go-bind-mp/pkg/config"
	"os"
)

func InitLog() {
	os.MkdirAll("log", 07777)

	logger, err := log.LoggerFromConfigAsBytes([]byte(config.Conf.Logs.LogConfigStr))
	if err != nil {
		log.Error(err)
		panic("init log fail")
	}
	log.ReplaceLogger(logger)
	log.Info("log工具配置完成")
}
