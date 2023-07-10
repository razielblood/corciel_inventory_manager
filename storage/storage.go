package storage

import (
	"github.com/razielblood/corciel_inventory_manager/types"
)

type Storage interface {
	CreateCategory(*types.Category) error
	CreateManufacturer(*types.Manufacturer) error
	CreateProduct(*types.Product) error

	UpdateCategory(*types.Product) error
	UpdateManufacturer(*types.Product) error
	UpdateProduct(*types.Product) error

	GetCategoryByID(int) (*types.Category, error)
	GetManufacturerByID(int) (*types.Manufacturer, error)
	GetProductByID(int) (*types.Product, error)

	GetCategories() ([]*types.Category, error)
	GetManufacturers() ([]*types.Manufacturer, error)
	GetProducts() ([]*types.Product, error)
}
