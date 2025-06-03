package domain

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"size:35;not null"`
	Email           string `gorm:"uniqueIndex;not null"`
	EmailVerifiedAt string
	Password        string
	TokenVersion    string
	Transactions    []Transaction
	TimeStamps
}
