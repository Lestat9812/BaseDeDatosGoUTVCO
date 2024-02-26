package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type MateriaService struct {
	materiaRepository ports.IMateriaRepository
}

func NewMateriaService(repository ports.IMateriaRepository) *MateriaService {
	return &MateriaService{
		materiaRepository: repository,
	}
}

func (s *MateriaService) GuardarMateria(materia *domains.Materia) error {
	err := s.materiaRepository.SaveMateria(materia)
	if err != nil {
		return err
	}
	return nil
}

func (s *MateriaService) ObtenerTodasMaterias() ([]*domains.Materia, error) {
	res, err := s.materiaRepository.AllMaterias()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MateriaService) ObtenerUnaMateria(id int) (*domains.Materia, error) {
	res, err := s.materiaRepository.FindMateriaById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MateriaService) ActualizarMateria(materia *domains.Materia) error {
	err := s.materiaRepository.UpdateMateria(materia)
	if err != nil {
		return err
	}
	return nil
}

func (s *MateriaService) EliminarMateria(id int) error {
	err := s.materiaRepository.DeleteMateria(id)
	if err != nil {
		return err
	}
	return nil
}
