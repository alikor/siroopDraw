package main

import (
	"strconv"
)

type CanvasCommand struct {
	width  int
	height int
}

func (c *CanvasCommand) Execute(canvas *Canvas) {
	tmp := NewCanvas(c.width, c.height)
	canvas.data = tmp.data
	canvas.width = tmp.width
	canvas.height = tmp.height
	drawBorder(canvas)
}

func drawBorder(c *Canvas) {

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {

			c.SetElement(j, i, " ")

			if j == 0 || j == (c.width-1) {
				c.SetElement(j, i, "|")
			}

			if i == 0 || i == (c.height-1) {
				c.SetElement(j, i, "-")
			}

		}
	}
}

func NewCanvasCommand(tokens []string) *CanvasCommand {
	width, _ := strconv.Atoi(tokens[1])
	height, _ := strconv.Atoi(tokens[2])

	cmd := CanvasCommand{
		width,
		height,
	}
	return &cmd
}
