package quote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/email"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/quote"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/ticker"
	"net/http"
)

func (srv service) QuoteConsumer() error {
	ctx := context.Background()
	for {
		select {
		case v, ok := <-srv.msgChan:
			if !ok {
				close(srv.msgChan)
				return errors.New("read from channel error")
			}

			var msg ticker.Ticker

			err := json.Unmarshal(v, &msg)
			if err != nil {
				return err
			}

			tx, err := srv.pg.NewTXContext(ctx)
			if err != nil {
				return err
			}

			alerts, err := tx.GetAlertByPriceAndCurrency(msg.FromSymbol, msg.ToSymbol, msg.Price)
			if err != nil {
				logrus.Error("couldn't read data from DB", "err", err)
				continue
			}

			for i := range alerts {
				a := alerts[i]
				emailNotification := email.Notification{}
				emailNotification.Text = fmt.Sprintf("Currency %s/%s has reached price %v",
					msg.FromSymbol,
					msg.ToSymbol, msg.Price)
				u, err := tx.GetUser(a.UserID)
				if err != nil {
					logrus.Error("couldn't read data from DB", "err", err)
					continue
				}

				emailNotification.To = u.Email
				emailNotification.From = "priceAlert@pa.com"
				b, err := json.Marshal(emailNotification)
				if err != nil {
					continue
				}

				//TODO create a client
				_, err = http.Post(srv.urlNS, "application/json", bytes.NewBuffer(b))
				if err != nil {
					logrus.Error("couldn't sent notification ", "err ", err)
					continue
				}

				_, err = tx.DeleteAlert(a.ID)
				if err != nil {
					logrus.Error("couldn't delete alert ", a.ID.String())
					continue
				}

			}
			if err := tx.Commit(); err != nil {
				continue
			}
		}
	}
}

func (srv service) UpdateCurrency(req quote.AddNewCurrencyRequest) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	r, err := http.Post(srv.urlQuote, "application/json", bytes.NewBuffer(b))
	if err != nil {
		logrus.Error("couldn't upd currencies", "err ", err)
		return err
	}

	if r.StatusCode != http.StatusOK {
		logrus.Error("couldn't upd currencies", "status code ", r.StatusCode)
		return fmt.Errorf("status code %d", r.StatusCode)
	}

	return nil
}
