package repository

import (
	entities "github.com/rattapon001/go-rest-api-template/internal/db/entity"
	"gorm.io/gorm"
)

type AllocateRepository interface {
	Find() []entities.Allocate
	Save(b *entities.Allocate) error
	GetAllocateDB() *gorm.DB
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
	err := r.DB.Save(&al).Error
	return err
}

func (r *allocateRepository) GetAllocateDB() *gorm.DB {
	return r.DB
}
