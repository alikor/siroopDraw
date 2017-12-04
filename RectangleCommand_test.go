package main

import (
	"testing"
)

func TestNewRectangleCommand(t *testing.T) {
	rectangleCmd := NewRectangleCommand([]string{"R", "1", "2", "3", "4"})

	if rectangleCmd.X1 != 1 {
		t.Errorf("X1 was not parsed")
	}

	if rectangleCmd.Y1 != 2 {
		t.Errorf("Y2 was not parsed")
	}

	if rectangleCmd.X2 != 3 {
		t.Errorf("X2 was not parsed")
	}

	if rectangleCmd.Y2 != 4 {
		t.Errorf("Y2 was not parsed")
	}
}

func TestRectangleCommandExecuteDraw(t *testing.T) {
	canvas := NewCanvas(3, 3)
	rectangleCmd := NewRectangleCommand([]string{"R", "0", "0", "4", "4"})

	rectangleCmd.Execute(canvas)

	result := make([][]string, 5)
	for i := range result {
		result[i] = make([]string, 5)
	}

	for i := 0; i < canvas.width; i++ {
		for j := 0; j < canvas.height; j++ {
			result[j][i], _ = canvas.Element(i, j)
		}
	}

	for i := range result[0] {
		if result[0][i] != "x" {
			t.Error("top element was not set")
		}
	}

	for i := range result[len(result)-1] {
		if result[0][i] != "x" {
			t.Error("bottom element was not set")
		}
	}

	for index := 1; index < (len(result) - 1); index++ {
		line := result[index]
		if line[0] != "x" {
			t.Error("Left column has not been set")
		}

		if line[len(line)-1] != "x" {
			t.Error("right column has not been set")
		}
	}
}
