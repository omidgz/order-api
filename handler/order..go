package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (this *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("C")
}

func (this *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("R/A")
}

func (this *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("R")
}

func (this *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("U")
}

func (this *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("D")
}
