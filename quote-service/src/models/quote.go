package models

type WSJSONParams struct {
	Action string   `json:"action"`
	Subs   []string `json:"subs"`
}

type WSResponse struct {
	Type       string   `json:"TYPE"`
	Market     string  `json:"MARKET"`
	FromSymbol string  `json:"FROMSYMBOL"`
	ToSymbol   string  `json:"TOSYMBOL"`
	Price      float64 `json:"PRICE"`
}

type AddNewCurrencyRequest struct {
	FromSymbol string  `json:"from_symbol"`
	ToSymbol   string  `json:"to_symbol"`
}
