package config

type Server struct {
	StartAddress    string   `mapstructure:"start_address" validate:"required"`
	IdleTimeout     int64    `mapstructure:"idle_timeout" validate:"required"`
	EchoAllowOrigin []string `mapstructure:"echo_allow_origin" validate:"required"`
}

func GetServerConfig() Server {
	c := getConfig().Server

	return Server{
		StartAddress:    c.StartAddress,
		IdleTimeout:     c.IdleTimeout,
		EchoAllowOrigin: c.EchoAllowOrigin,
	}
}
