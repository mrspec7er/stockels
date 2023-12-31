// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package object

type About struct {
	Title       *string      `json:"title,omitempty"`
	Info        []*Info      `json:"info,omitempty"`
	Description *Description `json:"description,omitempty"`
}

type Analytic struct {
	Financials     []*Financials   `json:"financials,omitempty"`
	KnowledgeGraph *KnowledgeGraph `json:"knowledge_graph,omitempty"`
	Graph          []*Graph        `json:"graph,omitempty"`
	Summary        *Summary        `json:"summary,omitempty"`
}

type Article struct {
	Source      *Source `json:"source"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	PublishedAt string  `json:"publishedAt"`
}

type Description struct {
	Snippet  *string `json:"snippet,omitempty"`
	Link     *string `json:"link,omitempty"`
	LinkText *string `json:"link_text,omitempty"`
}

type Financials struct {
	Title   *string    `json:"title,omitempty"`
	Results []*Results `json:"results,omitempty"`
}

type GenerateReportResponse struct {
	ReportURL string `json:"reportUrl"`
}

type GetStockData struct {
	StockSymbol     string `json:"stockSymbol"`
	SupportPrice    int    `json:"supportPrice"`
	ResistancePrice int    `json:"resistancePrice"`
}

type Graph struct {
	Price    *float64 `json:"price,omitempty"`
	Currency *string  `json:"currency,omitempty"`
	Date     *string  `json:"date,omitempty"`
	Volume   *float64 `json:"volume,omitempty"`
}

type Info struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
	Link  *string `json:"link,omitempty"`
}

type KeyStats struct {
	Stats []*Stats `json:"stats,omitempty"`
	Tags  []*Tags  `json:"tags,omitempty"`
}

type KnowledgeGraph struct {
	About    []*About  `json:"about,omitempty"`
	KeyStats *KeyStats `json:"key_stats,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type PriceMovement struct {
	Percentage *float64 `json:"percentage,omitempty"`
	Value      *float64 `json:"value,omitempty"`
	Movement   *string  `json:"movement,omitempty"`
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

type Results struct {
	Date  *string  `json:"date,omitempty"`
	Table []*Table `json:"table,omitempty"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Stats struct {
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Value       *string `json:"value,omitempty"`
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

type Summary struct {
	Title          *string        `json:"title,omitempty"`
	Stock          *string        `json:"stock,omitempty"`
	Exchange       *string        `json:"exchange,omitempty"`
	Price          *string        `json:"price,omitempty"`
	ExtractedPrice *float64       `json:"extracted_price,omitempty"`
	Currency       *string        `json:"currency,omitempty"`
	PriceMovement  *PriceMovement `json:"price_movement,omitempty"`
}

type Table struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Value       *string `json:"value,omitempty"`
	Change      *string `json:"change,omitempty"`
}

type Tags struct {
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
}

type TechnicalAnalytic struct {
	AverageSupportPrice    float64            `json:"averageSupportPrice"`
	AverageResistancePrice float64            `json:"averageResistancePrice"`
	Quarters               []*QuarterAnalytic `json:"quarters"`
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
