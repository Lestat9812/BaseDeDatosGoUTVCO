package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

// "gitlab.com/l-cm/api-giras/internals/core/domains"
// "gitlab.com/l-cm/api-giras/internals/core/ports"

type AlumnoService struct {
	encargadoRepository ports.IAlumnoRepository
}

func NewAlumnoService(repository ports.IAlumnoRepository) *AlumnoService {
	return &AlumnoService{
		encargadoRepository: repository,
	}
}

func (s *AlumnoService) GuardarAlumno(encargado *domains.Alumno) error {
	err := s.encargadoRepository.SaveAlumno(encargado)
	if err != nil {
		return err
	}
	return nil
}

func (s *AlumnoService) ObtenerTodosAlumnos() ([]*domains.Alumno, error) {
	res, err := s.encargadoRepository.AllAlumnos()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AlumnoService) ObtenerUnAlumno(id int) (*domains.Alumno, error) {
	res, err := s.encargadoRepository.FindAlumnoById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AlumnoService) ActualizarAlumno(encargado *domains.Alumno) error {
	err := s.encargadoRepository.UpdateAlumno(encargado)
	if err != nil {
		return err
	}
	return nil
}

func (s *AlumnoService) EliminarAlumno(id int) error {
	err := s.encargadoRepository.DeleteAlumno(id)
	if err != nil {
		return err
	}
	return nil
}
