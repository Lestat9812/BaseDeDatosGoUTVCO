package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
)

type EstudianteHandler struct {
	estudianteService ports.IEstudianteService
}

func NewEstudianteHandler(service ports.IEstudianteService) *EstudianteHandler {
	return &EstudianteHandler{
		estudianteService: service,
	}
}

func (h *EstudianteHandler) NuevoEstudiante(c *fiber.Ctx) error {
	estudiante := new(domains.Estudiante)
	if err := c.BodyParser(estudiante); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	err := h.estudianteService.GuardarEstudiante(estudiante)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(&domains.ResultResponse{
		Message: "Registrado",
	})
}

func (h *EstudianteHandler) TodosEstudiantes(c *fiber.Ctx) error {
	res, err := h.estudianteService.ObtenerTodosEstudiantes()
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *EstudianteHandler) UnEstudiante(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	res, err := h.estudianteService.ObtenerUnEstudiante(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *EstudianteHandler) EditarEstudiante(c *fiber.Ctx) error {
	estudiante := new(domains.Estudiante)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid bid",
		})
	}
	if err := c.BodyParser(estudiante); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	estudiante.ID = uint(id)
	error := h.estudianteService.ActualizarEstudiante(estudiante)
	if error != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Actualizado",
	})
}

func (h *EstudianteHandler) BorrarEstudiante(c *fiber.Ctx) error {
	id, error := c.ParamsInt("id")
	if error != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	err := h.estudianteService.EliminarEstudiante(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Eliminado",
	})
}
