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

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MariaDBStore{db: db}, nil
}

func (s *MariaDBStore) CreateCategory(category *types.Category) error {
	query := "insert into Categories (Name, Description) values (?, ?) returning ID"
	result, err := s.db.Query(query, category.Name, category.Description)
	if err != nil {
		return err
	}
	defer result.Close()
	result.Next()
	result.Scan(&category.ID)
	if category.ID == 0 {
		return fmt.Errorf("error creating category")
	}
	return nil
}
func (s *MariaDBStore) CreateManufacturer(manufacturer *types.Manufacturer) error {
	query := "insert into Manufacturers (Name) values (?) returning ID"
	result, err := s.db.Query(query, manufacturer.Name)
	result.Next()
	result.Scan(&manufacturer.ID)
	result.Close()
	return err
}
func (s *MariaDBStore) CreateProduct(product *types.Product) error {
	query := `insert into Products (Name, Description, WeightInKG, PiecesPerPackage, Image, Brand, Category) 
	values 
	(?, ?, ?, ?, ?, ?, ?) returning ID`
	result, err := s.db.Query(query, product.Name, product.Description, product.WeightInKG, product.PiecesPerPackage, product.Image, product.Brand.ID, product.Category.ID)
	result.Next()
	result.Scan(&product.ID)
	result.Close()
	return err
}

func (s *MariaDBStore) CreateBrand(brand *types.Brand) error {
	query := `insert into Brands (Name, Manufacturer) values (?, ?) returning ID`
	result, err := s.db.Query(query, brand.Name, brand.Manufacturer.ID)
	if err != nil {
		return err
	}
	defer result.Close()
	result.Next()
	result.Scan(&brand.ID)
	if brand.ID == 0 {
		return fmt.Errorf("error creating brand")
	}
	return nil
}

func (s *MariaDBStore) UpdateCategory(category *types.Category) error {
	query := "update Categories set Name = ?, Description = ? where ID = ?"
	_, err := s.db.Query(query, category.Name, category.Description, category.ID)
	return err
}
func (s *MariaDBStore) UpdateManufacturer(manufacturer *types.Manufacturer) error {
	query := "update Manufacturers set Name = ? where ID = ?"
	_, err := s.db.Query(query, manufacturer.Name, manufacturer.ID)

	return err
}
func (s *MariaDBStore) UpdateProduct(product *types.Product) error {
	query := "update Products set Name = ?, Description = ?, WeightInKG = ?, PiecesPerPackage = ?, Image = ?, Brand = ?, Category = ? where ID = ?"
	_, err := s.db.Query(query, product.Name, product.Description, product.WeightInKG, product.PiecesPerPackage, product.Image, product.Brand.ID, product.Category.ID, product.ID)

	return err
}

func (s *MariaDBStore) UpdateBrand(brand *types.Brand) error {
	query := "update Brands set Name = ?, Manufacturer = ? where ID = ?"
	_, err := s.db.Query(query, brand.Name, brand.Manufacturer.ID, brand.ID)
	return err
}

func (s *MariaDBStore) GetCategoryByID(categoryID int) (*types.Category, error) {
	query := "select * from Categories where ID = ?"
	rows, err := s.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	exists := rows.Next()
	if !exists {
		return nil, fmt.Errorf("category with id %v doesn't exists", categoryID)
	}
	category := new(types.Category)
	rows.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	rows.Close()
	return category, nil
}
func (s *MariaDBStore) GetManufacturerByID(manufacturerID int) (*types.Manufacturer, error) {
	query := "select * from Manufacturers where ID = ?"
	rows, err := s.db.Query(query, manufacturerID)
	if err != nil {
		return nil, err
	}
	exists := rows.Next()
	if !exists {
		return nil, fmt.Errorf("manufacturer with id %v doesn't exists", manufacturerID)
	}
	manufacturer := new(types.Manufacturer)
	rows.Scan(
		&manufacturer.ID,
		&manufacturer.Name,
	)
	rows.Close()
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

	product.Brand, _ = s.GetBrandByID(product.Brand.ID)
	product.Category, _ = s.GetCategoryByID(product.Category.ID)
	rows.Close()
	return product, nil
}

func (s *MariaDBStore) GetBrandByID(brandID int) (*types.Brand, error) {
	query := "select ID, Name, Manufacturer from Brands where ID = ?"
	result, err := s.db.Query(query, brandID)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	exists := result.Next()
	if !exists {
		return nil, fmt.Errorf("brand with id %v doesn't exists", brandID)

	}
	manufacturersMap, err := GetManufacturersAsMap(s)
	if err != nil {
		return nil, err
	}

	brand := new(types.Brand)
	err = parseBrand(result, brand)
	if err != nil {
		return nil, err
	}

	brand.Manufacturer = manufacturersMap[brand.Manufacturer.ID]

	return brand, nil
}

func (s *MariaDBStore) GetCategories() ([]*types.Category, error) {
	query := "select ID, Name, Description from Categories"
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
	rows.Close()
	return categories, nil
}
func (s *MariaDBStore) GetManufacturers() ([]*types.Manufacturer, error) {
	query := "select ID, Name from Manufacturers"
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
	rows.Close()
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
	brandsMap, err := GetBrandsAsMap(s)
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

		product.Brand = brandsMap[product.Brand.ID]
		product.Category = categoriesMap[product.Category.ID]

		products = append(products, product)
	}
	rows.Close()
	return products, nil
}

func (s *MariaDBStore) GetBrands() ([]*types.Brand, error) {
	query := "select ID, Name, Manufacturer from Brands"
	results, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	manufacturersMap, err := GetManufacturersAsMap(s)
	if err != nil {
		return nil, err
	}
	brands := []*types.Brand{}
	for results.Next() {
		brand := new(types.Brand)
		err := parseBrand(results, brand)
		if err != nil {
			return nil, err
		}
		brand.Manufacturer = manufacturersMap[brand.Manufacturer.ID]
		brands = append(brands, brand)
	}
	return brands, nil
}

func (s *MariaDBStore) GetUserByID(id string) (*types.User, error) {
	query := "select Username, FirstName, LastName, Email, Password from Users where Username = ?"
	rows := s.db.QueryRow(query, id)
	user := new(types.User)
	err := rows.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("username %v doesn't exist", id)
	}

	return user, nil
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
	rows.Close()
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
	rows.Close()
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
	rows.Close()
	return user, nil
}
