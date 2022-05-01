package dategen

import (
	"fmt"
	"time"
)

func GetPastDays(count int) []string {
	start := time.Now().AddDate(0, 0, -count)
	end := start.AddDate(1, 0, 0)

	result := make([]string, 0, count)
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		y, m, d := d.Date()
		result = append(result, fmt.Sprintf("%02d-%02d-%d", d, m, y))
	}
	return result
}
