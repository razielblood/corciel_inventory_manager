package storage

import (
	"database/sql"

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

	LoginUser(*types.LoginRequest) (*types.User, error)
	CreateUser(*types.CreateUserRequest) (*types.User, error)
}

func parseProduct(rows *sql.Rows, product *types.Product) error {
	var manufacturerID, categoryID int
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.WeightInKG,
		&product.PiecesPerPackage,
		&product.Image,
		&manufacturerID,
		&categoryID,
	)
	if err != nil {
		return err
	}

	product.Manufacturer = &types.Manufacturer{ID: manufacturerID}
	product.Category = &types.Category{ID: categoryID}

	return nil
}

func GetCategoriesAsMap(s Storage) (map[int]*types.Category, error) {
	categories, err := s.GetCategories()
	if err != nil {
		return nil, err
	}
	categoriesMap := make(map[int]*types.Category)
	for _, category := range categories {
		categoriesMap[category.ID] = category
	}
	return categoriesMap, nil
}

func GetManufacturersAsMap(s Storage) (map[int]*types.Manufacturer, error) {
	manufacturers, err := s.GetManufacturers()
	if err != nil {
		return nil, err
	}
	manufacturersMap := make(map[int]*types.Manufacturer)
	for _, manufacturer := range manufacturers {
		manufacturersMap[manufacturer.ID] = manufacturer
	}
	return manufacturersMap, nil
}
