package income

import "fmt"

type Service interface{}

type service struct {
	repository Repository
}

func NewLogic(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetData() {
	data := s.repository.GetByDate("2024-02-01 00:00:00", "2024-02-01 23:59:59")
	// bntb := 0
	// bsi := 0
	// cimb := 0

	for _, d := range data {
		fmt.Println(d)
		fmt.Println("+++")
	}
}
