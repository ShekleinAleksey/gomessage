package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ShekleinAleksey/gomessage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.Println("This is a debug message")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// })
	// if err != nil {
	// 	logrus.Fatalf("failed to initialize db: %s", err.Error())
	// }

	// repos := repository.NewRepository(db)
	// services := service.NewService(repos)
	// handlers := handler.NewHandler(services)
	// http.HandleFunc("/message", handleMessage)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)

	srv := new(gomessage.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
