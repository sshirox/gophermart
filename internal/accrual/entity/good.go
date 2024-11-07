package entity

type Good struct {
	Description string  `json:"description" required:"true"`
	Price       float64 `json:"price" required:"true"`
}
