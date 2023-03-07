package dto

type AddBatchInput struct {
	Sku string `binding:"required"`
	Ref string `binding:"required"`
	Qty int    `binding:"required"`
	Eta string
}
