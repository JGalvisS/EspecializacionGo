package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) error
	Delete(id int)error
}

type repository struct {
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// Update actualiza un producto existente por id
func (r *repository) Update(id int, p domain.Product)  error {
	if !r.validateCodeValue(p.CodeValue) {
		return errors.New("code value already exists")
	}

	update := false
	for key, producto := range r.list{
		if producto.Id == id {
			p.Id = id
			r.list[key] =p
			update = true
		}
	}

	if update == false {
		return errors.New("product not found")
	}
	return nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

//Delete elimna un producto por id
func (r *repository)Delete(id int)error {
	_,valida:= r.GetByID(id)
	if valida != nil {
		return valida
	}
	
	for key, produto := range r.list{
		if produto.Id == id {
			r.list = append(r.list[:key], r.list[key+1:]...)
		}
	}
	
	_, err:=r.GetByID(id)
	if err != nil {
		return nil
	}
	return err

}