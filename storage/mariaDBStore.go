package storage

import (
	"database/sql"
	"fmt"

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

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &MariaDBStore{db: db}, nil
}

func (s *MariaDBStore) CreateCategory(category *types.Category) error {
	query := "insert into Categories (Name, Description) values (?, ?)"
	_, err := s.db.Query(query, category.Name, category.Description)
	return err
}
func (s *MariaDBStore) CreateManufacturer(manufacturer *types.Manufacturer) error {
	query := "insert into Manufacturers (Name) values (?)"
	_, err := s.db.Query(query, manufacturer.Name)
	return err
}
func (s *MariaDBStore) CreateProduct(product *types.Product) error {
	query := `insert into Products (Name, Description, WeightInKG, PiecesPerPackage, Image, Manufacturer, Category) 
	values 
	(?, ?, ?, ?, ?, ?, ?)`
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
	parseProduct(rows, product)

	product.Manufacturer, _ = s.GetManufacturerByID(product.Manufacturer.ID)
	product.Category, _ = s.GetCategoryByID(product.Category.ID)

	return product, nil
}

func (s *MariaDBStore) GetCategories() ([]*types.Category, error) {
	query := "select (ID, Name, Description) from Categories"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	categories := []*types.Category{}
	for rows.Next() {
		category := new(types.Category)
		err := rows.Scan(&category.ID,
			&category.Name,
			&category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func (s *MariaDBStore) GetManufacturers() ([]*types.Manufacturer, error) {
	query := "select (ID, Name) from Manufacturers"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	manufacturers := []*types.Manufacturer{}
	for rows.Next() {
		manufacturer := new(types.Manufacturer)
		err := rows.Scan(&manufacturer.ID,
			&manufacturer.Name)
		if err != nil {
			return nil, err
		}
		manufacturers = append(manufacturers, manufacturer)
	}
	return manufacturers, nil
}
func (s *MariaDBStore) GetProducts() ([]*types.Product, error) {
	query := "select * from Products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	categoriesMap, err := GetCategoriesAsMap(s)
	if err != nil {
		return nil, err
	}
	manufacturersMap, err := GetManufacturersAsMap(s)
	if err != nil {
		return nil, err
	}
	products := []*types.Product{}
	for rows.Next() {
		product := new(types.Product)
		err := parseProduct(rows, product)
		if err != nil {
			return nil, err
		}

		product.Manufacturer = manufacturersMap[product.Manufacturer.ID]
		product.Category = categoriesMap[product.Category.ID]

		products = append(products, product)
	}
	return products, nil
}

func (s *MariaDBStore) LoginUser(loginRequest *types.LoginRequest) (*types.User, error) {
	query := "select Username, FirstName, LastName, Email from Users where Username = ? and Password= ?"
	rows, err := s.db.Query(query, loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	exists := rows.Next()
	if !exists {
		return nil, fmt.Errorf("username %v doesn't exists or the password is incorrect", loginRequest.Username)
	}
	user := new(types.User)
	rows.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)
	return user, nil
}

func (s *MariaDBStore) CreateUser(createUserRequest *types.CreateUserRequest) (*types.User, error) {
	query := "select Username from Users where Username = ?"
	rows, err := s.db.Query(query, createUserRequest.Username)
	if err != nil {
		return nil, err
	}
	exists := rows.Next()
	if exists {
		return nil, fmt.Errorf("username %v already exists", createUserRequest.Username)
	}

	query = "insert into Users (Username, Password, FirstName, LastName, Email) values (?, ?, ?, ?, ?)"
	_, err = s.db.Query(query, createUserRequest.Username, createUserRequest.Password, createUserRequest.FirstName, createUserRequest.LastName, createUserRequest.Email)
	if err != nil {
		return nil, err
	}

	query = "select Username, FirstName, LastName, Email from Users where Username = ?"
	rows, err = s.db.Query(query, createUserRequest.Username)
	if err != nil {
		return nil, err
	}
	exists = rows.Next()
	if !exists {
		return nil, fmt.Errorf("user %v couldn't be created", createUserRequest.Username)
	}
	user := new(types.User)
	rows.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)

	return user, nil
}
