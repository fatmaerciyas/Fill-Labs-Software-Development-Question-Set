package repos

import (
	"github.com/fatmaerciyas/Fill-Labs-Software-Development-Question-Set/tree/main/Question4/entities"
	"fmt"
)

type ProductRepo struct {
	products []entities.Product
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{make([]entities.Product, 0)}
}

func (p *ProductRepo) Create(partial entities.Product) entities.Product {
	newItem := partial
	newItem.ID = uint(len(p.products)) + 1
	p.products = append(p.products, newItem)
	return newItem
}

func (p *ProductRepo) GetList() []entities.Product {
	return p.products
}

func (p *ProductRepo) GetOne(id uint) (entities.Product, error) {
	for _, it := range p.products {
		if it.ID == id {
			return it, nil
		}
	}
	return entities.Product{}, fmt.Errorf("key '%d' not found", id)
}

func (p *ProductRepo) Update(id uint, amended entities.Product) (entities.Product, error) {
	for i, it := range p.products {
		if it.ID == id {
			amended.ID = id
			p.products = append(p.products[:i], p.products[i+1:]...)
			p.products = append(p.products, amended)
			return amended, nil
		}
	}
	return entities.Product{}, fmt.Errorf("key '%d' not found", amended.ID)
}

func (p *ProductRepo) DeleteOne(id uint) (bool, error) {
	for i, it := range p.products {
		if it.ID == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}


// C:\Users\fatma\source\repos\Fill-Labs-Software-Development-Question-Set\Question4
// C:\Users\fatma\source\repos\Fill-Labs-Software-Development-Question-Set\Question4\controllers\product-controller.go