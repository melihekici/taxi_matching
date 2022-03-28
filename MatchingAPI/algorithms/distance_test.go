package algorithms

import (
	"fmt"
	"matching/models"
	"math"
	"testing"
)

func TestHaversine(t *testing.T) {
	coordinatesTable := [][2][2]float64{
		{{137.5, 42.1}, {130.6, 55.2}},
		{{-102.6, 40.4}, {104.6, -53.8}},
		{{170.0, -32.9}, {0, 0}},
		{{-3.8, 7.5}, {-47.1, -16.2}},
		{{-180.0, -90.0}, {180.0, 90.0}},
		{{26.26, -13.13}, {26.28, -13.15}},
	}
	expected := []float64{1267630, 17960160, 16209790, 5325710, 20015120, 2990}

	for i, p := range coordinatesTable {
		p1, _ := models.NewPoint(p[0])
		p2, _ := models.NewPoint(p[1])
		calculated := Haversine(*p1, *p2)
		expect := expected[i]

		if (math.Abs(calculated-expect) / expect) > 0.002 {
			t.Error(fmt.Sprintf("Expected %.5f, Got %.5f", expect, calculated))
		}
	}
}
