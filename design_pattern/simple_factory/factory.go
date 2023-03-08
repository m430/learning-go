package simple_factory

import "fmt"

type Service interface {
	Create(name string) bool
}

func NewService(serviceType string) Service {
	if serviceType == "product" {
		return &productService{}
	}
	if serviceType == "order" {
		return &orderService{}
	}
	return nil
}

type productService struct{}

func (p *productService) Create(name string) bool {
	fmt.Printf("create product name %s success", name)
	return true
}

type orderService struct{}

func (o *orderService) Create(name string) bool {
	fmt.Printf("create order name %s success", name)
	return true
}
