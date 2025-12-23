package indicators

import (
	"math"
	"testing"
)

func TestER(t *testing.T) {
	prices := []float64{100, 29, 31, 13, 60, 80}
	period := 5

	erValues := ER(prices, period)

	if len(erValues) < 1 {
		t.Errorf("Expected at least one ER value, got %d", len(erValues))
	}

	expectedER := 0.12658227848101267

	tolerance := 1e-9
	if math.Abs(erValues[len(erValues)-1]-expectedER) > tolerance {
		t.Errorf("Expected ER value to be %.9f, but got %.9f", expectedER, erValues[len(erValues)-1])
	}
}
