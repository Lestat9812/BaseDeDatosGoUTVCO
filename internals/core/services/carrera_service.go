package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type CarreraService struct {
	carreraRepository ports.ICarreraRepository
}

func NewCarreraService(repository ports.ICarreraRepository) *CarreraService {
	return &CarreraService{
		carreraRepository: repository,
	}
}

func (s *CarreraService) GuardarCarrera(carrera *domains.Carrera) error {
	err := s.carreraRepository.SaveCarrera(carrera)
	if err != nil {
		return err
	}
	return nil
}

func (s *CarreraService) ObtenerTodasCarreras() ([]*domains.Carrera, error) {
	res, err := s.carreraRepository.AllCarreras()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CarreraService) ObtenerUnaCarrera(id int) (*domains.Carrera, error) {
	res, err := s.carreraRepository.FindCarreraById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CarreraService) ActualizarCarrera(carrera *domains.Carrera) error {
	err := s.carreraRepository.UpdateCarrera(carrera)
	if err != nil {
		return err
	}
	return nil
}

func (s *CarreraService) EliminarCarrera(id int) error {
	err := s.carreraRepository.DeleteCarrera(id)
	if err != nil {
		return err
	}
	return nil
}
