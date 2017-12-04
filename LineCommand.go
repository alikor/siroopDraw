package main

import (
	"strconv"
)

type LineCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Dirty would refactor
func (c *LineCommand) Execute(canvas *Canvas) {
	if c.X1 == c.X2 {
		if c.Y1 < c.Y2 {
			for i := c.Y1; i <= c.Y2; i++ {
				canvas.SetElement(c.X1, i, "x")
			}
		} else {
			for i := c.Y1; i >= c.Y2; i-- {
				canvas.SetElement(c.X1, i, "x")
			}
		}
	} else {
		if c.X1 < c.X2 {
			for i := c.X1; i <= c.X2; i++ {
				canvas.SetElement(i, c.Y1, "x")
			}
		} else {
			for i := c.X1; i >= c.X2; i-- {
				canvas.SetElement(i, c.Y1, "x")
			}
		}

	}
}

func NewLineCommand(tokens []string) *LineCommand {
	X1, _ := strconv.Atoi(tokens[1])
	Y1, _ := strconv.Atoi(tokens[2])
	X2, _ := strconv.Atoi(tokens[3])
	Y2, _ := strconv.Atoi(tokens[4])

	cmd := LineCommand{
		X1,
		Y1,
		X2,
		Y2,
	}
	return &cmd
}
