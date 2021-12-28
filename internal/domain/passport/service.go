package passport

import (
	"math"
	"strconv"
)

func (p *Passport) GetAllTowers() Towers {
	return p.Towers
}

func (p *Passport) GetTowerById(towerId string) *Tower {
	var tower *Tower
	// TODO: How I can do it more faster (change list to map as example)
	for _, t := range p.Towers.Towers {
		if t.ID == towerId {
			tower = &t
			break
		}
	}
	return tower
}

type minDistanceTower struct {
	distance float64
	tower    Tower
}

func (p *Passport) GetTowerByCoordinate(longitude, latitude float64) *Tower {
	minDistance := minDistanceTower{
		distance: math.Inf(1),
		tower:    Tower{},
	}
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
		if distance < minDistance.distance {
			minDistance.distance = distance
			minDistance.tower = tower
		}
		// TODO question: Should I exit from loop if distance start increase
	}
	return &minDistance.tower
}

type coordinatePoints struct {
	x float64
	y float64
}

func (cp *coordinatePoints) findSquareDistance(other coordinatePoints) float64 {
	return math.Pow(math.Abs(other.x-cp.x), 2.0) + math.Pow(math.Abs(other.y-cp.y), 2.0)
}
