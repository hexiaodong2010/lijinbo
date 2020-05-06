package config

import (
	"fmt"
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
	_, err := toml.DecodeFile(fmt.Sprintf("config/config-%s.toml", runEnv), &cfg)
	if err!=nil {
		log.Fatal(err)
	}
	RunMode = runEnv
	log.Println(cfg)
}
