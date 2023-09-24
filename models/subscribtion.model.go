package models

type Subscribtion struct {
	StockSymbol string `json:"stockId" gorm:"primaryKey;not null"`
	UserID int `json:"userId" gorm:"primaryKey;autoIncrement:false"`
	SupportPrice int `json:"supportPrice"`
	ResistancePrice int `json:"resistancePrice"`
	Stock Stock `json:"stock" gorm:"foreignKey:StockSymbol;references:Symbol"`
	User User `json:"user"`
}