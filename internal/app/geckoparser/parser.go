package geckoparser

import (
	"strconv"
	"sync"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	coingecko "github.com/superoo7/go-gecko/v3"
)

type GeckoParser struct {
	sync.Mutex
	Currencies []string
}

func New(wallets []string) *GeckoParser {
	return &GeckoParser{Currencies: wallets}
}

func (g *GeckoParser) GetPriceHistoryPeriod(currencyName string, days string) ([]model.History, error) {
	cg := coingecko.NewClient(nil)
	currencyHistory, err := cg.CoinsIDMarketChart(currencyName, "usd", days)
	if err != nil {
		return nil, err
	}

	historyByDays := make([]model.History, 0, len(*currencyHistory.Prices))
	for _, value := range *currencyHistory.Prices {
		m := model.History{
			Timestamp: floatToString(value[0]),
			Price:     value[1],
		}
		historyByDays = append(historyByDays, m)
	}
	return historyByDays, nil
}

func floatToString(input_num float32) string {

	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 0, 64)
}
