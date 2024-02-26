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
	grupoHandler      ports.IGrupoHandler
	periodoHandler    ports.IPeriodoHandler
	personalHandler   ports.IPersonalHandler
	materiaHandler    ports.IMateriaHandler
}

func NewServer(alHdlr ports.IAlumnoHandler, carHdlr ports.ICarreraHandler,
	estHdlr ports.IEstudianteHandler, grHdlr ports.IGrupoHandler,
	peHdlr ports.IPeriodoHandler, matHdlr ports.IMateriaHandler,
	persHdlr ports.IPersonalHandler) *Server {
	return &Server{
		alumnoHandler:     alHdlr,
		carreraHandler:    carHdlr,
		estudianteHandler: estHdlr,
		grupoHandler:      grHdlr,
		periodoHandler:    peHdlr,
		materiaHandler:    matHdlr,
		personalHandler:   persHdlr,
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

	personal := app.Group("/personal")
	personal.Post("/", s.personalHandler.NuevoPersonal)
	personal.Get("/todos", s.personalHandler.TodosPersonal)
	personal.Get("/:id", s.personalHandler.UnPersonal)
	personal.Put("/:id", s.personalHandler.EditarPersonal)
	personal.Delete("/:id", s.personalHandler.BorrarPersonal)

	periodo := app.Group("/periodo")
	periodo.Post("/", s.periodoHandler.NuevoPeriodo)
	periodo.Get("/todos", s.periodoHandler.TodosPeriodos)
	periodo.Get("/:id", s.periodoHandler.UnPeriodo)
	periodo.Put("/:id", s.periodoHandler.EditarPeriodo)
	periodo.Delete("/:id", s.periodoHandler.BorrarPeriodo)

	grupo := app.Group("/grupo")
	grupo.Post("/", s.grupoHandler.NuevoGrupo)
	grupo.Get("/todos", s.grupoHandler.TodosGrupos)
	grupo.Get("/:id", s.grupoHandler.UnGrupo)
	grupo.Put("/:id", s.grupoHandler.EditarGrupo)
	grupo.Delete("/:id", s.grupoHandler.BorrarGrupo)

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

	materia := app.Group("/materia")
	materia.Post("/", s.materiaHandler.NuevaMateria)
	materia.Get("/todas", s.materiaHandler.TodasMaterias)
	materia.Get("/:id", s.materiaHandler.UnaMateria)
	materia.Put("/:id", s.materiaHandler.EditarMateria)
	materia.Delete("/:id", s.materiaHandler.BorrarMateria)

	app.Post("/generateAl", middlewares.GenerarLogAlumno)
	app.Post("/generateMa", middlewares.GenerarLogMaestro)
	app.Get("/validate", middlewares.Verificar)
	app.Post("/refreshToken", middlewares.Refrescar)
	app.Get("/", middlewares.Authorizar, func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-Almacén 👋!")
	})
	return app
}
