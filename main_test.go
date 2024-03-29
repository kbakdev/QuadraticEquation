package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	a, b, c float64
	out     [2]float64
}{
	{1, 2, 3, [2]float64{0.41421356237309515, -2.414213562373095}}, // if below zero
	{2, 3, 4, [2]float64{0.4489578808281798, -1.9489578808281798}}, // random
	{1, 3, -4, [2]float64{-0.5, -5.5}},                             // if greater than zero
	{1, 2, 1, [2]float64{-1, -1}},                                  // discriminant == 0

}

func TestQuadraticEquation(t *testing.T) {

	for _, tC := range testCases {
		testname := fmt.Sprintf("a=%v,b=%v,c=%v,want:%v", tC.a, tC.b, tC.c, tC.out)
		t.Run(testname, func(t *testing.T) {
			ans := QuadraticEquation(tC.a, tC.b, tC.c)
			if ans != tC.out {
				t.Errorf("got %v, want %v", ans, tC.out)
			}

		})
	}
}

func TestScan(t *testing.T) {
	testCases := []struct {
		a, b, c float64
		out     [3]float64
	}{
		{1, 2, 3, [3]float64{1, 2, 3}},
		{2, 3, 4, [3]float64{2, 3, 4}},
		{3, 4, 5, [3]float64{3, 4, 5}},
	}
	for _, tC := range testCases {
		testname := fmt.Sprintf("a=%v,b=%v,c=%v,want:%v", tC.a, tC.b, tC.c, tC.out)
		t.Run(testname, func(t *testing.T) {
			ans := Scan(tC.a, tC.b, tC.c)
			if ans != tC.out {
				t.Errorf("got %v, want %v", ans, tC.out)
			}

		})
	}
}

func TestAsk(t *testing.T) {
	testCases := []struct {
		checked string
		out     string
	}{
		{"Y", "Y"},
		{"X", "X"},
	}
	for _, tC := range testCases {
		t.Run(tC.checked, func(t *testing.T) {
			ans := Ask(tC.checked)
			if ans != tC.out {
				t.Errorf("got %v, want %v", ans, tC.out)
			}
		})
	}
}
