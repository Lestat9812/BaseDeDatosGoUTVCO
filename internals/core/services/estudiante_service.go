package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type EstudianteService struct {
	estudianteRepository ports.IEstudianteRepository
}

func NewEstudianteService(repository ports.IEstudianteRepository) *EstudianteService {
	return &EstudianteService{
		estudianteRepository: repository,
	}
}

func (s *EstudianteService) GuardarEstudiante(estudiante *domains.Estudiante) error {
	err := s.estudianteRepository.SaveEstudiante(estudiante)
	if err != nil {
		return err
	}
	return nil
}

func (s *EstudianteService) ObtenerTodosEstudiantes() ([]*domains.Estudiante, error) {
	res, err := s.estudianteRepository.AllEstudiantes()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *EstudianteService) ObtenerUnEstudiante(id int) (*domains.Estudiante, error) {
	res, err := s.estudianteRepository.FindEstudianteById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *EstudianteService) ActualizarEstudiante(estudiante *domains.Estudiante) error {
	err := s.estudianteRepository.UpdateEstudiante(estudiante)
	if err != nil {
		return err
	}
	return nil
}

func (s *EstudianteService) EliminarEstudiante(id int) error {
	err := s.estudianteRepository.DeleteEstudiante(id)
	if err != nil {
		return err
	}
	return nil
}
