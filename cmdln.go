package main

import (
	"flag"
	"log"
)

type AppConfig struct {
	verbose     bool
	newRelicKey string
	database    string
	interval    int
	user        string
}

func HandleUserOptions() AppConfig {

	var config AppConfig

	flag.BoolVar(&config.verbose, "verbose", false, "Verbose mode")
	flag.StringVar(&config.newRelicKey, "key", "", "Newrelic license key (required)")
	flag.StringVar(&config.database, "database", "", "Database name (required)")
	flag.IntVar(&config.interval, "interval", 1, "Sampling interval [min]")
	flag.StringVar(&config.user, "user", "postgres", "Database user name")

	flag.Parse()

	if config.newRelicKey == "" ||
		config.database == "" {
		flag.PrintDefaults()
		log.Fatal("Required parameter missing.")
	}

	return config
}
