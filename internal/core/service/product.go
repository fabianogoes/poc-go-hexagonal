package service

import (
	"context"

	"github.com/demo/go-hexagonal/internal/core/domain"
	"github.com/demo/go-hexagonal/internal/core/port"
)

type ProductService struct {
	repository port.ProductRepository
}

func NewProductService(rep port.ProductRepository) *ProductService {
	return &ProductService{repository: rep}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	product, err := s.repository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, id uint64) (*domain.Product, error) {
	product, err := s.repository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) ListProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	product, err := s.repository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint64) error {
	err := s.repository.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
