package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type FakerService struct {
	fakerRepository ports.IFakerRepository
}

func NewFakerService(repository ports.IFakerRepository) *FakerService {
	return &FakerService{
		fakerRepository: repository,
	}
}

func (s *FakerService) GuardarFaker() error {
	err := s.fakerRepository.SaveFaker()
	if err != nil {
		return err
	}
	return nil
}
