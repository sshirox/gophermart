package entity

type Order struct {
	Number  ID       `json:"number" required:"true"`
	Status  string   `json:"status" required:"true"`
	Accrual *float64 `json:"accrual,omitempty"`
}
