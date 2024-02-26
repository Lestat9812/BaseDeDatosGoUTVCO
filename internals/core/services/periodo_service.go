package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type PeriodoService struct {
	periodoRepository ports.IPeriodoRepository
}

func NewPeriodoService(repository ports.IPeriodoRepository) *PeriodoService {
	return &PeriodoService{
		periodoRepository: repository,
	}
}

func (s *PeriodoService) GuardarPeriodo(periodo *domains.Periodo) error {
	err := s.periodoRepository.SavePeriodo(periodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *PeriodoService) ObtenerTodosPeriodos() ([]*domains.Periodo, error) {
	res, err := s.periodoRepository.AllPeriodos()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PeriodoService) ObtenerUnPeriodo(id int) (*domains.Periodo, error) {
	res, err := s.periodoRepository.FindPeriodoById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PeriodoService) ActualizarPeriodo(periodo *domains.Periodo) error {
	err := s.periodoRepository.UpdatePeriodo(periodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *PeriodoService) EliminarPeriodo(id int) error {
	err := s.periodoRepository.DeletePeriodo(id)
	if err != nil {
		return err
	}
	return nil
}
