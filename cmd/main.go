package main

import (
	"Todo_list/internal/repository"
	ps "Todo_list/internal/repository/postgres"
	"Todo_list/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	srv := service.NewService(service.NewTaskService(rep))

}

func initConfig() error {
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}
