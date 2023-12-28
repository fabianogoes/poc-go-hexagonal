package repository

import (
	"context"

	"github.com/demo/go-hexagonal/internal/core/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	// TODO: insert product into database
	return product, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uint64) (*domain.Product, error) {
	// TODO: select product from database
	return &domain.Product{
		ID:    id,
		Name:  "Product 1",
		Price: 1000,
		Stock: 10,
	}, nil
}

func (r *ProductRepository) ListProducts(ctx context.Context) ([]domain.Product, error) {
	// TODO: select products from database
	return []domain.Product{
		{
			ID:    1,
			Name:  "Product 1",
			Price: 1000,
			Stock: 10,
		},
		{
			ID:    2,
			Name:  "Product 2",
			Price: 2000,
			Stock: 20,
		},
	}, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	// TODO: update product in database
	return product, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uint64) error {
	// TODO: delete product from database
	return nil
}
