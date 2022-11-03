package maths

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(0, 0, 30), math.Pi},
		{SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{SimpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(GetTestName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)
			if c.angle != got {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{SimpleTime(0, 0, 30), Point{0, -1}},
		{SimpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(GetTestName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func SimpleTime(hours, mins, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, mins, seconds, 0, time.UTC)
}

func GetTestName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
