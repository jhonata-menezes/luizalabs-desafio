package config

import "testing"

func TestConfig(t *testing.T) {
	c := GetConfig()
	if c.App.Host != ":8080" {
		t.Error("host changed", c.App.Host)
	}

	postgresConfig := c.Postgres
	if postgresConfig.Host != "127.0.0.1" ||
		postgresConfig.Port != 5432 ||
		postgresConfig.Database != "luizalabs" ||
		postgresConfig.Password != "luizalabs" ||
		postgresConfig.Username != "luizalabs" ||
		postgresConfig.PoolSize == 0 {
		t.Error("config not found", postgresConfig)
	}
}
