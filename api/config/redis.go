package config

type Redis struct {
	Host      string `mapstructure:"host" validate:"required"`
	Port      string `mapstructure:"port" validate:"required"`
	SessionDB int    `mapstructure:"session_db" validate:"required"`
}

func GetRedisConfig() Redis {
	c := getConfig().Redis

	return Redis{
		Host:      c.Host,
		Port:      c.Port,
		SessionDB: c.SessionDB,
	}
}
