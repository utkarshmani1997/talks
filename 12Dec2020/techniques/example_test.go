package main

import (
	"testing"
	"time"
)

func TestSequential(t *testing.T) {
	// t.Fatal("not implemented")
	t.Run("Some test", func(t *testing.T) { time.Sleep(1 * time.Second) })
	t.Run("Some other test", func(t *testing.T) { time.Sleep(1 * time.Second) })
	t.Run("Some any other test", func(t *testing.T) { time.Sleep(1 * time.Second) })
}
