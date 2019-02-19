// Package space calculates some space-related queries
package space

import (
	"errors"
)

// Planet is a string
type Planet string

// Age takes an argument in seconds and calculates
// how long that would be in relation to the planets
// orbital period.
func Age(seconds float64, Planet string) float64 {
	var spaceAge float64
	switch Planet {
	case "Earth":
		spaceAge = 1
	case "Mercury":
		spaceAge = 0.2408467
	case "Venus":
		spaceAge = 0.61519726
	case "Mars":
		spaceAge = 1.8808158
	case "Jupiter":
		spaceAge = 11.862615
	case "Saturn":
		spaceAge = 29.447498
	case "Uranus":
		spaceAge = 84.016846
	case "Neptune":
		spaceAge = 164.79132
	default:
		panic(errors.New("Planet not defined"))
	}
	return seconds / ((365.25 * spaceAge) * 86400.0)
}
