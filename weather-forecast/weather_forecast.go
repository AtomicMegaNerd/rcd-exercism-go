// Package weather implements a single function called forecast that returns
// the current conditions for a specific city.
package weather

// CurrentCondition represents the current weather conditions for the specified location.
var CurrentCondition string

// CurrentLocation specifies which location they want to specify to get the current conditions.
var CurrentLocation string

// Forecast returns the current weather conditions for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
