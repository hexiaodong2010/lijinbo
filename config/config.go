package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

const (
	ProEnv  = "prod"
	TestEnv = "test"
	DevEnv  = "dev"
)

var (
	RunMode  = DevEnv
	MysqlDsn string
)
type conf struct {
	MysqlDsn string
	Redis   struct {
		Auth string
		Host string
		Port int
	}
}

func  InitConfig(runEnv string) {
	cfg := conf{}
	_, err := toml.DecodeFile("config/config-dev.toml", &cfg)
	if err!=nil {
		log.Println(err)
		return
	}
	RunMode = runEnv
	log.Println(cfg)
}
