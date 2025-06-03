package domain

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey"`
	Date            time.Time `gorm:"type:date;not null"`
	Type            string    `gorm:"type:enum('pemasukan', 'pengeluaran', 'pindah');not null;default:'pengeluaran'"`
	AccountID       uint
	Account         Account `gorm:"foreignKey:AccountID"`
	CategoryID      uint
	Category        Category `gorm:"foreignKey:CategoryID"`
	TargetAccountID *uint
	TargetAccount   Account `gorm:"foreignKey:AccountID"`
	Amount          float64 `gorm:"type:decimal(15,2);not null"`
	Descriptions    string  `gorm:"type:varchar(255)"`
	UserID          uint
	User            User `gorm:"foreignKey:UserID"`
	TimeStamps
}
