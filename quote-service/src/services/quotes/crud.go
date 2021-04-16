package quotes

import "fmt"

func (srv service) AddCurrencyPair(fromSymbol, toSymbol string) error {
	_, err := srv.redisCli.SAdd(currenciesKey, fmt.Sprintf(coinbaseSubscribeFormat,toSymbol, fromSymbol))
	if err != nil {
		return err
	}

	srv.close <- struct{}{}

	return srv.Start()
}
