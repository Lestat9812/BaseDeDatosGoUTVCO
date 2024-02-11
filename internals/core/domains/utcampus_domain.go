package domains

import (
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Utcampus) TableName() string {
	return "ut_campus"
}

type Utcampus struct {
	gorm.Model
	Nombre    string `gorm:"column:nombre"`
	Siglas    string `gorm:"column:siglas"`
	Ubicacion string `gorm:"column:ubicacion"`
}
