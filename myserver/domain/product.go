package domain

type Product struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	Price    int32  `json:"price"`
	Amount   int32  `json:"amount"`
}