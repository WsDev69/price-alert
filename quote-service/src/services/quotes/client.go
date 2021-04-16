package quotes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/models"
)

func (srv service) Start() error {
	var subs []string
	err := srv.redisCli.SMembers(currenciesKey, &subs)
	if err != nil {
		return err
	}

	jsonObj := models.WSJSONParams{Action: wsActionSub, Subs: subs}
	s, err := json.Marshal(jsonObj)
	if err != nil {
		return err
	}
	err = srv.conn.WriteMessage(websocket.TextMessage, s)
	if err != nil {
		return err
	}

	go srv.listener()

	return nil
}

func (srv service) listener() {
	srv.conn.SetPingHandler(func(appData string) error {
		return srv.conn.WriteMessage(websocket.PongMessage, nil)
	})
	for {
		select {
		case <-srv.close:
			return
		default:
			{
				_, message, err := srv.conn.ReadMessage()
				if err != nil {
					logrus.Error("read err:", err)
					continue
				}
				var msg models.WSResponse
				err = json.Unmarshal(message, &msg)
				if err != nil {
					logrus.Error("couldn't unmarshal response", err)
					logrus.Debug("data", string(message))
					continue
				}

				if msg.Type == wsTypeTicker {
					member := fmt.Sprintf(coinbaseSubscribeFormat, msg.ToSymbol, msg.FromSymbol)
					isMember, err := srv.redisCli.SIsMember(currenciesKey, member)
					if err != nil {
						logrus.Error("SIsMember err", err)
						continue
					}
					if isMember {
						if err := srv.messageBrokerSrv.Publish("currencies", message); err != nil {
							logrus.Error("message wasn't published, err", err)
						} else {
							logrus.Debug("message sent")
						}
					}
				}
			}
		}
	}
}
