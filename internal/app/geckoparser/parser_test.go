package geckoparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_GetPriceHistoryPeriod(t *testing.T) {
	g := GeckoParser{}
	actual, err := g.GetPriceHistoryPeriod("bitcoin", "365")
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)

}
