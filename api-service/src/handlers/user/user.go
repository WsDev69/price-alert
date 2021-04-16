package user

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/handlers/common"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/user"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req = user.User{}
	)

	httpErr := common.UnmarshalRequestBody(r, &req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	resp, httpErr := services.Get().User().Create(ctx, req)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = r.Context()
		userIDstr = mux.Vars(r)["user_id"]
	)

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		logrus.Error("can't parse userID ID",
			"userID", userIDstr, "err", err)
		common.SendHTTPError(w, httperrors.NewBadRequestError(err))
		return
	}

	resp, httpErr := services.Get().User().Get(ctx, userID)
	if httpErr != nil {
		common.SendHTTPError(w, httpErr)
		return
	}

	common.SendResponse(w, http.StatusOK, resp)
}
