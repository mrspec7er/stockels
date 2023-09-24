package models

type Subscribtion struct {
	StockSymbol string `json:"stockSymbol" gorm:"primaryKey;not null"`
	UserID int `json:"userId" gorm:"primaryKey;autoIncrement:false"`
	SupportPrice int `json:"supportPrice"`
	ResistancePrice int `json:"resistancePrice"`
	Stock *Stock `json:"stock,omitempty" gorm:"foreignKey:StockSymbol;references:Symbol"`
	User *User `json:"user,omitempty"`
}