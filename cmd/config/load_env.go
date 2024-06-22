package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBTestName string `mapstructure:"POSTGRES_DB_TEST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`

	GRPCPort     string `mapstructure:"GRPC_PORT"`
	AUTH_SERVICE string `mapstructure:"AUTH_SERVICE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("vocab")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
