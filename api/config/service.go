package config

type Service struct {
	SecretKey string `mapstructure:"secret_key" validate:"required"`
}

func GetServiceConfig() Service {
	c := getConfig().Service

	return Service{
		SecretKey: c.SecretKey,
	}
}
