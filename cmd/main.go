package main

import (
	"github.com/spf13/viper"
	"log"
	"vtb_api/internal/handler"
	"vtb_api/internal/repository"
	"vtb_api/internal/server"
	"vtb_api/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repo := repository.NewRepo()
	usecase := service.NewUseCase(repo)
	contr := handler.NewController(usecase)
	srv := server.NewServer(contr)
	if err := srv.Run("localhost:" + viper.GetString("port")); err != nil {
		log.Fatalf("Server eng with error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
