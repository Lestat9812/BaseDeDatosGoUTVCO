package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PersonalRepository struct {
	db *gorm.DB
}

func NewPersonalRepository(storage *gorm.DB) *PersonalRepository {
	return &PersonalRepository{
		db: storage,
	}
}

func (r *PersonalRepository) SavePersonal(personal *domains.Personal) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(personal.Password), 10)
	if err != nil {
		return errors.New(err.Error())
	} else {
		personal.Password = string(passwordHash)
	}
	result := r.db.Create(&personal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PersonalRepository) AllPersonal() ([]*domains.Personal, error) {
	var results []*domains.Personal

	if res := r.db.Preload(clause.Associations).
		Find(&results); res.Error != nil {
		return nil, res.Error
	}

	if len(results) == 0 {
		return nil, errors.New("data not found")
	} else {
		return results, nil
	}
}

func (r *PersonalRepository) FindPersonalById(id int) (*domains.Personal, error) {
	var personal domains.Personal
	result := r.db.
		First(&personal, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Personal{}, errors.New("data not found")
	}
	return &personal, nil
}
func (r *PersonalRepository) UpdatePersonal(personal *domains.Personal) error {
	result := r.db.Model(&domains.Personal{}).Where("id = ?", personal.ID).Updates(personal)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("personal no actualizado")
	}
	return nil
}

func (r *PersonalRepository) DeletePersonal(id int) error {
	var deletedPersonal domains.Personal
	result := r.db.Where("id = ?", id).Delete(&deletedPersonal)
	if result.RowsAffected == 0 {
		return errors.New("personal no borrado")
	}
	return nil
}
