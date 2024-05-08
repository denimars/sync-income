package income

import (
	"time"
)

type Income struct {
	Uuid              string
	Nuwb              int
	Ket               string
	Jumlah            int
	Tahun_ajar        string
	Metode_pembayaran string
	Uuid_santri       string
	Uuid_user         string
	CreatedAt         time.Time `gorm:"column:createdAt"`
	UpdatedAt         time.Time `gorm:"column:updatedAt"`
}

type Lembaga struct {
	BSI   int
	Tunai int
	BNTB  int
	CIMB  int
	Name  string
}
