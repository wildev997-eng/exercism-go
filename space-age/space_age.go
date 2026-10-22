package spaceage

type Planet string

func Age(seconds float64, planet Planet) float64 {
	if planet == "Sun" {
		return -1.0
	}
	earthAge := seconds / 31557600
	divider := 0.0
	switch {
	case planet == "Mercury":
		divider = 0.2408467
	case planet == "Venus":
		divider = 0.61519726
	case planet == "Earth":
		divider = 1.0
	case planet == "Mars":
		divider = 1.8808158
	case planet == "Jupiter":
		divider = 11.862615
	case planet == "Saturn":
		divider = 29.447498
	case planet == "Uranus":
		divider = 84.016846
	case planet == "Neptune":
		divider = 164.79132
	}
	return earthAge / divider
}
