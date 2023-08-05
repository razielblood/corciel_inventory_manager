package storage

import (
	"database/sql"

	"github.com/razielblood/corciel_inventory_manager/types"
)

type Storage interface {
	CreateCategory(*types.Category) error
	CreateManufacturer(*types.Manufacturer) error
	CreateProduct(*types.Product) error
	CreateBrand(*types.Brand) error

	UpdateCategory(*types.Category) error
	UpdateManufacturer(*types.Manufacturer) error
	UpdateProduct(*types.Product) error
	UpdateBrand(*types.Brand) error

	GetCategoryByID(int) (*types.Category, error)
	GetManufacturerByID(int) (*types.Manufacturer, error)
	GetProductByID(int) (*types.Product, error)
	GetBrandByID(int) (*types.Brand, error)

	GetCategories() ([]*types.Category, error)
	GetManufacturers() ([]*types.Manufacturer, error)
	GetProducts() ([]*types.Product, error)
	GetBrands() ([]*types.Brand, error)

	GetUserByID(string) (*types.User, error)
	LoginUser(*types.LoginRequest) (*types.User, error)
	CreateUser(*types.CreateUserRequest) (*types.User, error)
}

func parseBrand(rows *sql.Rows, brand *types.Brand) error {
	var manufacturerID int
	err := rows.Scan(
		&brand.ID,
		&brand.Name,
		&manufacturerID,
	)
	if err != nil {
		return err
	}

	brand.Manufacturer = &types.Manufacturer{ID: manufacturerID}

	return nil
}

func parseProduct(rows *sql.Rows, product *types.Product) error {
	var brandID, categoryID int
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.WeightInKG,
		&product.PiecesPerPackage,
		&product.Image,
		&brandID,
		&categoryID,
	)
	if err != nil {
		return err
	}

	product.Brand = &types.Brand{ID: brandID}
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

func GetBrandsAsMap(s Storage) (map[int]*types.Brand, error) {
	brands, err := s.GetBrands()
	if err != nil {
		return nil, err
	}
	brandsMap := make(map[int]*types.Brand)
	for _, brand := range brands {
		brandsMap[brand.ID] = brand
	}
	return brandsMap, nil
}
