package main

import (
	"testing"
)

func TestNewCanvasCommand(t *testing.T) {
	canvasCmd := NewCanvasCommand([]string{"C", "20", "4"})

	if canvasCmd.height != 4 {
		t.Errorf("canvas command is not the correct height")
	}

	if canvasCmd.width != 20 {
		t.Errorf("canvas command is not the correct width")
	}
}

func TestCanvasCommandExecute(t *testing.T) {
	canvasCmd := NewCanvasCommand([]string{"C", "20", "4"})

	if canvasCmd.height != 4 {
		t.Errorf("canvas command is not the correct height")
	}

	if canvasCmd.width != 20 {
		t.Errorf("canvas command is not the correct width")
	}
}

func TestExecuteDrawBorder(t *testing.T) {
	canvasCmd := NewCanvasCommand([]string{"C", "2", "2"})
	canvas := Canvas{}
	canvasCmd.Execute(&canvas)

	result := make([][]string, 4)
	for i := range result {
		result[i] = make([]string, 4)
	}

	for i := 0; i < canvas.width; i++ {
		for j := 0; j < canvas.height; j++ {
			result[j][i], _ = canvas.Element(i, j)
		}
	}

	for i := range result[0] {
		if result[0][i] != "-" {
			t.Error("Canvas element was not set")
		}
	}

	for i := range result[len(result)-1] {
		if result[0][i] != "-" {
			t.Error("Canvas element was not set")
		}
	}

	for index := 1; index < (len(result) - 1); index++ {
		line := result[index]
		if line[0] != "|" {
			t.Error("Left column has not been set")
		}

		if line[len(line)-1] != "|" {
			t.Error("right column has not been set")
		}
	}
}
