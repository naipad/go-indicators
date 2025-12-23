package main

import (
	"fmt"

	"github.com/naipad/go-indicators/v1"
)

func main() {

	inHigh := []float64{
		45.7, 46.5, 47.1, 46.8, 47.3, 47.5, 48.1, 48.3, 48.6, 48.7,
		49.2, 49.6, 49.7, 50.1, 50.4, 50.7, 51.0, 51.3, 51.6, 51.8,
		52.0, 52.3, 52.6, 52.8, 53.1, 53.3, 53.7, 53.9, 54.2, 54.4,
		54.7, 55.0, 55.2, 55.4, 55.7, 55.9, 56.2, 56.4, 56.6, 56.9,
		57.1, 57.3, 57.5, 57.7, 58.0, 58.2, 58.4, 58.6, 58.9, 59.1,
	}

	inLow := []float64{
		44.2, 45.0, 45.6, 45.3, 45.5, 45.7, 46.1, 46.2, 46.4, 46.5,
		46.9, 47.2, 47.3, 47.6, 47.8, 48.0, 48.2, 48.4, 48.5, 48.7,
		48.9, 49.1, 49.3, 49.4, 49.6, 49.8, 50.0, 50.2, 50.4, 50.6,
		50.8, 51.0, 51.2, 51.4, 51.6, 51.8, 52.0, 52.2, 52.4, 52.6,
		52.8, 53.0, 53.2, 53.4, 53.6, 53.8, 54.0, 54.2, 54.4, 54.6,
	}

	inClose := []float64{
		45.2, 45.9, 46.3, 46.1, 46.7, 47.1, 47.4, 47.5, 47.8, 48.0,
		48.3, 48.6, 48.8, 49.1, 49.4, 49.7, 50.0, 50.3, 50.6, 50.8,
		51.0, 51.3, 51.6, 51.9, 52.1, 52.4, 52.6, 52.9, 53.1, 53.4,
		53.6, 53.9, 54.1, 54.3, 54.6, 54.8, 55.0, 55.2, 55.4, 55.6,
		55.8, 56.0, 56.2, 56.4, 56.6, 56.8, 57.0, 57.2, 57.4, 57.6,
	}

	inVolume := []float64{
		1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900,
		2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900,
		3000, 3100, 3200, 3300, 3400, 3500, 3600, 3700, 3800, 3900,
		4000, 4100, 4200, 4300, 4400, 4500, 4600, 4700, 4800, 4900,
		5000, 5100, 5200, 5300, 5400, 5500, 5600, 5700, 5800, 5900,
	}

	inTimePeriod := 14
	printNum := 5

	// ATR
	atrValues := indicators.ATR(inHigh, inLow, inClose, inTimePeriod)
	fmt.Println("ATR:", atrValues[len(atrValues)-printNum:])

	// BBANDS
	inNbDevUp := 2.0
	inNbDevDn := 2.0
	inMAType := indicators.Sma
	upper, middle, lower := indicators.BBANDS(inClose, inTimePeriod, inNbDevUp, inNbDevDn, inMAType)
	fmt.Println("BBANDS Upper:", upper[len(upper)-printNum:])
	fmt.Println("BBANDS Middle:", middle[len(middle)-printNum:])
	fmt.Println("BBANDS Lower:", lower[len(lower)-printNum:])

	// CorrelationCoefficient
	correlationValues := indicators.CorrelationCoefficient(inHigh, inLow, inTimePeriod)
	fmt.Println("Correlation Coefficient:", correlationValues[len(correlationValues)-printNum:])

	// ER
	erValues := indicators.ER(inClose, inTimePeriod)
	fmt.Println("ER:", erValues[len(erValues)-printNum:])

	// EMA
	emaValues := indicators.EMA(inClose, inTimePeriod)
	fmt.Println("EMA:", emaValues[len(emaValues)-printNum:])

	// MACD
	macdValues, signalValues, histogramValues := indicators.MACD(inClose, 12, 26, 9)
	fmt.Println("MACD:", macdValues[len(macdValues)-printNum:])
	fmt.Println("MACD Signal:", signalValues[len(signalValues)-printNum:])
	fmt.Println("MACD Histogram:", histogramValues[len(histogramValues)-printNum:])

	// MA
	maValues := indicators.MA(inClose, inTimePeriod, inMAType)
	fmt.Println("MA:", maValues[len(maValues)-printNum:])

	// MFI
	mfiValues := indicators.MFI(inHigh, inLow, inClose, inVolume, inTimePeriod)
	fmt.Println("MFI:", mfiValues[len(mfiValues)-printNum:])

	// OBV
	obvValues := indicators.OBV(inClose, inVolume)
	fmt.Println("OBV:", obvValues[len(obvValues)-printNum:])

	// RSI
	rsiValues := indicators.RSI(inClose, inTimePeriod)
	fmt.Println("RSI:", rsiValues[len(rsiValues)-printNum:])
}
