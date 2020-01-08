package  main

import (
	"log"
	_"net/http"
	"os"
	_"github.com/gorilla/handlers"
	"rest-and-go/store"
	"net/http"
)

func main()  {
	// get the port from env variable
	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("Port must be set")
	}

	router := store.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))

}