// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package object

type Article struct {
	Source      *Source `json:"source"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	PublishedAt string  `json:"publishedAt"`
}

type GenerateReportResponse struct {
	ReportURL string `json:"reportUrl"`
}

type GetStockData struct {
	StockSymbol     string `json:"stockSymbol"`
	SupportPrice    int    `json:"supportPrice"`
	ResistancePrice int    `json:"resistancePrice"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type QuarterAnalytic struct {
	Quarter          string  `json:"quarter"`
	SupportPrice     float64 `json:"supportPrice"`
	SupportDate      string  `json:"supportDate"`
	SupportVolume    int     `json:"supportVolume"`
	ResistancePrice  float64 `json:"resistancePrice"`
	ResistanceDate   string  `json:"resistanceDate"`
	ResistanceVolume int     `json:"resistanceVolume"`
}

type Register struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StockAnalytic struct {
	AverageSupportPrice    float64            `json:"averageSupportPrice"`
	AverageResistancePrice float64            `json:"averageResistancePrice"`
	Quarters               []*QuarterAnalytic `json:"quarters"`
}

type StockData struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Sector               string  `json:"sector"`
	Logo                 string  `json:"logo"`
	Website              string  `json:"website"`
	OpenPrice            string  `json:"openPrice"`
	ClosePrice           string  `json:"closePrice"`
	HighestPrice         string  `json:"highestPrice"`
	LowestPrice          string  `json:"lowestPrice"`
	Volume               string  `json:"volume"`
	LastUpdate           string  `json:"lastUpdate"`
	SupportPercentage    float64 `json:"supportPercentage"`
	ResistancePercentage float64 `json:"resistancePercentage"`
}

type StockDetail struct {
	Info  *StockData          `json:"info"`
	Price []*StockDetailPrice `json:"price"`
}

type StockDetailPrice struct {
	Date   string `json:"date"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	Volume int    `json:"volume"`
}

type Subscribtion struct {
	StockSymbol     string `json:"stockSymbol"`
	UserID          int    `json:"userId"`
	SupportPrice    int    `json:"supportPrice"`
	ResistancePrice int    `json:"resistancePrice"`
}

type User struct {
	ID         int     `json:"id"`
	FullName   string  `json:"fullName"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	IsVerified bool    `json:"isVerified"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	DeletedAt  *string `json:"deletedAt,omitempty"`
}
