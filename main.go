package main

import (
	"log"
	"net/http"

	"github.com/SkaisgirisMarius/order-processor.git/db"
	"github.com/SkaisgirisMarius/order-processor.git/server"
)

func main() {
	log.Println("Starting Order Processor")
	database, err := db.ConnectToMySQL()
	if err != nil {
		log.Fatal("Could not start MySQL server. ", err)
	}
	r := server.NewRouter(database)
	defer database.Close()

	http.ListenAndServe(":8080", r)

}
