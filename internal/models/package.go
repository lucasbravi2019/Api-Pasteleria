package models

type Package struct {
	Id       int64
	Metric   string
	Quantity float64
}

func NewPackage(id int64, metric string, quantity float64) *Package {
	return &Package{
		Id:       id,
		Metric:   metric,
		Quantity: quantity,
	}
}
