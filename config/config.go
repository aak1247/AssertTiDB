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
		Port     int    `default:"4000" env:"ASSERT_TIDB_PORT"`
		Username string `default:"" env:"ASSERT_TIDB_USERNAME"`
		Password string `default:"" env:"ASSERT_TIDB_PASSWORD"`
		Database string `require:"true" default:"" env:"ASSERT_TIDB_DATABASE"`
		Params   string `default:"" env:"ASSERT_TIDB_DB_PARAMS"`
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

func (this *Tidb) GetUrl() string {
	if this.URL != "" {
		return this.URL
	} else {
		var url string
		url += this.Username
		if this.Password != "" {
			url += ":" + this.Password
		}
		if this.Username != "" {
			url += "@"
		}
		url += "tcp(" + this.Host + ":" + strconv.Itoa(this.Port) + ")/" + this.Database
		if this.Params != "" {
			url += "?" + this.Params
		}
		return url
		// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	}
}
