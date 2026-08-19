package main

import (
	"fmt"
	"os"
)

// predictive_analytics - Time series prediction engine
func predictive_analytics(path string) {
	fmt.Println("========================================")
	fmt.Println("  Predictive-Analytics")
	fmt.Println("  Time series prediction engine")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	predictive_analytics(path)
}
