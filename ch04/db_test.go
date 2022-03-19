package ch04

import (
	"GolandProjectsunit_test_golang/ch04/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDB(ctrl)
	m.
		EXPECT().
		Get(gomock.Eq("liwenzhou.com")).
		Return(1, nil).
		Times(1)

	if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
		t.Fatal()
	}
}
