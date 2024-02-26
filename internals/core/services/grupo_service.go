package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type GrupoService struct {
	grupoRepository ports.IGrupoRepository
}

func NewGrupoService(repository ports.IGrupoRepository) *GrupoService {
	return &GrupoService{
		grupoRepository: repository,
	}
}

func (s *GrupoService) GuardarGrupo(grupo *domains.Grupo) error {
	err := s.grupoRepository.SaveGrupo(grupo)
	if err != nil {
		return err
	}
	return nil
}

func (s *GrupoService) ObtenerTodosGrupos() ([]*domains.Grupo, error) {
	res, err := s.grupoRepository.AllGrupos()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GrupoService) ObtenerUnGrupo(id int) (*domains.Grupo, error) {
	res, err := s.grupoRepository.FindGrupoById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GrupoService) ActualizarGrupo(grupo *domains.Grupo) error {
	err := s.grupoRepository.UpdateGrupo(grupo)
	if err != nil {
		return err
	}
	return nil
}

func (s *GrupoService) EliminarGrupo(id int) error {
	err := s.grupoRepository.DeleteGrupo(id)
	if err != nil {
		return err
	}
	return nil
}
