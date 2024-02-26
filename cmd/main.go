package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/adapters/handlers"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/adapters/repositories"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/middlewares"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/services"
	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/server"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	type Configuration struct {
		Host       string `env:"HOST"`
		Port       string `env:"PORT"`
		DbUser     string `env:"DBUSER"`
		DbPassword string `env:"DBPASSWORD"`
		DbName     string `env:"DBNAME"`
		DbHost     string `env:"DBHOST"`
		DbPort     string `env:"DBPORT"`
		DbSslmode  string `env:"DBSSLMODE"`
		DbTimezone string `env:"DBTIMEZONE"`
		//
		JWTDbUser        string `env:"JWTDBUSER"`
		JWTDbPassword    string `env:"JWTDBPASSWORD"`
		JWTDbName        string `env:"JWTDBNAME"`
		JWTDbHost        string `env:"JWTDBHOST"`
		JWTDbPort        string `env:"JWTDBPORT"`
		JWTDbSslmode     string `env:"JWTDBSSLMODE"`
		JWTJWTDbTimezone string `env:"JWTDBTIMEZONE"`
	}

	errno := godotenv.Load()
	if errno != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Configuration
	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		log.Println("err", err)
		os.Exit(1)
	}

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	// jwtDbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.JWTDbUser, cfg.JWTDbPassword, cfg.JWTDbHost, cfg.JWTDbPort, cfg.JWTDbName)

	db, err := gorm.Open(mysql.Open(dbConnString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.SetupJoinTable(&domains.Alumno{}, "Grupos", &domains.AlumnoGrupo{})
	db.SetupJoinTable(&domains.Personal{}, "Perfiles", &domains.UsuariosPerfiles{})
	db.AutoMigrate(domains.Materia{}, domains.Perfil{}, domains.Estudiante{}, domains.Alumno{}, domains.Grupo{}, domains.GrupoMateria{}, domains.Calificacion{}, domains.Personal{})

	middlewares.Init(db)

	alumnoRepository := repositories.NewAlumnoRepository(db)
	alumnoService := services.NewAlumnoService(alumnoRepository)
	alumnoHandler := handlers.NewAlumnoHandler(alumnoService)

	personalRepository := repositories.NewPersonalRepository(db)
	personalService := services.NewPersonalService(personalRepository)
	personalHandler := handlers.NewPersonalHandler(personalService)

	periodoRepository := repositories.NewPeriodoRepository(db)
	periodoService := services.NewPeriodoService(periodoRepository)
	periodoHandler := handlers.NewPeriodoHandler(periodoService)

	materiaRepository := repositories.NewMateriaRepository(db)
	materiaService := services.NewMateriaService(materiaRepository)
	materiaHandler := handlers.NewMateriaHandler(materiaService)

	grupoRepository := repositories.NewGrupoRepository(db)
	grupoService := services.NewGrupoService(grupoRepository)
	grupoHandler := handlers.NewGrupoHandler(grupoService)

	estudianteRepository := repositories.NewEstudianteRepository(db)
	estudianteService := services.NewEstudianteService(estudianteRepository)
	estudianteHandler := handlers.NewEstudianteHandler(estudianteService)

	carreraRepository := repositories.NewCarreraRepository(db)
	carreraService := services.NewCarreraService(carreraRepository)
	carreraHandler := handlers.NewCarreraHandler(carreraService)

	server := server.NewServer(
		alumnoHandler,
		carreraHandler,
		estudianteHandler,
		grupoHandler,
		periodoHandler,
		materiaHandler,
		personalHandler,
	)

	server.Initizalize().Listen(":8000")
}
