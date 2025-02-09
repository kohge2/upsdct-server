package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	DbDsn      string `mapstructure:"DB_DSN"`
	EncryptKey string `mapstructure:"ENCRYPT_KEY"`
	//SessionSecret string `mapstructure:"SESSION_SECRET"` セッションID方式でログインする場合はここでセッションIDの秘密鍵を管理
}

var Env *EnvConfig

func InitEnvConfig() (*EnvConfig, error) {
	if os.Getenv("ENV") == "test" {
		Env = &EnvConfig{}
		return Env, nil
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	var config *EnvConfig

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to unmarshal config", err)
	}

	Env = config
	return Env, nil
}

func init() {
	InitEnvConfig()
}
