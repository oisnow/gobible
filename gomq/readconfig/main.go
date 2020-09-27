package readconfig

import (
	"log"

	"github.com/Unknwon/goconfig"
)

func getConfigFile(filePath string) (configFile *goconfig.ConfigFile) {
	cfg, err := goconfig.LoadConfigFile(filePath)
	if err != nil {
		log.Fatal("can not load config")
	}
	return cfg
}
