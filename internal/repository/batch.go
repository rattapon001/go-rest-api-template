package repository

import (
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
	return r.DB.Save(&b).Error
}
func (r *batchRepository) FindForAllocate(b *[]entity.Batches, qty int, sku string) error {
	return r.DB.Where("qty >= ? and sku = ?", qty, sku).Order("qty asc").Find(&b).Error
}

func (r *batchRepository) UpdateQty(qty int) error {
	return r.DB.Model(&entity.Batches{}).Update("qty", qty).Error
}
