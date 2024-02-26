package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type PersonalService struct {
	personalRepository ports.IPersonalRepository
}

func NewPersonalService(repository ports.IPersonalRepository) *PersonalService {
	return &PersonalService{
		personalRepository: repository,
	}
}

func (s *PersonalService) GuardarPersonal(personal *domains.Personal) error {
	err := s.personalRepository.SavePersonal(personal)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonalService) ObtenerTodosPersonal() ([]*domains.Personal, error) {
	res, err := s.personalRepository.AllPersonal()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) ObtenerUnPersonal(id int) (*domains.Personal, error) {
	res, err := s.personalRepository.FindPersonalById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) ActualizarPersonal(personal *domains.Personal) error {
	err := s.personalRepository.UpdatePersonal(personal)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonalService) EliminarPersonal(id int) error {
	err := s.personalRepository.DeletePersonal(id)
	if err != nil {
		return err
	}
	return nil
}
