package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devfullcycle/goexpert/9-API/internal/dto"
	"github.com/devfullcycle/goexpert/9-API/internal/entity"
	"github.com/devfullcycle/goexpert/9-API/internal/infra/database"
	entityPKG "github.com/devfullcycle/goexpert/9-API/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary      Create Product
// @Description  API para criação de produto
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request    body     dto.CreateProductInput  	true	"product request"
// @Success      201
// @Failure      500 	{object}  Error
// @Router       /products [post]
// @Security	ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary      Get Product
// @Description  API para Ler um produto
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id		path   		string  	true	"product ID"	Format(uuid)
// @Success      200	{object}	entity.Product
// @Failure      404
// @Failure      500 	{object}  	Error
// @Router       /products/{id} [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary      Update Product
// @Description  API para atualização de produto
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id			path		string						true	"product ID" Format(uuid)
// @Param        request	body		dto.CreateProductInput		true	"product request"
// @Success      200
// @Failure      404
// @Failure      500 	{object}  Error
// @Router       /products/{id} [put]
// @Security	ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary      Delete Product
// @Description  API para Eliminar o produto
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id			path		string						true	"product ID" Format(uuid)
// @Success      200
// @Failure      404
// @Failure      500 	{object}  Error
// @Router       /products/{id} [delete]
// @Security	ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// List Product godoc
// @Summary      List Products
// @Description  API para Listar os produtos
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page   query		string  	false	"page number"
// @Param        limit	query		string  	false	"limit"
// @Success      200	{array}		entity.Product
// @Failure      404
// @Failure      500 	{object}  	Error
// @Router       /products [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
