package income

import (
	"fmt"
	"strconv"
	"sync-finance/util"
)

type Service interface {
	GetData()
	cekListLembaga(institution []Lembaga, institutionName string, total int, method string) []Lembaga
}

type service struct {
	repository Repository
}

func NewLogic(repository Repository) *service {
	return &service{repository}
}

func (s *service) cekListLembaga(institution []Lembaga, institutionName string, total int, method string) []Lembaga {
	statusLoopBreak := false
	for i := range institution {

		if institution[i].Name == institutionName {
			if method == "bntb" {
				bntb := institution[i].BNTB + total
				institution[i].BNTB = bntb
				statusLoopBreak = true
			} else if method == "bsi" {
				bsi := institution[i].BSI + total
				institution[i].BSI = bsi
				statusLoopBreak = true
			} else if method == "cimb" {
				cimb := institution[i].CIMB + total
				institution[i].CIMB = cimb
				statusLoopBreak = true
			} else if method == "tunai" {
				tunai := institution[i].Tunai + total
				institution[i].Tunai = tunai
				statusLoopBreak = true
			}

		}
		if statusLoopBreak {
			break
		}
	}
	if !statusLoopBreak {

		inst_ := Lembaga{
			Name:  institutionName,
			Tunai: 0,
			BNTB:  0,
			CIMB:  0,
			BSI:   0,
		}
		if method == "bntb" {
			inst_.BNTB = total
		} else if method == "bsi" {
			inst_.BSI = total
		} else if method == "cimb" {
			inst_.CIMB = total
		} else if method == "tunai" {
			inst_.Tunai = total
		}
		institution = append(institution, inst_)
	}
	return institution
}

func (s *service) GetData() {
	var lembaga_ []Lembaga
	start := util.FormatToUtc("2024-02-01 00:00:00")
	end := util.FormatToUtc("2024-05-01 23:59:59")
	data := s.repository.GetByDate(start, end)

	for _, d := range data {
		institutionName := util.Institution(strconv.Itoa(d.Nuwb))
		lembaga_ = s.cekListLembaga(lembaga_, institutionName, d.Jumlah, d.Metode_pembayaran)
	}

	fmt.Println(lembaga_)
	fmt.Println(start)
	fmt.Println(end)
}
