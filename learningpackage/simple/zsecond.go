package simple

import "fmt"

func init() {
	fmt.Println("Ssecond package initialized")
}

// Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculated(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
