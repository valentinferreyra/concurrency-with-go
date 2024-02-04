package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thingTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("expected 5, got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	tests := []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quarter second delay", 250 * time.Millisecond},
		{"half second delay", 500 * time.Millisecond},
	}

	for _, tt := range tests {
		orderFinished = []string{}
		eatTime = tt.delay
		sleepTime = tt.delay
		thingTime = tt.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: expected 5, got %d", tt.name, len(orderFinished))
		}
	}
}
