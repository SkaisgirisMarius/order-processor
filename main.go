package main

import (
	"github.com/SkaisgirisMarius/order-processor.git/db"
	"github.com/SkaisgirisMarius/order-processor.git/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting Order Processor")
	database, err := db.ConnectToMySQL()
	if err != nil {
		log.Fatal("Could not start MySQL server. ", err)
	}
	r := server.NewRouter(database)

	server.StartServer(r)
}
