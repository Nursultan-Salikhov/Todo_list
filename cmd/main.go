package main

import (
	"Todo_list/internal/repository"
	ps "Todo_list/internal/repository/postgres"
	"Todo_list/internal/service"
	"Todo_list/internal/transport/handler"
	"Todo_list/internal/transport/server"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("initConfig failed: %s", err)
		return
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := ps.NewPostgresDB(ps.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("init db failed: %s", err)
		return
	}

	rep := repository.NewRepository(ps.NewTask(db))
	services := service.NewService(service.NewTaskService(rep))
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("running http server failed: %s", err)
		}
	}()

	logrus.Println("APP started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting donw", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on DB connection close", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}
