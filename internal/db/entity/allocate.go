package entity

type Allocate struct {
	Id        uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Sku       string `json:"sku"`
	OrderLine string `json:"order_line"`
	Qty       int    `json:"qty"`
}
