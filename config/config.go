package config

import (
	"../utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Token         string `yaml:"token"`
	SynUrl        string `yaml:"synUrl"`
	ListenAddress string `yaml:"listenAddress"`
	FileBedPath   string `yaml:"fileBedPath"`
}

const configFilePath string = "config.yml"

var log = logrus.New()
var config = Config{"token", "http://127.0.0.1:8880", "0.0.0.0:8880", "file_bed"}

func init() {
	log.Info("加载配置文件")
	existAndIsFile, _ := utils.ExistAndIsFile(configFilePath)
	if !existAndIsFile {
		writer, err := utils.CreateFile(configFilePath)
		if err != nil {
			log.WithFields(logrus.Fields{"configFilePath": configFilePath, "err": err}).Panic("配置文件不存在，但创建配置文件失败")
			return
		}
		defer writer.Close()
		bytes, err := yaml.Marshal(&config)
		if err != nil {
			log.WithFields(logrus.Fields{"config": config, "err": err}).Panic("创建配置文件初始配置失败")
			return
		}
		writer.Write(bytes)
		log.Info("成功创建初始化配置文件，请填写配置文件后再启动")
		os.Exit(0)
	}
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.WithFields(logrus.Fields{"configFilePath": configFilePath, "err": err}).Panic("读取配置文件失败")
		return
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.WithFields(logrus.Fields{"configFilePath": configFilePath, "err": err}).Panic("解析配置文件失败")
		return
	}
	log.Info("加载配置文件成功")
}

func GetConfig() Config {
	return config
}
