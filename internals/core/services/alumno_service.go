package services

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
)

type AlumnoService struct {
	alumnoRepository ports.IAlumnoRepository
}

func NewAlumnoService(repository ports.IAlumnoRepository) *AlumnoService {
	return &AlumnoService{
		alumnoRepository: repository,
	}
}

func (s *AlumnoService) GuardarAlumno(alumno *domains.Alumno) error {
	err := s.alumnoRepository.SaveAlumno(alumno)
	if err != nil {
		return err
	}
	return nil
}

func (s *AlumnoService) ObtenerTodosAlumnos() ([]*domains.AlumnoSinNada, error) {
	res, err := s.alumnoRepository.AllAlumnos()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AlumnoService) ObtenerUnAlumno(id int) (*domains.Alumno, error) {
	res, err := s.alumnoRepository.FindAlumnoById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *AlumnoService) ObtenerUnAlumnoGrupo(id int) (*domains.Alumno, error) {
	res, err := s.alumnoRepository.FindAlumnoGrupoById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *AlumnoService) ObtenerUnAlumnoCalificaciones(id int) (*domains.Alumno, error) {
	res, err := s.alumnoRepository.FindAlumnoCalificacionesById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AlumnoService) ActualizarAlumno(alumno *domains.Alumno) error {
	err := s.alumnoRepository.UpdateAlumno(alumno)
	if err != nil {
		return err
	}
	return nil
}

func (s *AlumnoService) EliminarAlumno(id int) error {
	err := s.alumnoRepository.DeleteAlumno(id)
	if err != nil {
		return err
	}
	return nil
}
