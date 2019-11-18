package test

import (
	"go-specification-sample/src/test/mock_human"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestHuman(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_human.NewMockHuman(ctrl)

	m.EXPECT().Greeting("hello").Return("hello kngfjsr")

	result := m.Greeting("hello")
	expect := "hello kngfjsr"

	if result != expect {
		t.Error("\n result:", result, "expect:", expect)
	}
	t.Log("finish TestHuman")
}
