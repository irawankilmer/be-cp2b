package domain

type Category struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"uniqueIndex" json:"name"`
	Type         string `gorm:"type:enum('pemasukan','pengeluaran','pindah');default:'pemasukan'" json:"type"`
	Descriptions string `gorm:"type:text" json:"descriptions"`
	TimeStamps
}
