package types

type Product struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	WeightInKG       float64   `json:"weight_in_kg"`
	PiecesPerPackage int       `json:"pieces_per_package"`
	Image            string    `json:"image"`
	Brand            *Brand    `json:"brand"`
	Category         *Category `json:"category"`
}

func CreateProduct(name string, description string, weightInKG float64, piecesPerPackage int, image string, brand int, category int) *Product {
	return &Product{
		Name:             name,
		Description:      description,
		WeightInKG:       weightInKG,
		PiecesPerPackage: piecesPerPackage,
		Image:            image,
		Brand:            &Brand{ID: brand},
		Category:         &Category{ID: category},
	}
}

type CreateProductRequest struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	WeightInKG       float64 `json:"weight_in_kg"`
	PiecesPerPackage int     `json:"pieces_per_package"`
	Image            string  `json:"image"`
	Brand            int     `json:"brand"`
	Category         int     `json:"category"`
}

type UpdateProductRequest struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	WeightInKG       float64 `json:"weight_in_kg"`
	PiecesPerPackage int     `json:"pieces_per_package"`
	Image            string  `json:"image"`
	Brand            int     `json:"brand"`
	Category         int     `json:"category"`
}

type GetProductResponse struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	WeightInKG       float64 `json:"weight_in_kg"`
	PiecesPerPackage int     `json:"pieces_per_package"`
	Image            string  `json:"image"`
	Brand            int     `json:"brand"`
	Category         int     `json:"category"`
}
