package models

type TagRelationship struct {
	Tag1ID            uint `gorm:"primaryKey"`
	Tag2ID            uint `gorm:"primaryKey"`
	Score             float64
	CoOccurrenceCount int
}
