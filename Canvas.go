package main

import (
	"errors"
)

type Canvas struct {
	width  int
	height int
	data   []string
}

type Point struct {
	col int
	row int
}

func (c *Canvas) SetElement(col int, row int, value string) {

	c.data[c.width*row+col] = value
}

func (c *Canvas) Element(col int, row int) (string, error) {
	if col < 0 || col >= c.width {
		return "", errors.New("col is out of range")
	}

	if row < 0 || row >= c.height {
		return "", errors.New("row is out of range")
	}

	return c.data[c.width*row+col], nil
}

func (c *Canvas) NeighbourElements(col int, row int) []Point {
	newSlice := make([]Point, 0)

	_, erTop := c.Element(col, row-1)
	if erTop == nil {
		newSlice = append(newSlice, Point{col, row - 1})
	}

	_, erBottom := c.Element(col, row+1)
	if erBottom == nil {
		newSlice = append(newSlice, Point{col, row + 1})
	}

	_, erLeft := c.Element(col-1, row)
	if erLeft == nil {
		newSlice = append(newSlice, Point{col - 1, row})
	}

	_, erRight := c.Element(col+1, row)
	if erRight == nil {
		newSlice = append(newSlice, Point{col + 1, row})
	}

	return newSlice

}

func NewCanvas(width int, height int) *Canvas {
	width += 2
	height += 2
	size := width * height
	canvas := Canvas{
		width:  width,
		height: height,
	}
	canvas.data = make([]string, size)
	for i := 0; i < canvas.height; i++ {
		for j := 0; j < canvas.width; j++ {

			canvas.SetElement(j, i, " ")
		}
	}
	return &canvas
}
