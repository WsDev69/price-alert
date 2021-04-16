package alert

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/handlers/common"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/alert"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services"
	"net/http"
)

func AddAlert(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		req       = alert.Alert{}
		userIDstr = r.URL.Query().Get("user_id")
	)

	httpErr := common.UnmarshalRequestBody(r, &req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		logrus.Error("can't parse userID",
			"userID", userIDstr, "err", err)
		common.SendHTTPError(w, httperrors.NewBadRequestError(err))
		return
	}

	resp, httpErr := services.Get().Alert().Create(ctx, userID, req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}

func GetAlertsByUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		userIDstr = mux.Vars(r)["user_id"]
	)

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		logrus.Error("can't parse userID",
			"userID", userIDstr, "err", err)
		common.SendHTTPError(w, httperrors.NewBadRequestError(err))
		return
	}

	limit, offset, httpErr := common.GetLimitAndOffset(r.URL.Query())
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	resp, httpErr := services.Get().Alert().GetAllByUserID(ctx, limit, offset, userID)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        = r.Context()
		alertIDstr = mux.Vars(r)["alert_id"]
	)

	alertID, err := uuid.Parse(alertIDstr)
	if err != nil {
		logrus.Error("can't parse alert ID",
			"alertID", alertID, "err", err)
		common.SendHTTPError(w, httperrors.NewBadRequestError(err))
		return
	}

	resp, httpErr := services.Get().Alert().Delete(ctx, alertID)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}

func EditAlert(w http.ResponseWriter, r *http.Request) {
	var (
		ctx        = r.Context()
		alertIDstr = mux.Vars(r)["alert_id"]
		req        = alert.Alert{}
	)

	httpErr := common.UnmarshalRequestBody(r, &req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	alertID, err := uuid.Parse(alertIDstr)
	if err != nil {
		logrus.Error("can't parse alert ID",
			"alertID", alertID, "err", err)
		common.SendHTTPError(w, httperrors.NewBadRequestError(err))
		return
	}

	resp, httpErr := services.Get().Alert().Update(ctx, req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}
