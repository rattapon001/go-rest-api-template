package entity

type Batches struct {
	Id  uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Sku string `json:"sku"`
	Ref string `json:"ref"`
	Qty int    `json:"qty"`
	Eta string `json:"eta"`
}
