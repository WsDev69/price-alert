package handlers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/handlers/mail"
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

	v1Router.Handle("/mail/send", publicChain.ThenFunc(mail.SendMail)).Methods(http.MethodPost)

	return router
}
