package repositories

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

var folio int = rand.Intn(80) + 1

type FakerRepository struct {
	db *gorm.DB
}

func NewFakerRepository(storage *gorm.DB) *FakerRepository {
	return &FakerRepository{
		db: storage,
	}
}

func (r *FakerRepository) SaveFaker() error {
	var preficha domains.Preficha
	var aspirante domains.Aspirante
	var estudiante domains.Estudiante
	var carrera domains.Carrera
	var password = "123456"
	for i := 0; i < 10; i++ {
		var car int = rand.Intn(9) + 1
		var name = faker.FirstName()
		var apP = faker.LastName()
		var apM = faker.LastName()
		result := r.db.Create(&domains.Preficha{
			Nombre:    name,
			ApellidoP: apP,
			ApellidoM: &apM,
			CarFicha:  car,
		})
		if result.Error != nil {
			return result.Error
		}
		r.db.Model(&domains.Preficha{}).Select("id").Last(&preficha)
		folio = folio + i
		fmt.Println(folio)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			return err
		} else {
			result2 := r.db.Create(&domains.Aspirante{
				Fecha:      time.Now().String(),
				Folio:      folio,
				CarFinal:   &car,
				PeriodoID:  1,
				Password:   string(passwordHash),
				PrefichaID: preficha.ID + uint(i),
			})
			if result2.Error != nil {
				return result2.Error
			}
		}
		r.db.Model(&domains.Aspirante{}).Select("id, car_final, folio").Last(&aspirante)

		fmt.Println(i)
		fmt.Println(aspirante.Folio)
		result3 := r.db.Create(&domains.Estudiante{
			Nombre:      name,
			ApellidoP:   apP,
			ApellidoM:   &apM,
			PrefichaID:  preficha.ID + uint(i),
			AspiranteID: aspirante.ID + uint(i),
			CarFicha:    *aspirante.CarFinal,
		})
		if result3.Error != nil {
			return result3.Error
		}
		r.db.Model(&domains.Estudiante{}).Select("id, car_ficha").Last(&estudiante)
		r.db.Model(&domains.Carrera{}).Select("nombre").Where("id=?", car).First(&carrera)
		mat := "UT"
		if strings.Contains(carrera.Nombre, "Tecn") {
			mat = mat + "TI"
		}
		if strings.Contains(carrera.Nombre, "Mec") {
			mat = mat + "ME"
		}
		if strings.Contains(carrera.Nombre, "Gast") {
			mat = mat + "GA"
		}
		if strings.Contains(carrera.Nombre, "Gest") {
			mat = mat + "TI"
		}
		mat = mat + strings.Replace(strconv.Itoa(time.Now().Year()), "20", "", 1) + "20" + strconv.Itoa(folio)
		fmt.Println(mat)
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			return err
		} else {
			result4 := r.db.Create(&domains.Alumno{
				Matricula: mat,
				// Folio:      11 + i,
				// PeriodoID:  1,
				Password:     string(passwordHash),
				CarreraID:    uint(car),
				EstudianteID: estudiante.ID + uint(i),
			})
			if result4.Error != nil {
				return result4.Error
			}
		}

	}
	// result := r.db.Create(&domains.Alumno{})
	return nil
}
