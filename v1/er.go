package indicators

import "math"

// ER returns the Efficiency Ratio for the given period.
func ER(inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	if inTimePeriod < 1 || inTimePeriod > len(inClose) {
		return outReal
	}

	// Loop through each point in the data starting after the inTimePeriod
	for i := inTimePeriod; i < len(inClose); i++ {
		priceChange := inClose[i] - inClose[i-inTimePeriod] // Price change over the period
		totalMovement := 0.0

		// Calculate the total movement for the given period
		for j := i - inTimePeriod + 1; j <= i; j++ {
			totalMovement += math.Abs(inClose[j] - inClose[j-1])
		}

		// Calculate ER (Efficiency Ratio)
		if totalMovement != 0 {
			outReal[i] = math.Abs(priceChange) / totalMovement
		} else {
			outReal[i] = 0 // Avoid division by zero
		}
	}

	return outReal
}
