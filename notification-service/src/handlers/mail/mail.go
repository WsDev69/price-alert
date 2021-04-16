package mail

import (
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/handlers/common"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/models"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/services"
	"net/http"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req = models.EmailNotification{}
	)

	httpErr := common.UnmarshalRequestBody(r, &req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	httpErr = services.Get().GetMail().Send(ctx, req.From, req.Text, []string{req.To})
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, struct{}{})
}
