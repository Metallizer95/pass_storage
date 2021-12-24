package passport

import (
	"math"
	"strconv"
)

func (p *Passport) GetTowerByCoordinate(longitude, latitude float64) {
	minDistance := math.Inf(1)
	inputPoint := coordinatePoints{x: longitude, y: latitude}
	for _, tower := range p.Towers.Towers {

		long, err := strconv.ParseFloat(tower.Longitude, 64)
		if err != nil {
			continue
		}

		lat, err := strconv.ParseFloat(tower.Latitude, 64)
		if err != nil {
			continue
		}
		c := coordinatePoints{x: long, y: lat}
		distance := c.findSquareDistance(inputPoint)
		if distance < minDistance {
			minDistance = distance
		}
		// TODO question: Should I exit from loop if distance start increase
	}
}

type coordinatePoints struct {
	x float64
	y float64
}

func (cp *coordinatePoints) findSquareDistance(other coordinatePoints) float64 {
	return math.Pow(math.Abs(other.x-cp.x), 2.0) + math.Pow(math.Abs(other.y-cp.y), 2.0)
}
