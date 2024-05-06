package income

import "gorm.io/gorm"

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
	r.db.Table("pembayaran").Where("createdAt between ? and ?", start, end).Find(&income)
	return income
}
