package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type UtcampusService struct {
	utcampusRepository ports.IUtcampusRepository
}

func NewUtcampusService(repository ports.IUtcampusRepository) *UtcampusService {
	return &UtcampusService{
		utcampusRepository: repository,
	}
}

func (s *UtcampusService) GuardarUtcampus(utcampus *domains.Utcampus) error {
	err := s.utcampusRepository.SaveUtcampus(utcampus)
	if err != nil {
		return err
	}
	return nil
}

func (s *UtcampusService) ObtenerTodosUtcampus() ([]*domains.Utcampus, error) {
	res, err := s.utcampusRepository.AllUtcampus()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UtcampusService) ObtenerUnUtcampus(id int) (*domains.Utcampus, error) {
	res, err := s.utcampusRepository.FindUtcampusById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UtcampusService) ActualizarUtcampus(utcampus *domains.Utcampus) error {
	err := s.utcampusRepository.UpdateUtcampus(utcampus)
	if err != nil {
		return err
	}
	return nil
}

func (s *UtcampusService) EliminarUtcampus(id int) error {
	err := s.utcampusRepository.DeleteUtcampus(id)
	if err != nil {
		return err
	}
	return nil
}
