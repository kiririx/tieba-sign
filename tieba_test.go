package main

import (
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	ti := time.NewTicker(time.Second)
	for a := range ti.C {
		t.Log(a)
	}
}
