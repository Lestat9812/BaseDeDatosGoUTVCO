package repositories

import (
	"errors"

	"github.com/Lestat9812/BaseDeDatosGoUTVCO/internals/core/domains"
	"gorm.io/gorm"
)

type GrupoRepository struct {
	db *gorm.DB
}

func NewGrupoRepository(storage *gorm.DB) *GrupoRepository {
	return &GrupoRepository{
		db: storage,
	}
}

func (r *GrupoRepository) SaveGrupo(grupo *domains.Grupo) error {
	result := r.db.Create(&grupo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GrupoRepository) AllGrupos() ([]*domains.Grupo, error) {
	var results []*domains.Grupo

	if res := r.db.
		Find(&results); res.Error != nil {
		return nil, res.Error
	}

	if len(results) == 0 {
		return nil, errors.New("data not found")
	} else {
		return results, nil
	}
}

func (r *GrupoRepository) FindGrupoById(id int) (*domains.Grupo, error) {
	var grupo domains.Grupo
	result := r.db.
		First(&grupo, "id = ?", id)

	if result.RowsAffected == 0 {
		return &domains.Grupo{}, errors.New("data not found")
	}
	return &grupo, nil
}
func (r *GrupoRepository) UpdateGrupo(grupo *domains.Grupo) error {
	result := r.db.Model(&domains.Grupo{}).Where("id = ?", grupo.ID).Updates(grupo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("grupo no actualizado")
	}
	return nil
}

func (r *GrupoRepository) DeleteGrupo(id int) error {
	var deletedGrupo domains.Grupo
	result := r.db.Where("id = ?", id).Delete(&deletedGrupo)
	if result.RowsAffected == 0 {
		return errors.New("grupo no borrado")
	}
	return nil
}
