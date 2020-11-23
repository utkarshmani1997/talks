package main

import "testing"

type mockUser struct {
}

func (u mockUser) Insert() (int, error) {
	return 0, nil
}

func TestInsertUser(t *testing.T) {
	id, err := insertUser(mockUser{})
	if err != nil {
		t.Fatal(err)
	}
	if id != 0 {
		t.Fatalf("Expected id: %d, got: %d", 0, id)
	}
}
