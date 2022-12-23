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
	got := divison(9, 3)
	want := float64(3)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSin(t *testing.T) {
	got := Sin(math.Pi)
	want := float64(0.0)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCos(t *testing.T) {
	got := Cos(math.Pi)
	want := -1.0
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestTan(t *testing.T) {
	got := Tan(math.Pi)
	want := 0.0
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSqr(t *testing.T) {
	got := Sqr(4)
	want := float64(2)
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
