package seismo

import (
	"hayai/constants"

	"github.com/jftuga/geodist"
)

func CalculateEquivalentMagnitude(earthquakeMagnitude float64, earthquakeLat, earthquakeLon, userLat, userLon float64) float64 {
	var earthquake = geodist.Coord{Lat: earthquakeLat, Lon: earthquakeLon}
	var user = geodist.Coord{Lat: userLat, Lon: userLon}
	_, km, _ := geodist.VincentyDistance(earthquake, user)
	return earthquakeMagnitude - km*constants.MagnitudeDecreasePerKM
}
