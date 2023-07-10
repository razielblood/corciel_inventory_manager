package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/razielblood/corciel_inventory_manager/types"
)

type MariaDBStore struct {
	db *sql.DB
}

func NewMariaDBStore(dbUsername, dbPass, dbHost, dbPort, dbName string) (*MariaDBStore, error) {

	cfg := mysql.Config{
		User:                 dbUsername,
		Passwd:               dbPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", dbHost, dbPort),
		DBName:               dbName,
		AllowNativePasswords: true,
	}

	log.Printf("Connecting to Maria DB using the following info: %+v\n", cfg)

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &MariaDBStore{db: db}, nil
}

func (db *MariaDBStore) CreateCategory(category *types.Category) error {
	fmt.Println("Under construction =)")
	return nil
}
func (db *MariaDBStore) CreateManufacturer(manufacturer *types.Manufacturer) error {
	fmt.Println("Under construction =)")
	return nil
}
func (s *MariaDBStore) CreateProduct(product *types.Product) error {
	query := `insert into Products (Name, Description, WeightInKG, PiecesPerPackage, Image, Manufacturer, Category) 
	values 
	(?, ?, ?, ?, ?, ?, ?)`
	log.Printf("Executed query: %v\n", query)
	_, err := s.db.Query(query, product.Name, product.Description, product.WeightInKG, product.PiecesPerPackage, product.Image, product.Manufacturer.ID, product.Category.ID)
	return err
}
func (db *MariaDBStore) UpdateCategory(product *types.Product) error {
	fmt.Println("Under construction =)")
	return nil
}
func (db *MariaDBStore) UpdateManufacturer(product *types.Product) error {
	fmt.Println("Under construction =)")
	return nil
}
func (db *MariaDBStore) UpdateProduct(product *types.Product) error {
	fmt.Println("Under construction =)")
	return nil
}
func (s *MariaDBStore) GetCategoryByID(categoryID int) (*types.Category, error) {
	query := "select * from Categories where ID = ?"
	rows, err := s.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	category := new(types.Category)
	rows.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	return category, nil
}
func (s *MariaDBStore) GetManufacturerByID(manufacturerID int) (*types.Manufacturer, error) {
	query := "select * from Manufacturers where ID = ?"
	rows, err := s.db.Query(query, manufacturerID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	manufacturer := new(types.Manufacturer)
	rows.Scan(
		&manufacturer.ID,
		&manufacturer.Name,
	)
	return manufacturer, nil
}
func (s *MariaDBStore) GetProductByID(productID int) (*types.Product, error) {
	query := "select * from Products where ID = ?"
	rows, err := s.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	exists := rows.Next()
	if !exists {
		return nil, fmt.Errorf("product with id %v doesn't exists", productID)
	}
	product := new(types.Product)
	s.parseProduct(rows, product)
	return product, nil
}

func (db *MariaDBStore) GetCategories() ([]*types.Category, error) {
	fmt.Println("Under construction =)")
	return nil, nil
}
func (db *MariaDBStore) GetManufacturers() ([]*types.Manufacturer, error) {
	fmt.Println("Under construction =)")
	return nil, nil
}
func (s *MariaDBStore) GetProducts() ([]*types.Product, error) {
	query := "select * from Products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	products := []*types.Product{}
	for rows.Next() {
		product := new(types.Product)
		err := s.parseProduct(rows, product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (s *MariaDBStore) parseProduct(rows *sql.Rows, product *types.Product) error {
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
	manufacturer, err := s.GetManufacturerByID(manufacturerID)
	if err != nil {
		return err
	}
	category, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	product.Manufacturer = manufacturer
	product.Category = category
	return nil
}
