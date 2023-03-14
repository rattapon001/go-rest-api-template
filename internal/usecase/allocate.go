package usecase

import (
	"log"

	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/db/entity"
	"github.com/rattapon001/go-rest-api-template/internal/repository"
)

type AllocateUseCase interface {
	AllocateCreate(allocateInput *dto.AllocateInput) (*entity.Allocate, error)
}

type allocateUseCase struct {
	repository repository.AllocateRepository
}

func NewAllocateUseCase(r repository.AllocateRepository) *allocateUseCase {
	return &allocateUseCase{repository: r}
}

func (u *allocateUseCase) AllocateCreate(allocateInput *dto.AllocateInput) (*entity.Allocate, error) {
	log.Println(allocateInput)

	batchRepo := repository.NewBatchRepository(u.repository.GetAllocateDB())
	batch := batchRepo.Find()
	log.Println(batch)
	allocate := entity.Allocate{}
	allocate.Sku = "sdsad"
	allocate.Qty = 1
	allocate.OrderLine = "dsafsad"
	err := u.repository.Save(&allocate)
	return &allocate, err
}
