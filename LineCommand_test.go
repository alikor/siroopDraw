package main

import (
	"strconv"
	"testing"
)

func TestNewLineCommand(t *testing.T) {
	lineCmd := NewLineCommand([]string{"L", "1", "2", "3", "4"})

	if lineCmd.X1 != 1 {
		t.Errorf("X1 was not parsed")
	}

	if lineCmd.Y1 != 2 {
		t.Errorf("Y2 was not parsed")
	}

	if lineCmd.X2 != 3 {
		t.Errorf("X2 was not parsed")
	}

	if lineCmd.Y2 != 4 {
		t.Errorf("Y2 was not parsed")
	}
}

func TestLineCommandExecuteDrawVerticalTopToBottom(t *testing.T) {
	canvas := NewCanvas(3, 3)
	lineCmd := NewLineCommand([]string{"L", "3", "0", "3", "4"})

	lineCmd.Execute(canvas)

	for row := 0; row < 5; row++ {
		el, _ := canvas.Element(3, row)
		if el != "x" {
			msg := "Element x5 y" + strconv.Itoa(row) + " does not contain x"
			t.Error(msg)
		}
	}
}

func TestLineCommandExecuteDrawVerticalBottomToTop(t *testing.T) {
	canvas := NewCanvas(3, 3)
	lineCmd := NewLineCommand([]string{"L", "3", "4", "3", "0"})

	lineCmd.Execute(canvas)

	for row := 0; row < 5; row++ {
		el, _ := canvas.Element(3, row)
		if el != "x" {
			msg := "Element x5 y" + strconv.Itoa(row) + " does not contain x"
			t.Error(msg)
		}
	}
}

func TestLineCommandExecuteDrawHorizontalLeftToRight(t *testing.T) {
	canvas := NewCanvas(3, 3)
	lineCmd := NewLineCommand([]string{"L", "0", "3", "5", "3"})

	lineCmd.Execute(canvas)

	for row := 0; row < 5; row++ {
		el, _ := canvas.Element(row, 3)
		if el != "x" {
			msg := "Element x" + strconv.Itoa(row) + " y3 does not contain x"
			t.Error(msg)
		}
	}
}

func TestLineCommandExecuteDrawHorizontalRightToLeft(t *testing.T) {
	canvas := NewCanvas(3, 3)
	lineCmd := NewLineCommand([]string{"L", "5", "3", "0", "3"})

	lineCmd.Execute(canvas)

	for row := 0; row < 5; row++ {
		el, _ := canvas.Element(row, 3)
		if el != "x" {
			msg := "Element x" + strconv.Itoa(row) + " y3 does not contain x"
			t.Error(msg)
		}
	}
}
