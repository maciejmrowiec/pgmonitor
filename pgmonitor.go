package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	platform "github.com/yvasiyarov/newrelic_platform_go"
	"log"
	"os"
)

func main() {
	config := HandleUserOptions()

	db, err := OpenDatabaseConnection(config.database, config.user)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	plugin := platform.NewNewrelicPlugin("0.0.1", config.newRelicKey, config.interval*60)
	plugin.AddComponent(InitTupleComponent(db, hostName, config.verbose))
	plugin.AddComponent(InitTupleSummaryComponent(db, hostName, config.verbose))
	plugin.AddComponent(InitTableSizeComponent(db, hostName, config.verbose))

	plugin.Verbose = config.verbose
	plugin.Run()
}

func OpenDatabaseConnection(database string, user string) (*sql.DB, error) {
	connectionUri := "postgres://" + user + "@localhost:5432/" + database + "?sslmode=disable"
	db, err := sql.Open("postgres", connectionUri)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
