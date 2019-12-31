package day01

import "testing"

func TestCalculateFuelForMass(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		got := calculateFuelForMass(c.in)
		if got != c.want {
			t.Errorf("calculateFuelForMass(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCalculateFuelForModule(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		got := calculateFuelForModule(c.in)
		if got != c.want {
			t.Errorf("calculateFuelForModule(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
