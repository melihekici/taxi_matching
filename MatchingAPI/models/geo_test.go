package models

import "testing"

func TestNewLatitude(t *testing.T) {
	for l := -180.; l <= 180.; l = l + 0.25 {
		_, err := newLatitude(l)
		if err != nil {
			t.Error("value ", l, " should be valid for latitude")
		}
	}
}

func TestNewLongitude(t *testing.T) {
	for l := -90.; l <= 90.; l = l + 0.25 {
		_, err := newLongitude(l)
		if err != nil {
			t.Error("value ", l, " should be valid for longitude")
		}
	}
}

func TestNewCoordinates(t *testing.T) {
	for lat := -180.; lat <= 180.; lat = lat + 0.25 {
		for long := 90.; long <= 90.; long = long + 0.25 {
			c, err := NewCoordinates([2]float64{lat, long})
			if err != nil {
				t.Errorf("value %v,%v should be valid for coordiantes", lat, long)
			}
			if c[0] != lat || c[1] != long {
				t.Errorf("expected (%v,%v) got (%v)", lat, long, c)
			}
		}
	}
}

func TestNewCoordinates_Fail(t *testing.T) {
	testTable := [][2]float64{
		{-181, 5},
		{191, 10},
		{15, -91},
		{14, 91},
		{190, 100},
	}

	for _, c := range testTable {
		_, err := NewCoordinates(c)
		if err == nil {
			t.Errorf("value %v,%v should not be valid for coordinates", c[0], c[1])
		}
	}
}

func TestNewPoint(t *testing.T) {
	for lat := -180.; lat <= 180.; lat = lat + 0.25 {
		for long := 90.; long <= 90.; long = long + 0.25 {
			p, err := NewPoint([2]float64{lat, long})
			if err != nil {
				t.Errorf("value %v,%v should be valid for coordiantes", lat, long)
			}
			if p.Coordinates[0] != lat || p.Coordinates[1] != long {
				t.Errorf("expected coordinates (%v,%v) got (%v)", lat, long, p.Coordinates)
			}
		}
	}
}

func TestNewPoint_Fail(t *testing.T) {
	testTable := [][2]float64{
		{-181, 5},
		{191, 10},
		{15, -91},
		{14, 91},
		{190, 100},
	}

	for _, c := range testTable {
		_, err := NewPoint(c)
		if err == nil {
			t.Errorf("value %v,%v should not be valid for coordiantes", c[0], c[1])
		}
	}
}
