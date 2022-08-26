package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MongoDBConfig MongoDBConfig
}

type MongoDBConfig struct {
	Region   string `envconfig:"DOCUMENTDB_REGION" required:"false" default:"ap-northeast-1"`
	Endpoint string `envconfig:"DOCUMENTDB_ENDPOINT" required:"false" default:"0.0.0.0:27017"`
	User     string `envconfig:"DOCUMENTDB_USER" required:"false" default:"admin"`
	Password string `envconfig:"DOCUMENTDB_PASSWORD" required:"false" default:"admin"`
}

var c *Config

func Init() *Config {
	var cnf Config
	if c != nil {
		return c
	}
	err := envconfig.Process("", &cnf)
	if err != nil {
		panic(err)
	}
	c = &cnf
	return c
}

func GetConfig() *Config {
	if c == nil {
		c = Init()
	}
	return c
}
