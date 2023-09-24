package models

type Stock struct {
	Symbol string `json:"symbol" gorm:"primaryKey;not null"`
	Name string `json:"name"`
	Sector string `json:"sector"`
	Website string `json:"website"`
	Logo string `json:"logo"`
	Description string `json:"description"`
	OpenPrice string `json:"openPrice"`
	ClosePrice string `json:"closePrice"`
	HighestPrice string `json:"highestPrice"`
	LowestPrice string `json:"lowestPrice"`
	Volume string `json:"volume"`
	LastUpdate string `json:"lastUpdate"`
	Subscribtion []Subscribtion `json:"subscribtions"`
}