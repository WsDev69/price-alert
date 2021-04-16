package mail

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/smtp"

	httperrors "github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/models/http-errors"

)

func (srv smtpService) Send(ctx context.Context, from, message string, to []string) httperrors.HTTPError {
	if err := smtp.SendMail(srv.host+":1025", srv.auth, from, to, []byte(message)); err != nil {
		logrus.Error("mail wasn't sent ", "err ", err)
		return httperrors.NewInternalServerError(err)
	}
	return nil
}
