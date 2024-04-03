package conf

import (
	"bufio"
	"encoding/json"

	"os"

	"github.com/JustACP/AutoDDNS/config"
	"github.com/JustACP/AutoDDNS/logging"
)

var configPath = "./config.json"

func LoadConfig() *config.Config {

	var Conf config.Config

	_, err := os.Stat(configPath)
	if err != nil {
		logging.Info("config file not exist!, will create config file")
		createConfig()
		os.Exit(1)
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("read config error, %v", err)
	}

	err = json.Unmarshal([]byte(configFile), &Conf)
	if err != nil {
		logging.Error("unmarshal config error, %v", err)
	}

	logging.Info("load config file success")
	return &Conf
}

func createConfig() {
	configFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		logging.Error("create config file error %v", err)
	}

	writer := bufio.NewWriter(configFile)

	confJSON, err := json.MarshalIndent(config.GetTemplateConfig(), " ", "  ")

	if err != nil {
		logging.Error("marshal config error, %v", err)
	}

	writer.Write([]byte(confJSON))
	writer.Flush()

}
