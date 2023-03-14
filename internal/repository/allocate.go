package repository

import (
	"log"

	entities "github.com/rattapon001/go-rest-api-template/internal/db/entity"
	"gorm.io/gorm"
)

type AllocateRepository interface {
	Find() []entities.Batches
	Save(b *entities.Batches) error
}

type allocateRepository struct {
	DB *gorm.DB
}

func NewAllocateRepository(DB *gorm.DB) *allocateRepository {
	return &allocateRepository{DB: DB}
}

func (r *allocateRepository) Find() []entities.Allocate {
	al := []entities.Allocate{}
	r.DB.Find(&al)
	return al
}

func (r *allocateRepository) Save(al *entities.Allocate) error {
	log.Println(*al)
	err := r.DB.Save(&al).Error
	return err
}
