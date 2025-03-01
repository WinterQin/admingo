package models

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
}

func (p *Product) ModelName() string {
	return "products"
}

func (p *Product) NewInstance() interface{} {
	return &Product{}
}

func (p *Product) GetPrimaryKey() string {
	return "id"
}
