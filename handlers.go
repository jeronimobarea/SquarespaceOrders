package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchOrdersAndFilter(w http.ResponseWriter, r *http.Request) {
	/**
	Handler for manually check for updates.
	*/
	err := CheckOrders()

	if err != nil {
		fmt.Println(err)
	}
	res := `{ "inserted": true, "status": "ok" }`
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(res))
}

func GetFilteredOrders(w http.ResponseWriter, r *http.Request) {
	/**
	Handler for getting user orders by an email.
	*/
	email := r.URL.Query().Get("email")
	data, _ := GetOrdersByEmail(email)

	serialized, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}
