package income

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetByDate(start string, end string) []Income
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByDate(start string, end string) []Income {
	var income []Income
	err := r.db.Table("pembayaran").Where("createdAt between ? and ?", start, end).Find(&income).Error
	fmt.Println(err)
	return income
}
