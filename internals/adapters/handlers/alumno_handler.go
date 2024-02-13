package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AlumnoHandler struct {
	alumnoService ports.IAlumnoService
}

func NewAlumnoHandler(service ports.IAlumnoService) *AlumnoHandler {
	return &AlumnoHandler{
		alumnoService: service,
	}
}

func (h *AlumnoHandler) NuevoAlumno(c *fiber.Ctx) error {
	alumno := new(domains.Alumno)
	if err := c.BodyParser(alumno); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	err := h.alumnoService.GuardarAlumno(alumno)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(&domains.ResultResponse{
		Message: "Registrado",
	})
}

func (h *AlumnoHandler) TodosAlumnos(c *fiber.Ctx) error {
	res, err := h.alumnoService.ObtenerTodosAlumnos()
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *AlumnoHandler) UnAlumno(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	res, err := h.alumnoService.ObtenerUnAlumno(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *AlumnoHandler) EditarAlumno(c *fiber.Ctx) error {
	alumno := new(domains.Alumno)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid bid",
		})
	}
	if err := c.BodyParser(alumno); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	alumno.ID = uint(id)
	error := h.alumnoService.ActualizarAlumno(alumno)
	if error != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Actualizado",
	})
}

func (h *AlumnoHandler) BorrarAlumno(c *fiber.Ctx) error {
	id, error := c.ParamsInt("id")
	if error != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	err := h.alumnoService.EliminarAlumno(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Eliminado",
	})
}
