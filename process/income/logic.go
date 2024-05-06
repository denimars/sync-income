package income

import (
	"fmt"
	"sync-finance/util"
)

type Service interface{}

type service struct {
	repository Repository
}

func NewLogic(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetData() {
	start := util.FormatToUtc("2024-02-01 00:00:00")
	end := util.FormatToUtc("2024-02-01 23:59:59")

	data := s.repository.GetByDate(start, end)
	// bntb := 0
	// bsi := 0
	// cimb := 0

	for _, d := range data {
		fmt.Println(d.CreatedAt)
		fmt.Println("+++")
	}
	fmt.Println(start)
	fmt.Println(end)
}
