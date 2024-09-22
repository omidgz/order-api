package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func (this *Handler) ID(r *http.Request) uint {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}

func (this *Handler) Decode(r *http.Request, target any) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return nil
}
