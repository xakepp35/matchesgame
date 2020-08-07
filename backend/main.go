/*
	Игра "спички".
	Matches game fullstack example.
	Rules:
		There is 30 safety matches on the deck.
		2 players can take 1,2 or 3 math in a turn.
		To win, a player must force opponent to take last match

	Author: S.Leonovich
	Date: 07.08.2020
*/

package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

)

func initLogger() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func initViper(configPath string) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("matches")

	viper.SetDefault("loglevel", "debug")
	viper.SetDefault("listen", ":3000")
	viper.SetDefault("db.url", "postgres://matches@db/matches?sslmode=disable&pool_max_conns=10")

	if configPath != "" {
		log.Infof("Parsing config: %s", configPath)
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Unable to read config file: %s", err)
		}
	} else {
		log.Infof("Config file is not specified.")
	}
}

func configureLogger() {
	logLevelString := viper.GetString("loglevel")
	logLevel, err := log.ParseLevel(logLevelString)
	if err != nil {
		log.Fatalf("Unable to parse loglevel: %s", logLevelString)
	}
	log.SetLevel(logLevel)
}

// каркас взят с https://github.com/afiskon/go-rest-service-example/blob/master/cmd/rest-service-example/main.go
// кобра разумеется удалена как совершенно ненужное тут дело
func main() {
	initLogger()
	initViper(os.Getenv("CONFIG_PATH"))
	configureLogger()

	dbURL := viper.GetString("db.url")
	log.Infof("Using DB URL: %s", dbURL)

	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	defer pool.Close()
	log.Infof("pgx pool connected!")

	listenAddr := viper.GetString("listen")
	log.Infof("Starting HTTP server at %s...", listenAddr)
	http.Handle("/", initHandlers(pool))
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}

	log.Info("HTTP server terminated")
}

