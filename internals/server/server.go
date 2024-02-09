package server

import (
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	alumnoHandler ports.IAlumnoHandler
}

func NewServer(alumnoHandler ports.IAlumnoHandler,

// eHandler ports.IEntradaHandler, mHandler ports.IMunicipioHandler, lHandler ports.ILocalidadHandler, rHandler ports.IRegionHandler,
// iHandler ports.IInmuebleHandler, dHandler ports.IProgramaHandler, sHandler ports.ISalidaHandler, cHandler ports.ICategoriaHandler,
// pHandler ports.IProductoHandler, preHandler ports.IPresentacionHandler, epHandler ports.IEntradaProductoHandler, invHandler ports.IInventarioHandler,
// spHandler ports.ISalidaProductoHandler, encHandler ports.IEncargadoHandler, perHandler ports.IPersonaHandler, usHandler ports.IUsuarioHandler,
// externos ports.IExternoHandler,
) *Server {
	return &Server{
		alumnoHandler: alumnoHandler,
		// entradaHandler:         eHandler,
		// municipioHandler:       mHandler,
		// localidadHandler:       lHandler,
		// regionHandler:          rHandler,
		// inmuebleHandler:        iHandler,
		// programaHandler:        dHandler,
		// salidaHandler:          sHandler,
		// categoriaHandler:       cHandler,
		// productoHandler:        pHandler,
		// presentacionHandler:    preHandler,
		// entradaProductoHandler: epHandler,
		// inventarioHandler:      invHandler,
		// salidaProductoHandler:  spHandler,
		// encargadoHandler:       encHandler,
		// personaHandler:         perHandler,
		// usuarioHandler:         usHandler,
		// externos:               externos,
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-Almacén 👋!")
	})

	// app.Post("/generate", middlewares.Generar)
	// app.Get("/validate", middlewares.Verificar)
	// app.Post("/refreshToken", middlewares.Refrescar)
	/*app.Get("/", handlers.Authorizar, func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-Almacén 👋!")
	})


	app.Post("/generate", handlers.Generar)
	app.Get("/validate", handlers.Verificar)
	app.Get("/refreshToken", handlers.Refrescar)*/

	// localidades := app.Group("/localidades")
	// localidades.Get("/", s.externos.ObtenerTodasLocalidadesExterno)

	// municipios := app.Group("/municipios")
	// municipios.Get("/", s.externos.ObtenerTodasMunicipiosExterno)

	// catAreas := app.Group("/catAreas")
	// catAreas.Get("/", s.externos.ObtenerTodasCatAreas)

	// usuarios := app.Group("/usuarios")
	// usuarios.Get("/", middlewares.Usuarios)

	// inventario := app.Group("/inventario")
	// inventario.Post("/", s.inventarioHandler.NuevoInventario)
	// inventario.Get("/todo", s.inventarioHandler.TodosInventarios)
	// inventario.Get("/todo/inmueble/:id", s.inventarioHandler.TodosInventariosPorInmueble)
	// inventario.Get("/todo/producto/:id", s.inventarioHandler.TodosInventariosPorProducto)

	// entrada := app.Group("/entrada")
	// //entrada.Use(handlers.Authorizar)
	// entrada.Post("/", s.entradaHandler.NuevaEntrada)
	// entrada.Post("/:id/:cantidad", s.entradaHandler.NuevaEntradaProducto) //No
	// entrada.Get("/todas", s.entradaHandler.TodasEntradas)
	// entrada.Get("/:id", s.entradaHandler.UnaEntrada)
	// entrada.Put("/:id", s.entradaHandler.EditaEntrada)
	// entrada.Delete("/:id", s.entradaHandler.BorrarEntrada)
	// entrada.Post("/:id", s.entradaProductoHandler.NuevaEntradaProducto)

	// municipio := app.Group("/municipio")
	// // municipio.Use(handlers.Authorizar)
	// municipio.Post("/", s.municipioHandler.NuevoMunicipio)
	// // municipio.Put("/:id", s.municipionHandler.EditarMunicipio)
	// municipio.Get("/", s.municipioHandler.BuscarPor)
	// municipio.Get("/todos", s.municipioHandler.TodosMunicipios)

	// localidad := app.Group("/localidad")
	// // localidad.Use(handlers.Authorizar)
	// localidad.Post("/", s.localidadHandler.NuevaLocalidad)
	// // localidad.Put("/:id", s.localidadHandler.EditarLocalidad)
	// localidad.Get("/", s.localidadHandler.BuscarPor)
	// localidad.Get("/:municipioId", s.localidadHandler.BuscarPorMunicipio)
	// // localidad.Get("/:municipioId", s.localidadHandler.BuscarPor)
	// localidad.Get("/todas", s.localidadHandler.TodasLocalidades)

	// region := app.Group("/region")
	// // localidad.Use(handlers.Authorizar)
	// region.Post("/", s.regionHandler.NuevaRegion)
	// // localidad.Put("/:id", s.localidadHandler.EditarLocalidad)
	// region.Get("/", s.regionHandler.BuscarPorNombre)
	// region.Get("/todas", s.regionHandler.TodasRegiones)

	// inmueble := app.Group("/inmueble")
	// inmueble.Post("/", s.inmuebleHandler.NuevoInmueble)
	// inmueble.Get("/todos", s.inmuebleHandler.TodosInmuebles)
	// inmueble.Get("/:id", s.inmuebleHandler.UnInmueble)
	// inmueble.Put("/:id", s.inmuebleHandler.EditaInmueble)
	// inmueble.Delete("/:id", s.inmuebleHandler.BorrarInmueble)

	// programa := app.Group("/programa")
	// programa.Post("/", s.programaHandler.NuevoPrograma)
	// programa.Get("/todos", s.programaHandler.TodosProgramas)
	// programa.Get("/:id", s.programaHandler.UnPrograma)
	// programa.Put("/:id", s.programaHandler.EditaPrograma)
	// programa.Delete("/:id", s.programaHandler.BorrarPrograma)

	// salida := app.Group("/salida")
	// salida.Post("/", s.salidaHandler.NuevaSalida)
	// salida.Get("/todas", s.salidaHandler.TodasSalidas)
	// salida.Get("/:id", s.salidaHandler.UnaSalida)
	// salida.Put("/:id", s.salidaHandler.EditaSalida)
	// salida.Delete("/:id", s.salidaHandler.BorrarSalida)
	// salida.Post("/:id", s.salidaProductoHandler.NuevaSalidaProducto)

	// categoria := app.Group("/categoria")
	// categoria.Post("/", s.categoriaHandler.NuevaCategoria)
	// categoria.Get("/todas", s.categoriaHandler.BuscarTodo)
	// categoria.Get("/:id", s.categoriaHandler.UnaCategoria)
	// categoria.Put("/:id", s.categoriaHandler.ActualizarCategoria)
	// categoria.Delete("/:id", s.categoriaHandler.EliminarCategoria)

	// producto := app.Group("/producto")
	// producto.Post("/", s.productoHandler.NuevoProducto)
	// producto.Get("/todos", s.productoHandler.TodosProductos)
	// producto.Get("/:id", s.productoHandler.UnProducto)
	// producto.Put("/:id", s.productoHandler.EditaProducto)
	// producto.Delete("/:id", s.productoHandler.BorrarProducto)

	// presentacion := app.Group("/presentacion")
	// presentacion.Post("/", s.presentacionHandler.NuevaPresentacion)
	// presentacion.Get("/todas", s.presentacionHandler.TodasPresentaciones)
	// presentacion.Get("/:id", s.presentacionHandler.UnaPresentacion)
	// presentacion.Put("/:id", s.presentacionHandler.EditaPresentacion)
	// presentacion.Delete("/:id", s.presentacionHandler.BorrarPresentacion)

	return app
}
