package domains

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	PersonalID uint   `gorm:"not null"`
	PeriodoID  uint   `gorm:"not null"`
	Extension  string `gorm:"not null"`
}
