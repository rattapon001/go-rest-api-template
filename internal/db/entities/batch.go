package entities

type Batch struct {
	Id  uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Sku string `json:"sku"`
	Ref string `json:"ref"`
	Qty int    `json:"qty"`
	Eta string `json:"eta"`
}

type Tabler interface {
	TableName() string
}

func (Batch) TableName() string {
	return "batch"
}
