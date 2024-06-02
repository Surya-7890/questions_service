package models

type Company struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"unique;not null"`
	Questions []Question `gorm:"foreignKey:ID"`
}
