package server

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/middlewares"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	alumnoHandler     ports.IAlumnoHandler
	carreraHandler    ports.ICarreraHandler
	estudianteHandler ports.IEstudianteHandler
}

func NewServer(alHdlr ports.IAlumnoHandler, carHdlr ports.ICarreraHandler, estHdlr ports.IEstudianteHandler) *Server {
	return &Server{
		alumnoHandler:     alHdlr,
		carreraHandler:    carHdlr,
		estudianteHandler: estHdlr,
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

	estudiante := app.Group("/estudiante")
	estudiante.Post("/", s.estudianteHandler.NuevoEstudiante)
	estudiante.Get("/todos", s.estudianteHandler.TodosEstudiantes)
	estudiante.Get("/:id", s.estudianteHandler.UnEstudiante)
	estudiante.Put("/:id", s.estudianteHandler.EditarEstudiante)
	estudiante.Delete("/:id", s.estudianteHandler.BorrarEstudiante)

	carrera := app.Group("/carrera")
	carrera.Post("/", s.carreraHandler.NuevaCarrera)
	carrera.Get("/todas", s.carreraHandler.TodasCarreras)
	carrera.Get("/:id", s.carreraHandler.UnaCarrera)
	carrera.Put("/:id", s.carreraHandler.EditarCarrera)
	carrera.Delete("/:id", s.carreraHandler.BorrarCarrera)

	app.Post("/generateAl", middlewares.GenerarLogAlumno)
	app.Post("/generateMa", middlewares.GenerarLogMaestro)
	app.Get("/validate", middlewares.Verificar)
	app.Post("/refreshToken", middlewares.Refrescar)
	app.Get("/", middlewares.Authorizar, func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-Almacén 👋!")
	})
	return app
}
