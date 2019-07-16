package postgres

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/jhonata-menezes/luizalabs-desafio/back/config"
)

var conn *pg.DB
var connected = false

func init() {
	GetConnection()
}

func GetConnection() *pg.DB {
	if connected {
		connected = true
		return conn
	}
	config := config.GetConfig().Postgres
	conn = pg.Connect(&pg.Options{
		User: config.Username,
		Password: config.Password,
		Database: config.Database,
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		PoolSize: config.PoolSize,
	})
	return conn
}
