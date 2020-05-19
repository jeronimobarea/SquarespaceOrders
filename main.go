package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func createClient(ctx context.Context) *firestore.Client {
	/**
	Google Firestore client configuration.
	*/
	projectID := projectID

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func main() {
	/**
	Ticker configuration for checking the api every 24h.
	*/
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for t := range ticker.C {
			_ = t
			fmt.Println("Checking Squarespace API on -> ", time.Now())
			_ = InsertManyOrdersFiltered()
		}
	}()

	/**
	Gorilla mux router configuration.
	*/
	r := mux.NewRouter()

	r.HandleFunc("/orders", FetchOrdersFilter).Methods("POST")
	r.HandleFunc("/orders", GetFilteredOrders).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server listening on http://localhost:8000")
	_ = http.ListenAndServe(":8000", r)
}