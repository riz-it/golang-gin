package model

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:varchar(50)" json:"description"`
}
