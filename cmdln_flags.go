package main

import (
	"flag"
	"log"
)

var config struct {
	verbose       bool
	new_relic_key string
	database      string
	interval      int
	user          string
}

func HandleUserOptions() {
	flag.BoolVar(&config.verbose, "verbose", false, "Verbose mode")
	flag.StringVar(&config.new_relic_key, "key", "", "Newrelic license key (required)")
	flag.StringVar(&config.database, "database", "", "Database name (required)")
	flag.IntVar(&config.interval, "interval", 1, "Sampling interval [min]")
	flag.StringVar(&config.user, "user", "postgres", "Database user name")

	flag.Parse()

	if config.new_relic_key == "" ||
		config.database == "" {
		flag.PrintDefaults()
		log.Fatal("Required parameter missing.")
	}
}
