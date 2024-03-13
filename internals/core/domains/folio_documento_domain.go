package domains

import (
	"gorm.io/gorm"
)

func (FolioDocumento) TableName() string {
	return "folio_documento"
}

type FolioDocumento struct {
	gorm.Model
	Tipo  string `gorm:"not null"`
	Folio Folio
}
