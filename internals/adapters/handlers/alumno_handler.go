package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
	// "gitlab.com/l-cm/api-encargados/internals/core/domains"
	// "gitlab.com/l-cm/api-encargados/internals/core/ports"
)

type AlumnoHandler struct {
	encargadoService ports.IAlumnoService
}

func NewAlumnoHandler(service ports.IAlumnoService) *AlumnoHandler {
	return &AlumnoHandler{
		encargadoService: service,
	}
}

func (h *AlumnoHandler) NuevoAlumno(c *fiber.Ctx) error {
	don := new(domains.Alumno)
	if err := c.BodyParser(don); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	err := h.encargadoService.GuardarAlumno(don)
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
	res, err := h.encargadoService.ObtenerTodosAlumnos()
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
	res, err := h.encargadoService.ObtenerUnAlumno(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *AlumnoHandler) EditaAlumno(c *fiber.Ctx) error {
	encargado := new(domains.Alumno)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid bid",
		})
	}
	if err := c.BodyParser(encargado); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	encargado.ID = uint(id)
	error := h.encargadoService.ActualizarAlumno(encargado)
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
	err := h.encargadoService.EliminarAlumno(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Eliminado",
	})
}
