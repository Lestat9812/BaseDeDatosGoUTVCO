package handlers

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
)

type FakerHandler struct {
	fakerService ports.IFakerService
}

func NewFakerHandler(service ports.IFakerService) *FakerHandler {
	return &FakerHandler{
		fakerService: service,
	}
}

func (h *FakerHandler) NuevoFaker(c *fiber.Ctx) error {
	// faker := new()
	// if err := c.BodyParser(faker); err != nil {
	// 	return c.Status(400).JSON(&domains.ErrorResponse{
	// 		Message: "Invalid body parser",
	// 	})
	// }
	err := h.fakerService.GuardarFaker()
	if err != nil {
		return c.Status(404).JSON(&domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(&domains.ResultResponse{
		Message: "Registrado",
	})
}
