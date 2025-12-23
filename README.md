# go-indicators

Some indicator algorithms

## Example

```go
package main

import (
	"fmt"

	"github.com/naipad/go-indicators/v1"
)

func main() {
	h := []float64{45.7, 46.5, 47.1, 46.8, 47.3}
	l := []float64{44.2, 45.0, 45.6, 45.3, 45.5}
	c := []float64{45.2, 45.9, 46.3, 46.1, 46.7}
	// v := []float64{1000, 1100, 1200, 1300, 1400}
	inTimePeriod := 3
	values := indicators.ATR(h, l, c, inTimePeriod)
	fmt.Println("ATR:", values)
}

```
