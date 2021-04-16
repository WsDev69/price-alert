package ticker

type Ticker struct {
	Type       string  `json:"TYPE"`
	Market     string  `json:"MARKET"`
	FromSymbol string  `json:"FROMSYMBOL"`
	ToSymbol   string  `json:"TOSYMBOL"`
	Price      float64 `json:"PRICE"`
}
