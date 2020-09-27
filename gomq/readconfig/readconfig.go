package readconfig

import (
	"log"

	"github.com/Unknwon/goconfig"
)

// GetConfigFile  is a public funtion
func GetConfigFile(filePath string) (configFile *goconfig.ConfigFile) {
	cfg, err := goconfig.LoadConfigFile(filePath)
	if err != nil {
		log.Fatal("can not load config,error:", err)
	}
	return cfg
}
