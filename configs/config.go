package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	Server    Server
	Postgres  Postgres
	JWTSecret Auth
}

type Server struct {
	Host string
	Port string
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type Auth struct {
	Secret      string
	TokenExpiry time.Duration
}

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	tokenExpiry, err := time.ParseDuration(viper.GetString("TOKEN_EXPIRY"))
	if err != nil {
		log.Fatalf("error when parsing TOKEN_EXPIRY: %s", err)
	}

	Config = config{
		Server: Server{
			Host: viper.GetString("SERVER_HOST"),
			Port: viper.GetString("SERVER_PORT"),
		},
		Postgres: Postgres{
			Username: viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
			DBName:   viper.GetString("POSTGRES_DB"),
		},
		JWTSecret: Auth{
			Secret:      viper.GetString("SECRET_KEY"),
			TokenExpiry: tokenExpiry,
		},
	}
}
