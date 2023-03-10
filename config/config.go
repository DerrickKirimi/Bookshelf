package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Debug  bool `env:"DB_HOST, required"`
	Db     DBConf
	Server serverConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE, required"`
	Debug        bool          `env:"SERVER_DEBUG, required"`
}

type DBConf struct {
	Host     	string `env:"DB_HOST, required"`
	Port     	int    `env:"DB_PORT, required"`
	Password    string `env:"DB_PASS, required"`
	Username 	string `env:"DB_USER, required"`
	DbName     	string `env:"DB_NAME, required"`
}

func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}

/*func NewDb() *ConfDB {
	var c ConfDB
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
*/
