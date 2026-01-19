package sharecalc

import (
	"math"
)

func CalculateShares(betAmount float64, members []MemberBet) ([]CalculationResult, error) {
	var results []CalculationResult

	currentTakenPercent := 0.0

	for _, m := range members {

		actualHoldPercent := m.Keep - currentTakenPercent

		if actualHoldPercent < 0 {
			actualHoldPercent = 0
		}

		amount := (betAmount * actualHoldPercent) / 100.0

		results = append(results, CalculationResult{
			UserID:      m.UserID,
			Level:       m.Level,
			HoldPercent: toFixed(actualHoldPercent, 2),
			HoldAmount:  toFixed(amount, 2),
		})

		currentTakenPercent = m.Keep
	}

	return results, nil
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}
