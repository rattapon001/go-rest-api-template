package repository

import (
	"log"

	"github.com/rattapon001/go-rest-api-template/internal/db/entity"
	"gorm.io/gorm"
)

type BatchRepository interface {
	Find() []entity.Batches
	Save(b *entity.Batches) error
	FindForAllocate(b *[]entity.Batches, qty int, sku string) error
}

type batchRepository struct {
	DB *gorm.DB
}

func NewBatchRepository(DB *gorm.DB) *batchRepository {
	return &batchRepository{DB: DB}
}

func (r *batchRepository) Find() []entity.Batches {
	batch := []entity.Batches{}
	r.DB.Find(&batch)
	return batch
}

func (r *batchRepository) Save(b *entity.Batches) error {
	log.Println(*b)
	err := r.DB.Save(&b).Error
	return err
}

func (r *batchRepository) FindForAllocate(b *[]entity.Batches, qty int, sku string) error {
	err := r.DB.Where("qty >= ? and sku = ?", qty, sku).Find(&b).Error
	return err
}
