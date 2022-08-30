package main

import (
	"flag"
	"log"
	"os"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	router := app.routes()

	err := router.Run("localhost:8080")
	if err != nil {
		log.Println(err)
	}
}
