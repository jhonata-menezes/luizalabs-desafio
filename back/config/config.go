package config

import (
	"github.com/crgimenes/goconfig"
	"log"
)

type Postgres struct {
	Host string `yaml:"host" cfgDefault:"127.0.0.1"`
	Port int `yaml:"port" cfgDefault:"5432"`
	Database string `yaml:"database" cfgDefault:"luizalabs"`
	Username string `yaml:"username" cfgDefault:"luizalabs"`
	Password string `yaml:"password" cfgDefault:"luizalabs"`
	PoolSize int `yaml:"poolsize" cfgDefault:"30"`
}

type Config struct {
	Postgres Postgres `yaml:"postgres"`
}

var configData Config

func GetConfig() Config {
	return configData
}

func init() {
	goconfig.File = "config.yml"
	err := goconfig.Parse(&configData)
	if err != nil {
		log.Println("arquivo de configuracao nao encontrado", err)
	}

}
