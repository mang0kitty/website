package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mang0kitty/website/src/handlers"
)

func main() {

	listenAddr := "8080"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.Handle("/", handlers.Handle())
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))

}
