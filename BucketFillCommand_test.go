package main

import (
	"testing"
)

func TestNewBucketFillCommand(t *testing.T) {
	bucketFillCmd := NewBucketFillCommand([]string{"B", "1", "2", "o"})

	if bucketFillCmd.X != 1 {
		t.Errorf("X was not parsed")
	}

	if bucketFillCmd.Y != 2 {
		t.Errorf("Y was not parsed")
	}

	if bucketFillCmd.Color != "o" {
		t.Errorf("Color was not parsed")
	}
}

func TestNewBucketFillCommandExecuteFillAll(t *testing.T) {
	canvas := NewCanvas(3, 3)

	bucketFillCmd := NewBucketFillCommand([]string{"B", "1", "1", "o"})
	bucketFillCmd.Execute(canvas)

	result := make([][]string, 5)
	for i := range result {

		result[i] = make([]string, 5)
	}

	for i := 0; i < canvas.width; i++ {
		for j := 0; j < canvas.height; j++ {
			result[j][i], _ = canvas.Element(i, j)
		}
	}

	for i := range result {
		for j := range result[i] {
			if result[i][j] != "o" {
				t.Error("Element not filled")
			}
		}
	}
}

func TestBucketFillCommandExecuteFillHalf(t *testing.T) {
	canvas := NewCanvas(3, 3)

	lineCmd := LineCommand{2, 0, 2, 4}
	lineCmd.Execute(canvas)
	bucketFillCmd := NewBucketFillCommand([]string{"B", "1", "1", "o"})
	bucketFillCmd.Execute(canvas)

	result := make([][]string, 5)
	for i := range result {

		result[i] = make([]string, 5)
	}

	for i := 0; i < canvas.width; i++ {
		for j := 0; j < canvas.height; j++ {
			result[j][i], _ = canvas.Element(i, j)
		}
	}

	for i := 0; i < len(result)/2; i++ {
		for j := 0; j < len(result[i])/2; j++ {
			if result[i][j] != "o" {
				t.Error("Element not filled")
			}
		}
	}
}
