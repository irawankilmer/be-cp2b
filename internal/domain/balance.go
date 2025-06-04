package domain

type Balance struct {
	ID        uint     `gorm:"primaryKey"`
	AccountID uint     `gorm:"not null;unique"`
	Balance   float64  `gorm:"type:decimal(15,2)"`
	Account   *Account `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE"`
	TimeStamps
}
