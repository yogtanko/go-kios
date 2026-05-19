package config

import "os"

type Config struct {
	Addr      string
	Db        dbConfig
	JwtSecret []byte
}

type dbConfig struct {
	DBUrl string
}

func GetConfig() *Config {
	return &Config{
		Addr: ":8080",
		Db: dbConfig{
			DBUrl: os.Getenv("DBUrl"),
		},
		JwtSecret: []byte("ini,rahasia"),
	}
}
