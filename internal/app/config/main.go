package config

import (
	"flag"
	"os"
)

type Config struct {
	serverAddress string
	dbAddress     string
}

func (c *Config) GetServerAddress() string {
	return c.serverAddress
}

func (c *Config) GetDBAddress() string {
	return c.dbAddress
}

func (c *Config) SetServerAddress(value string) {
	c.serverAddress = value
}

func (c *Config) SetDBAddress(value string) {
	c.dbAddress = value
}

func (c *Config) ParseFlags() {
	flag.StringVar(&c.serverAddress, "s", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.dbAddress, "d", "", "database address")

	flag.Parse()

	if envStartAddress, in := os.LookupEnv("SERVER_ADDRESS"); in {
		c.SetServerAddress(envStartAddress)
	}

	if envDBAddress, in := os.LookupEnv("DB_ADDRESS"); in {
		c.SetDBAddress(envDBAddress)
	}
}
