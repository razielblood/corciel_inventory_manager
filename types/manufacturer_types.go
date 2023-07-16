package types

type Manufacturer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateManufacturer(name string) *Manufacturer {
	return &Manufacturer{
		Name: name,
	}
}

type CreateManufacturerRequest struct {
	Name string `json:"name"`
}

type UpdateManufacturerRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
