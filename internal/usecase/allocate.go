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

type Test struct {
	Qty int
}

func NewAllocateUseCase(r repository.AllocateRepository) *allocateUseCase {
	return &allocateUseCase{repository: r}
}

func (u *allocateUseCase) AllocateCreate(allocateInput *dto.AllocateInput) (*entity.Allocate, error) {
	log.Println(allocateInput)
	var err error = nil
	batchRepo := repository.NewBatchRepository(u.repository.GetAllocateDB())
	batches := []entity.Batches{}
	batchRepo.FindForAllocate(&batches, allocateInput.Qty, allocateInput.Sku)
	log.Println(batches)
	allocate := entity.Allocate{}
	if len(batches) > 0 {
		batch := batches[0]
		batch.Qty = batch.Qty - allocateInput.Qty
		allocate.Sku = allocateInput.Sku
		allocate.Qty = allocateInput.Qty
		allocate.OrderLine = allocateInput.OrderLine
		err = u.repository.Save(&allocate)
		if err != nil {
			log.Println(err)
		}
		err = batchRepo.Save(&batch)
		if err != nil {
			log.Println(err)
		}

	}
	return &allocate, err
}
