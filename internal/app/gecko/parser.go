package gecko

import (
	"fmt"
	"net/http"
	"time"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	coingecko "github.com/superoo7/go-gecko/v3"
)

type Task struct {
	wallet string
	date   string
}

type Gecko struct {
}

func (e *Gecko) Run() error {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	gc := coingecko.NewClient(httpClient)
	if _, err := gc.Ping(); err != nil {
		return err
	}

	fmt.Println(gc)
	return nil
}

func GetPriceForLasyYear(wallet string) ([]model.PricePerTimestamp, error) {
	cg := coingecko.NewClient(nil)
	info, err := cg.CoinsIDMarketChart(wallet, "usd", "365")
	if err != nil {
		return nil, err
	}

	prices := make([]model.PricePerTimestamp, 0, len(*info.Prices))

	for _, v := range *info.Prices {
		m := &model.PricePerTimestamp{}
		m.Timestamp = int64(v[0])
		m.Price = v[1]
		prices = append(prices, *m)
	}

	return prices, nil
}

func GetPriceForWalletsForLasyYear(wallets ...string) (map[string][]model.PricePerTimestamp, error) {
	results := make(map[string][]model.PricePerTimestamp, len(wallets))
	for _, wallet := range wallets {
		prices, err := GetPriceForLasyYear(wallet)
		if err != nil {
			return nil, err
		}
		results[wallet] = prices
	}
	return results, nil
}
