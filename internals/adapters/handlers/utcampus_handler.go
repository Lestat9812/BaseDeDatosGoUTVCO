package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UtcampusHandler struct {
	utcampusService ports.IUtcampusService
}

func NewUtcampusHandler(service ports.IUtcampusService) *UtcampusHandler {
	return &UtcampusHandler{
		utcampusService: service,
	}
}

func (h *UtcampusHandler) NuevoUtcampus(c *fiber.Ctx) error {
	utcampus := new(domains.Utcampus)
	if err := c.BodyParser(utcampus); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	err := h.utcampusService.GuardarUtcampus(utcampus)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(&domains.ResultResponse{
		Message: "Registrado",
	})
}

func (h *UtcampusHandler) TodosUtcampus(c *fiber.Ctx) error {
	res, err := h.utcampusService.ObtenerTodosUtcampus()
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *UtcampusHandler) UnUtcampus(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	res, err := h.utcampusService.ObtenerUnUtcampus(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *UtcampusHandler) EditarUtcampus(c *fiber.Ctx) error {
	utcampus := new(domains.Utcampus)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid bid",
		})
	}
	if err := c.BodyParser(utcampus); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	utcampus.ID = uint(id)
	error := h.utcampusService.ActualizarUtcampus(utcampus)
	if error != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Actualizado",
	})
}

func (h *UtcampusHandler) BorrarUtcampus(c *fiber.Ctx) error {
	id, error := c.ParamsInt("id")
	if error != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	err := h.utcampusService.EliminarUtcampus(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Eliminado",
	})
}
