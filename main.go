package main

import (
	"flag"
	"log"
	"os"

	"github.com/razielblood/corciel_inventory_manager/api"
	"github.com/razielblood/corciel_inventory_manager/storage"
)

func main() {
	listenPort := flag.String("port", "3550", "Listen address for the API")
	listenAddr := flag.String("host", "localhost", "Host to listen to")
	flag.Parse()

	dbUsername := os.Getenv("CORCIEL_INVENTORY_DB_USERNAME")
	dbPass := os.Getenv("CORCIEL_INVENTORY_DB_PASSWORD")
	dbHost := os.Getenv("CORCIEL_INVENTORY_DB_HOST")
	dbPort := os.Getenv("CORCIEL_INVENTORY_DB_PORT")
	dbName := os.Getenv("CORCIEL_INVENTORY_DB_NAME")

	store, err := storage.NewMariaDBStore(dbUsername, dbPass, dbHost, dbPort, dbName)

	if err != nil {
		log.Fatalf("Could not create connection to database: %v", err.Error())
		return
	}

	server := api.NewAPIServer(*listenAddr, *listenPort, store)

	server.Run()

}
