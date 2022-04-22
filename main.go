package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zenreach/go-gadget/pkg/handlers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := 3000

	fh := http.FileServer(http.Dir("./images/"))
	r := mux.NewRouter()

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", fh))
	http.Handle("/", r)
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/auth/{provider}", handlers.AuthHandler).Methods("GET")
	r.HandleFunc("/auth/{provider}/callback", handlers.AuthCallBackHandler).Methods("GET")
	// r.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))

	srv := &http.Server{
		Addr:         "127.0.0.1:" + strconv.Itoa(port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("starting server on port %v", port)
	log.Fatal(srv.ListenAndServe())

}
