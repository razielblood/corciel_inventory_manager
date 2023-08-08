package types

type Brand struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

func CreateBrand(name string, manufacturer int) *Brand {
	return &Brand{
		Name:         name,
		Manufacturer: &Manufacturer{ID: manufacturer},
	}
}

type CreateBrandRequest struct {
	Name         string `json:"name"`
	Manufacturer int    `json:"manufacturer"`
}

type UpdateBrandRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer int    `json:"manufacturer"`
}

type GetBrandResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer int    `json:"manufacturer"`
}
