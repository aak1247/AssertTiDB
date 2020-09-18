package config

import (
	"strconv"
)

type (
	Config struct {
		Env  string `env:"ASSERT_TIDB_ENV" default:"dev"`
		Tidb Tidb
	}
	Tidb struct {
		Host     string `default:"127.0.0.1" env:"ASSERT_TIDB_HOST"`
		Port     int    `default:"27017" env:"ASSERT_TIDB_PORT"`
		Username string `default:"" env:"ASSERT_TIDB_USERNAME"`
		Password string `default:"" env:"ASSERT_TIDB_PASSWORD"`
		Database string `require:"true" default:"" env:"ASSERT_TIDB_DATABASE"`
		URL      string `default:"" env:"ASSERT_TIDB_CONNECT_STRING"`
	}
)

var (
	config *Config = nil
)

func GetEnv() *Config {
	if config != nil {
		return config
	}
	cfg := new(Config)
	err := env.Fill(cfg)
	if err != nil {
		panic(err)
	}
	config = cfg
	return config
}

func (this *Config) GetAddr() string {
	return GetEnv().Host + ":" + strconv.Itoa(GetEnv().Port)
}

func (this *Mongodb) GetUrl() string {
	if this.URL != "" {
		return this.URL
	} else {
		return "mongodb://" + this.Host + ":" + strconv.Itoa(this.Port)
	}
}
