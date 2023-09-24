package models

type Subscribtion struct {
	StockID int `json:"stockId" gorm:"primaryKey;autoIncrement:false"`
	UserID int `json:"userId" gorm:"primaryKey;autoIncrement:false"`
	SupportPrice int `json:"supportPrice"`
	ResistancePrice int `json:"resistancePrice"`
	Stock Stock `json:"stock"`
	User User `json:"user"`
}