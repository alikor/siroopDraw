package main

import (
	"reflect"
	"testing"
)

func TestParseForCanvasCommand(t *testing.T) {
	expected := CanvasCommand{}
	result := Parse("C 20 4")

	if reflect.TypeOf(result) != reflect.TypeOf(&expected) {
		t.Errorf("command not a canvas but was %s", reflect.TypeOf(&expected))
	}
}

func TestParseForQuiteCommand(t *testing.T) {
	expected := QuiteCommand{}
	result := Parse("Q")

	if reflect.TypeOf(result) != reflect.TypeOf(&expected) {
		t.Errorf("command not a quite but was %s", reflect.TypeOf(&expected))
	}
}

func TestParseForUnknownCommand(t *testing.T) {
	expected := UnknownCommand{}
	result := Parse("X")

	if reflect.TypeOf(result) != reflect.TypeOf(&expected) {
		t.Errorf("command not a quite but was %s", reflect.TypeOf(&expected))
	}
}
