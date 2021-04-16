package handlers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/handlers/alert"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/handlers/user"
	"net/http"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router { //nolint:funlen
	var (
		router    = mux.NewRouter()
		apiRouter = router.PathPrefix(config.Config.URLPrefix).Subrouter()
		v1Router  = apiRouter.PathPrefix("/v1").Subrouter()

		publicChain = alice.New()
	)

	// alert
	v1Router.Handle("/price/alert", publicChain.ThenFunc(alert.GetAlertsByUser)).Methods(http.MethodGet)
	v1Router.Handle("/price/alert/{alert_id}", publicChain.ThenFunc(alert.DeleteAlert)).Methods(http.MethodDelete)
	v1Router.Handle("/price/alert/{alert_id}", publicChain.ThenFunc(alert.EditAlert)).Methods(http.MethodPatch)
	v1Router.Handle("/price/alert", publicChain.ThenFunc(alert.AddAlert)).Methods(http.MethodPost)

	// user
	v1Router.Handle("/user", publicChain.ThenFunc(user.CreateUser)).Methods(http.MethodPost)
	v1Router.Handle("/user/{user_id}", publicChain.ThenFunc(user.GetUser)).Methods(http.MethodGet)

	return router
}
