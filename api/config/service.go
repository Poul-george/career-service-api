package config

type Service struct {
	SecretKey string `mapstructure:"secret_key"`
}

func GetServiceConfig() Service {
	c := getConfig().Service

	return Service{
		SecretKey: c.SecretKey,
	}
}
