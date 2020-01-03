package main

import "testing"

func TestValueValid(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
		{223456, true},
	}
	for _, c := range cases {
		got := passwordValid(c.in)
		if got != c.want {
			t.Errorf("passwordValid(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestValueValid2(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{112233, true},
		{223450, false},
		{123789, false},
		{223456, true},
		{123444, false},
		{111122, true},
	}
	for _, c := range cases {
		got := passwordValid2(c.in)
		if got != c.want {
			t.Errorf("passwordValid2(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
