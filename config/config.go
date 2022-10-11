package config

import (
	"errors"
	"github.com/spf13/viper"
)

type DatabaseConfigurations struct {
	ConnectionString string
}
type KafkaConfiguration struct {
	URL   string
	Topic string
}
type Configurations struct {
	Database DatabaseConfigurations
	Kafka    KafkaConfiguration
}

func SetUpViper(Configurations Configurations) (Configurations, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./resources")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return Configurations, errors.New("failed to read config file")
	}
	err := viper.Unmarshal(&Configurations)
	if err != nil {
		return Configurations, errors.New("failed to Convert to object")
	}
	return Configurations, nil
}
