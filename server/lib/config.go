package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration interface {
	// constructor
	NewConfig() Config

	// @getter: Config.Port
	GetPort() int16
	// @setter: Config.Port
	SetPort(port int16)
	// @logger: Config.Port
	LogginPort()

	// @setter: ApiKeys (from .env file in the root of the project: server folder directory).
	SetApiKeys()
	// @getter: ApiKeys (from .env file in the root of the project: server folder directory).
	GetApiKeys(key string) string

	/* @DEPRECATED!!!
	// @setter: env_variables of project.
	SetEnv(key string, value string)
	// @getter: env_variables of project.
	GetEnv() map[string]string
	*/
}

type Config struct {
	// Port number of server.
	Port int16
	// All APIs authorization keys for 3rd party microservices.
	ApiKeys map[string]string
}

func NewConfig() Config {
	/*
		@description:
			- Constructor of config struct.
			- It will be accessable from Configure interface.
		i dunno, but i use that port
		for a port num problem.
	*/
	return Config{Port: 5632, ApiKeys: nil}
}

func (c Config) LoggingPort() {
	/*
		@description:
			- It'll be print the current port number of golang server.
	*/
	fmt.Println("Current port is: ", c.Port)
}

func (c *Config) SetPort(port int16) {
	/*
		@description:
			- Setter of Port number in Config struct.

		@params:
			- port: int16 = wanted port number of project server to serve.
	*/
	c.Port = port
}

func (c *Config) GetPort() int16 {
	/*
		@description:
			- Getter of Port number in Config struct.
	*/
	return c.Port
}

func (c *Config) SetApiKeys() {
	/*
		@description:
			- Getter of API Keys map from .env file in the root of golang project.
	*/
	godotenv.Load()
	api_keys := make(map[string]string, 3)

	// setting
	api_keys["geocode"] = os.Getenv("API_KEY_GEOCODE")

	// put in config
	c.ApiKeys = api_keys
}

func (c Config) GetApiKeys(key string) string {
	/*
		@description:
			- Getter for Api Key (map[string]string).
			-- The keys are stored in .env files
	*/
	return c.ApiKeys[key]
}
