package main

import (
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Some test", func(t *testing.T) {
		t.Parallel()
		time.Sleep(1 * time.Second)
	})
	t.Run("Some other test", func(t *testing.T) {
		t.Parallel()
		time.Sleep(1 * time.Second)
	})
	t.Run("Some any other test", func(t *testing.T) {
		t.Parallel()
		time.Sleep(1 * time.Second)
	})
}
