package main

import (
	_ "cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"net/http"
)

func FetchAPI(url string) (FilteredResult, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var order FilteredResult
	err = decoder.Decode(&order)

	fmt.Println("Query made!!!", url)
	return order, nil
}

func CheckOrders() error {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var orders []interface{}
	var do = true
	var nextPageUrl = ""
	var rawData FilteredResult
	var initialCall = true

	rawData, _ = FetchAPI(squareSpaceURL)

	for do {
		initialResults := rawData.Result

		do = rawData.Pagination.HasNextPage
		nextPageUrl = rawData.Pagination.NextPageUrl

		for i := 0; i < len(initialResults); i++ {
			/**
			Fill the orders Interface.
			*/
			orders = append(orders, initialResults[i])

			/**
			Insert the data on the users collection and create the user doc with the email as id.
			*/
			_, _ = client.Collection("users").Doc(
				initialResults[i].CustomerEmail).Create(ctx, initialResults[i].UserData)

			fmt.Println("Order ID -> " + initialResults[i].ID)

			for j := 0; j < len(initialResults[i].Items); j++ {
				/**
				Fill the user email doc with the orders.
				*/
				initialResults[i].Items[j].CreatedOn = initialResults[i].CreatedOn
				_, _ = client.Collection("users").Doc(
					initialResults[i].CustomerEmail).Collection("orders").Doc(
					initialResults[i].Items[j].ProductName).Set(ctx, initialResults[i].Items[j])
				fmt.Println("Product ID -> " + initialResults[i].Items[j].ProductID)
			}
		}

		if do && nextPageUrl != "" && !initialCall {
			orders = []interface{}{}
			rawData, _ = FetchAPI(nextPageUrl)
			initialResults = rawData.Result
		}
		initialCall = false
	}
	return nil
}

func GetOrdersByEmail(email string) ([]SimplifiedLineItem, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var userOrders []SimplifiedLineItem
	data := client.Collection("users").Doc(email).Collection("orders").Documents(ctx)
	for {
		doc, err := data.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var order SimplifiedLineItem
		_ = doc.DataTo(&order)
		userOrders = append(userOrders, order)
		fmt.Println(doc.Data())
	}
	return userOrders, nil
}
