package repository

import (
	"log"

	"github.com/rattapon001/go-rest-api-template/internal/db/entities"
	"gorm.io/gorm"
)

type BatchRepository interface {
	Find() []entities.Batch
	Save(b *entities.Batch) error
}

type batchRepository struct {
	DB *gorm.DB
}

func NewBatchRepository(DB *gorm.DB) *batchRepository {
	return &batchRepository{DB: DB}
}

func (r *batchRepository) Find() []entities.Batch {
	batch := []entities.Batch{}
	r.DB.Find(&batch)
	return batch
}

func (r *batchRepository) Save(b *entities.Batch) error {
	log.Println(*b)
	err := r.DB.Save(&b).Error
	return err
}
