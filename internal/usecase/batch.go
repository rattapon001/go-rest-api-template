package usecase

import (
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/db/entities"
	"github.com/rattapon001/go-rest-api-template/internal/repository"
)

type BatchUseCase interface {
	AddBatch(batchInput *dto.AddBatchInput) (*entities.Batch, error)
}

type batchUseCase struct {
	repository repository.BatchRepository
}

func NewServiceCreate(r repository.BatchRepository) *batchUseCase {
	return &batchUseCase{repository: r}
}

func (u batchUseCase) AddBatch(batchInput *dto.AddBatchInput) (*entities.Batch, error) {
	batch := entities.Batch{}
	batch.Sku = batchInput.Sku
	batch.Qty = batchInput.Qty
	batch.Ref = batchInput.Ref
	batch.Eta = batchInput.Eta

	err := u.repository.Save(&batch)
	return &batch, err
}
