package dto

type PackageDTO struct {
	Id       *int64   `json:"id,omitempty"`
	Metric   *string  `json:"metric,omitempty"`
	Quantity *float64 `json:"quantity,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

func NewPackageDTO(id *int64, metric *string, quantity *float64, price *float64) *PackageDTO {
	return &PackageDTO{
		Id:       id,
		Metric:   metric,
		Quantity: quantity,
		Price:    price,
	}
}
