package main

import (
	"strconv"
)

type RectangleCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (rectangleCmd *RectangleCommand) Execute(c *Canvas) {

	width := rectangleCmd.X2 - rectangleCmd.X1
	height := rectangleCmd.Y2 - rectangleCmd.Y1

	line1 := LineCommand{
		rectangleCmd.X1,
		rectangleCmd.Y1,
		(rectangleCmd.X1 + width),
		rectangleCmd.Y1,
	}
	line1.Execute(c)

	line2 := LineCommand{
		rectangleCmd.X1,
		rectangleCmd.Y1,
		rectangleCmd.X1,
		(rectangleCmd.Y1 + height),
	}
	line2.Execute(c)

	line3 := LineCommand{
		rectangleCmd.X1,
		(rectangleCmd.Y1 + height),
		rectangleCmd.X2,
		rectangleCmd.Y2,
	}
	line3.Execute(c)

	line4 := LineCommand{
		(rectangleCmd.X1 + width),
		rectangleCmd.Y1,
		rectangleCmd.X2,
		rectangleCmd.Y2,
	}
	line4.Execute(c)
}

func NewRectangleCommand(tokens []string) *RectangleCommand {
	x1, _ := strconv.Atoi(tokens[1])
	y1, _ := strconv.Atoi(tokens[2])
	x2, _ := strconv.Atoi(tokens[3])
	y2, _ := strconv.Atoi(tokens[4])

	cmd := RectangleCommand{
		x1,
		y1,
		x2,
		y2,
	}
	return &cmd
}
