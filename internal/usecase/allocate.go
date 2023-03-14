package usecase

import (
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

func NewAllocateService(r repository.AllocateRepository) *allocateUseCase {
	return &allocateUseCase{repository: r}
}
