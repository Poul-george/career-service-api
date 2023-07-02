package config

type MySQL struct {
	User string `mapstructure:"user" validate:"required"`
	Host string `mapstructure:"host" validate:"required"`
	// HostReader string `mapstructure:"host_reader" validate:"required"`
	Port int `mapstructure:"port" validate:"required"`
	// PortReader int    `mapstructure:"port_reader" validate:"required"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database" validate:"required"`
	Encoding string `mapstructure:"encoding" validate:"required"`
	// Debug      bool   `mapstructure:"debug"`
	// LoggerColorful   bool `mapstructure:"logger_colorful"`
	// ConnMaxLifetime   uint64 `mapstructure:"conn_max_lifetime"`
	// ConnMaxIdletime   uint64 `mapstructure:"conn_max_idletime"`
	// MaxIdleConns   int `mapstructure:"max_idle_conns"`
	// MaxOpenConns   int `mapstructure:"max_open_conns"`
	// SqlMode string `mapstructure:"sql_mode"`
}

func GetMySQLConfig() MySQL {
	c := getConfig().MySQL

	return MySQL{
		User:     c.User,
		Host:     c.Host,
		Port:     c.Port,
		Password: c.Password,
		Database: c.Database,
		Encoding: c.Encoding,
	}
}
