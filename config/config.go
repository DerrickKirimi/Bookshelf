package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server serverConf
}

type serverConf struct {
	Port        	int				`env:"SERVER_PORT,required"`
	TimeoutRead	 	time.Duration 	`env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite 	time.Duration 	`env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle		time.Duration 	`env:"SERVER_TIMEOUT_IDLE, required"`
	Debug			bool			`env:"SERVER_DEBUG, required"`		
}


func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c);err != nil{
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