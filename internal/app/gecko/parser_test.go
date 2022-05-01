package gecko

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPriceForLasyYear(t *testing.T) {
	val, _ := GetPriceForLasyYear("bitcoin")

	assert.GreaterOrEqual(t, len(val), 365)
}

func TestGetPriceForWalletsForLasyYear(t *testing.T) {
	wallets, _ := GetPriceForWalletsForLasyYear("bitcoin", "ethereum")
	assert.EqualValues(t, 2, len(wallets))
	for name, wallet := range wallets {
		fmt.Println(name)
		assert.GreaterOrEqual(t, len(wallet), 365)
	}
}
