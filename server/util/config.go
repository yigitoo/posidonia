package util

type Configure interface {
	NewConfig() Config
	GetPort() int16
	SetPort(port int16)
	LogginPort()
}

type Config struct {
	Port string
}

func (c Config) LoggingPort() {
	fmt.Println("Current port is: ", c.Port)
}

func (c *Config) SetPort(port int16) {
	fmt.Println("Current port is:")
}

func (c Config) GetPort() int16 {
	return c.Port
}
