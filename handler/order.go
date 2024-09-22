package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/omidgz/order-api/model"
	"gorm.io/gorm"
)

type Order struct {
	Handler
}

func NewOrderHandler(db *gorm.DB) *Order {
	return &Order{
		Handler: Handler{db: db},
	}
}

func (this *Order) DB() *gorm.DB {
	return this.db.Model(&model.Order{})
}

func (this *Order) Create(w http.ResponseWriter, r *http.Request) {
	order := model.Order{}
	this.Decode(r, &order)

	this.DB().Create(&order)
	w.WriteHeader(http.StatusCreated)
}

func (this *Order) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}

	// later can add the page size as a param
	pageSize := 50
	offset := (page - 1) * pageSize

	var orders []model.Order
	this.DB().Limit(pageSize).Offset(offset).Find(&orders)
	json.NewEncoder(w).Encode(&orders)
	w.WriteHeader(http.StatusOK)
}

func (this *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	order := model.Order{}
	result := this.DB().First(&order, this.ID(r))
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&order)
	w.WriteHeader(http.StatusOK)
}

func (this *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	order := model.Order{}
	this.Decode(r, &order)
	json.NewEncoder(w).Encode(&order)

	result := this.DB().Where("id = ?", order.ID).Updates(&order)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(&order)
	w.WriteHeader(http.StatusOK)
}

func (this *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	result := this.DB().Delete(&model.Order{}, this.ID(r))
	if result.RowsAffected == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
