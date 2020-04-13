package _1_simple_factory

import (
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("jerry")
	t.Log(s)
	//if s != "Hi, Tom" {
	//	t.Fatal("Type1 test fail")
	//}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("jerry")
	t.Log(s)
}
