package server

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/middlewares"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	alumnoHandler  ports.IAlumnoHandler
	carreraHandler ports.ICarreraHandler
}

func NewServer(alumnoHandler ports.IAlumnoHandler, carreraHandler ports.ICarreraHandler) *Server {
	return &Server{
		alumnoHandler:  alumnoHandler,
		carreraHandler: carreraHandler,
	}
}

func (s *Server) Initizalize() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	alumno := app.Group("/alumno")
	alumno.Post("/", s.alumnoHandler.NuevoAlumno)
	alumno.Get("/todos", s.alumnoHandler.TodosAlumnos)
	alumno.Get("/:id", s.alumnoHandler.UnAlumno)
	alumno.Put("/:id", s.alumnoHandler.EditarAlumno)
	alumno.Delete("/:id", s.alumnoHandler.BorrarAlumno)

	carrera := app.Group("/carrera")
	carrera.Post("/", s.carreraHandler.NuevaCarrera)
	carrera.Get("/todas", s.carreraHandler.TodasCarreras)
	carrera.Get("/:id", s.carreraHandler.UnaCarrera)
	carrera.Put("/:id", s.carreraHandler.EditarCarrera)
	carrera.Delete("/:id", s.carreraHandler.BorrarCarrera)

	app.Post("/generate", middlewares.GenerarLogAlumno)
	app.Get("/validate", middlewares.Verificar)
	app.Post("/refreshToken", middlewares.Refrescar)
	app.Get("/", middlewares.Authorizar, func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-Almacén 👋!")
	})
	return app
}
