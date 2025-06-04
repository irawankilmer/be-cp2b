package domain

type Account struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"uniqueIndex;not null"`
	Descriptions string  `gorm:"type:text"`
	Balance      Balance `gorm:"constraint:OnDelete:CASCADE"`
	TimeStamps
}
