package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/viper"
	"vtb_api/internal/handler"
	"vtb_api/internal/repository"
	"vtb_api/internal/server"
	"vtb_api/internal/service"
)

// @title Rent Transport
// @version 1.0
// @description WEB API

// @host localhost:8000
// @basePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repo := repository.NewRepo(db)
	usecase := service.NewUseCase(repo)
	contr := handler.NewController(usecase)
	srv := server.NewServer(contr)
	if err := srv.Run("localhost:" + viper.GetString("port")); err != nil {
		logrus.Fatalf("Server eng with error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
