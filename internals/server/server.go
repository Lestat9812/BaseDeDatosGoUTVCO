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
	utcampusHandler   ports.IUtcampusHandler
	materiaHandler    ports.IMateriaHandler
}

func NewServer(alHdlr ports.IAlumnoHandler, carHdlr ports.ICarreraHandler,
	estHdlr ports.IEstudianteHandler, grHdlr ports.IGrupoHandler,
	peHdlr ports.IPeriodoHandler, persHdlr ports.IPersonalHandler,
	utcHdlr ports.IUtcampusHandler, matHdlr ports.IMateriaHandler) *Server {
	return &Server{
		alumnoHandler:     alHdlr,
		carreraHandler:    carHdlr,
		estudianteHandler: estHdlr,
		grupoHandler:      grHdlr,
		periodoHandler:    peHdlr,
		personalHandler:   persHdlr,
		utcampusHandler:   utcHdlr,
		materiaHandler:    matHdlr,
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

	// app.Use(middlewares.Autorizar)  no funciona porque bloquea el login también

	alumno := app.Group("/alumno", middlewares.Autorizar)
	alumno.Post("/", s.alumnoHandler.NuevoAlumno)
	alumno.Get("/todos", s.alumnoHandler.TodosAlumnos)
	alumno.Get("/:id", s.alumnoHandler.UnAlumno)
	alumno.Get("/grupo/:id", s.alumnoHandler.UnAlumnoGrupo)
	alumno.Get("/calificaciones/:id", s.alumnoHandler.UnAlumnoCalificaciones)
	alumno.Put("/:id", s.alumnoHandler.EditarAlumno)
	alumno.Delete("/:id", s.alumnoHandler.BorrarAlumno)

	personal := app.Group("/personal", middlewares.Autorizar)
	personal.Post("/", s.personalHandler.NuevoPersonal)
	personal.Get("/todos", s.personalHandler.TodosPersonal)
	personal.Get("/:id", s.personalHandler.UnPersonal)
	personal.Put("/:id", s.personalHandler.EditarPersonal)
	personal.Delete("/:id", s.personalHandler.BorrarPersonal)

	utcampus := app.Group("/utcampus", middlewares.Autorizar)
	utcampus.Post("/", s.utcampusHandler.NuevoUtcampus)
	utcampus.Get("/todos", s.utcampusHandler.TodosUtcampus)
	utcampus.Get("/:id", s.utcampusHandler.UnUtcampus)
	utcampus.Put("/:id", s.utcampusHandler.EditarUtcampus)
	utcampus.Delete("/:id", s.utcampusHandler.BorrarUtcampus)

	periodo := app.Group("/periodo", middlewares.Autorizar)
	periodo.Post("/", s.periodoHandler.NuevoPeriodo)
	periodo.Get("/todos", s.periodoHandler.TodosPeriodos)
	periodo.Get("/:id", s.periodoHandler.UnPeriodo)
	periodo.Put("/:id", s.periodoHandler.EditarPeriodo)
	periodo.Delete("/:id", s.periodoHandler.BorrarPeriodo)

	grupo := app.Group("/grupo", middlewares.Autorizar)
	grupo.Post("/", s.grupoHandler.NuevoGrupo)
	grupo.Get("/todos", s.grupoHandler.TodosGrupos)
	grupo.Get("/:id", s.grupoHandler.UnGrupo)
	grupo.Put("/:id", s.grupoHandler.EditarGrupo)
	grupo.Delete("/:id", s.grupoHandler.BorrarGrupo)

	estudiante := app.Group("/estudiante", middlewares.Autorizar)
	estudiante.Post("/", s.estudianteHandler.NuevoEstudiante)
	estudiante.Get("/todos", s.estudianteHandler.TodosEstudiantes)
	estudiante.Get("/:id", s.estudianteHandler.UnEstudiante)
	estudiante.Put("/:id", s.estudianteHandler.EditarEstudiante)
	estudiante.Delete("/:id", s.estudianteHandler.BorrarEstudiante)

	carrera := app.Group("/carrera", middlewares.Autorizar)
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
	// app.Get("/", middlewares.Authorizar, s.alumnoHandler.TodosAlumnos)
	// func(c *fiber.Ctx) error {
	// 	return c.SendString("Bienvenido, Servicios Escolares UTVCO 👋!")
	// })
	return app
}
