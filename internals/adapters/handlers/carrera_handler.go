package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
)

type CarreraHandler struct {
	carreraService ports.ICarreraService
}

func NewCarreraHandler(service ports.ICarreraService) *CarreraHandler {
	return &CarreraHandler{
		carreraService: service,
	}
}

func (h *CarreraHandler) NuevaCarrera(c *fiber.Ctx) error {
	carrera := new(domains.Carrera)
	if err := c.BodyParser(carrera); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	err := h.carreraService.GuardarCarrera(carrera)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(&domains.ResultResponse{
		Message: "Registrado",
	})
}

func (h *CarreraHandler) TodasCarreras(c *fiber.Ctx) error {
	res, err := h.carreraService.ObtenerTodasCarreras()
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *CarreraHandler) UnaCarrera(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	res, err := h.carreraService.ObtenerUnaCarrera(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": &res,
	})
}

func (h *CarreraHandler) EditarCarrera(c *fiber.Ctx) error {
	carrera := new(domains.Carrera)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid bid",
		})
	}
	if err := c.BodyParser(carrera); err != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid body parser",
		})
	}
	carrera.ID = uint(id)
	error := h.carreraService.ActualizarCarrera(carrera)
	if error != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Actualizado",
	})
}

func (h *CarreraHandler) BorrarCarrera(c *fiber.Ctx) error {
	id, error := c.ParamsInt("id")
	if error != nil {
		return c.Status(400).JSON(&domains.ErrorResponse{
			Message: "Invalid id",
		})
	}
	err := h.carreraService.EliminarCarrera(id)
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: error.Error(),
		})
	}

	return c.Status(200).JSON(&domains.ResultResponse{
		Message: "Eliminado",
	})
}
