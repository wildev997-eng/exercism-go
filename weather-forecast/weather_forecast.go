// Package weather means that this file belong to the weather package.
package weather

var (
	//CurrentCondition means that current condition will always be a string and it ask the user for current condition.
	CurrentCondition string
	//CurrentLocation means that current location will always be a string and it ask user for the current location.
	CurrentLocation string
)

// Forecast function ask the name and weather condition of the city and then it return the information to the user.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
