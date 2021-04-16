package quotes

import (
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/handlers/common"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/models"
	httperrors "github.com/wsdev69/price-alert/quote-service/v0.0.1/src/models/http-errors"

	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/services"
	"net/http"
)

func UpdateCurrencies(w http.ResponseWriter, r *http.Request) {
	var req models.AddNewCurrencyRequest

	if httpErr := common.UnmarshalRequestBody(r, &req); httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	if err := services.Get().GetQuote().AddCurrencyPair(req.ToSymbol, req.FromSymbol);
		err != nil {
		common.SendHTTPError(w, httperrors.NewInternalServerError(err))
		return
	}

	common.SendResponse(w, http.StatusOK, struct{}{})
}
