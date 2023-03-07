package dto

type AllocateInput struct {
	Sku       string `binding:"required"`
	OrderLine string `binding:"required"`
	Qty       int    `binding:"required"`
}
