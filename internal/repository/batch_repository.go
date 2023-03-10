package repository

import (
	"github.com/rattapon001/go-rest-api-template/internal/db/entities"
	"gorm.io/gorm"
)

type BatchRepo struct {
	DB *gorm.DB
}

func (r *BatchRepo) Find() []entities.BatchEntity {
	batch := []entities.BatchEntity{}
	r.DB.Find(&batch)
	return batch
}

func (r *BatchRepo) Save(b entities.BatchEntity) (entities.BatchEntity, error) {
	err := r.DB.Save(&b).Error
	return b, err
}
