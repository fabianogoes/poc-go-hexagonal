package handler

import (
	"strconv"

	"github.com/demo/go-hexagonal/internal/core/domain"
	"github.com/demo/go-hexagonal/internal/core/port"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(svc port.ProductService) *ProductHandler {
	return &ProductHandler{svc}
}

type createProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	product := domain.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	response, err := h.service.CreateProduct(ctx, &product)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	handleSuccess(ctx, response)
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	product, err := h.service.GetProduct(ctx, id)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	response := product

	handleSuccess(ctx, response)
}

func (h *ProductHandler) ListProducts(ctx *gin.Context) {
	products, err := h.service.ListProducts(ctx)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	response := products

	handleSuccess(ctx, response)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	product := domain.Product{
		ID:    id,
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	response, err := h.service.UpdateProduct(ctx, &product)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	handleSuccess(ctx, response)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	err = h.service.DeleteProduct(ctx, id)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
