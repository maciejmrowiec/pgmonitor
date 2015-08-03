package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// build number set on during linking
var minversion string

type AppConfig struct {
	verbose     bool
	newRelicKey string
	database    string
	interval    int
	user        string
	host        string
	version     bool
}

func HandleUserOptions() AppConfig {

	var config AppConfig

	flag.BoolVar(&config.verbose, "verbose", false, "Verbose mode")
	flag.StringVar(&config.newRelicKey, "key", "", "Newrelic license key (required)")
	flag.StringVar(&config.database, "database", "", "Database name (required)")
	flag.IntVar(&config.interval, "interval", 1, "Sampling interval [min]")
	flag.StringVar(&config.user, "user", "postgres", "Database user name")
	flag.StringVar(&config.host, "host", "localhost:5432", "Database host")
	flag.BoolVar(&config.version, "version", false, "Print version")

	flag.Parse()

	if config.version {
		fmt.Printf("Build: %s\n", minversion)
		os.Exit(0)
	}

	if config.newRelicKey == "" ||
		config.database == "" {
		flag.PrintDefaults()
		log.Fatal("Required parameter missing.")
	}

	return config
}
