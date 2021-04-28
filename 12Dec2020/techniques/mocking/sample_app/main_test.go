package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/utkarshmani1997/talks/12Dec2020/techniques/mocking/sample_app/database/mock_database"
)

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

func TestInsertUser2(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockUser := mock_database.NewMockUsers(mockCtl)
	mockUser.EXPECT().Insert().Return(0, fmt.Errorf("%v", "insert error")).Times(1)
	_, err := insertUser(mockUser)
	if err == nil {
		t.Fatal(err)
	}
}
