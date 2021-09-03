package ginp

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	CONF_FILE = "application.yaml"
)

type ServerConfig struct {
	Port int32
	Name string
}

type SysConfig struct {
	Server *ServerConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{
		Port: 8081,
		Name: "default",
	}}
}

func InitConfig() *SysConfig {
	conf := NewSysConfig()
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b, conf)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return conf
}

func LoadConfigFile() []byte {
	conf, err := ioutil.ReadFile(CONF_FILE)
	if err != nil {
		return nil
	}
	return conf
}
