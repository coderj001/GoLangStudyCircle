package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(7, 3)
	want := float64(10)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSub(t *testing.T) {
	got := sub(7, 3)
	want := float64(4)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestMulti(t *testing.T) {
	got := multi(7, 3)
	want := float64(21)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestDiv(t *testing.T) {
	got := div(9, 3)
	want := float64(3)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSin(t *testing.T) {
	tolerance := 0.000001
	got := sin(math.Pi)
	want := float64(0.0)
	if math.Abs(got-want) > tolerance {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCos(t *testing.T) {
	got := cos(math.Pi)
	want := -1.0
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestTan(t *testing.T) {
	tolerance := 0.000001
	got := tan(math.Pi)
	want := 0.0
	if math.Abs(got-want) > tolerance {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSqr(t *testing.T) {
	got := sqr(4)
	want := float64(2)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
