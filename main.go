package main

import (
	"flag"
	"fmt"
	"net/http"
	"tc/controllers"
	"tc/models"

	"github.com/gorilla/mux"
)

func main() {
	boolPtr := flag.Bool(
		"prod",
		false,
		"Provide this flag in production. This ensures that a .config file is provided before the application starts.",
	)
	flag.Parse()

	cfg := LoadConfig(*boolPtr)
	services, err := models.NewServices(
		models.WithTezTools(),
	)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	homeC := controllers.NewHomepage(services.TezTools)

	r.HandleFunc("/", homeC.Index).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	// Assets
	assetHandler := http.FileServer(http.Dir("./dist/"))
	assetHandler = http.StripPrefix("/dist/", assetHandler)
	r.PathPrefix("/dist/").Handler(assetHandler)

	fmt.Printf("Starting the server on: %d...\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
}
