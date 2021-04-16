package quote

type AddNewCurrencyRequest struct {
	FromSymbol string `json:"from_symbol"`
	ToSymbol   string `json:"to_symbol"`
}
