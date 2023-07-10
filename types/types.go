package types

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(name, description string) *Category {
	return &Category{
		Name:        name,
		Description: description,
	}
}

type Manufacturer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateManufacturer(name string) *Manufacturer {
	return &Manufacturer{
		Name: name,
	}
}

type Product struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	WeightInKG       float64       `json:"weight_in_kg"`
	PiecesPerPackage int           `json:"pieces_per_package"`
	Image            string        `json:"image"`
	Manufacturer     *Manufacturer `json:"manufacturer"`
	Category         *Category     `json:"category"`
}

func CreateProduct(name string, description string, weightInKG float64, piecesPerPackage int, image string, manufacturer int, category int) *Product {
	return &Product{
		Name:             name,
		Description:      description,
		WeightInKG:       weightInKG,
		PiecesPerPackage: piecesPerPackage,
		Image:            image,
		Manufacturer:     &Manufacturer{ID: manufacturer},
		Category:         &Category{ID: category},
	}
}

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateManufacturerRequest struct {
	Name string `json:"name"`
}

type CreateProductRequest struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	WeightInKG       float64 `json:"weight_in_kg"`
	PiecesPerPackage int     `json:"pieces_per_package"`
	Image            string  `json:"image"`
	Manufacturer     int     `json:"manufacturer"`
	Category         int     `json:"category"`
}

type GetProductResponse struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	WeightInKG       float64 `json:"weight_in_kg"`
	PiecesPerPackage int     `json:"pieces_per_package"`
	Image            string  `json:"image"`
	Manufacturer     int     `json:"manufacturer"`
	Category         int     `json:"category"`
}

type APIError struct {
	Message string `json:"message"`
}

type User struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func CreateLoginRequest(username, password string) *LoginRequest {
	return &LoginRequest{Username: username, Password: password}
}

type CreateUserRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Email     string `form:"email" json:"email"`
}

func CreateCreateUserRequest(username, password, firstName, lastName, email string) *CreateUserRequest {
	return &CreateUserRequest{Username: username, Password: password, FirstName: firstName, LastName: lastName, Email: email}
}
