package services

import "github.com/awwabi/monorepo-academics/cuti/internal/repositories"

type KelasService struct {
	kelasRepo repositories.KelasRepository
}

func NewKelasService(kelasRepo repositories.KelasRepository) *KelasService {
	return &KelasService{kelasRepo}
}

func (s *KelasService) GetKelas() ([]repositories.KelasSimple, error) {
	return s.kelasRepo.GetKelas()
}
