package handlers

import (
	"fmt"
	"net/http"
)

type Order struct {
}

// This will create all the orders
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a order")
}

// This will get all the orders
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a order")
}

// Getting results by ID
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a order")
}

// Updating by ID
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a order")
}

// Deleting by ID
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a order")
}
