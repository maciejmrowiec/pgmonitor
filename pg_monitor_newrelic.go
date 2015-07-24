package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	nr "github.com/yvasiyarov/newrelic_platform_go"
	"log"
	"os"
)

func main() {

	HandleUserOptions()

	db, err := OpenDatabaseConnection(config.database, config.user)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	table_stats := InitTableStatsComponent(hostname, config.verbose)

	plugin := nr.NewNewrelicPlugin("0.0.1", config.new_relic_key, config.interval*60)
	plugin.AddComponent(table_stats)
	plugin.Run()

}

func OpenDatabaseConnection(database string, user string) (*sql.DB, error) {
	connection_uri := "postgres://" + user + "@localhost:5432/" + database + "?sslmode=disable"
	db, err := sql.Open("postgres", connection_uri)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
