package main

import (
	"strconv"
)

type BucketFillCommand struct {
	X     int
	Y     int
	Color string
}

func (b *BucketFillCommand) Execute(c *Canvas) {
	Fill(b.X, b.Y, b.Color, c)
}

func Fill(x int, y int, color string, c *Canvas) {

	if el, err := c.Element(x, y); el != " " || err != nil {
		return
	}

	c.SetElement(x, y, color)
	for _, point := range c.NeighbourElements(x, y) {
		Fill(point.col, point.row, color, c)
	}

}

func NewBucketFillCommand(tokens []string) *BucketFillCommand {
	x, _ := strconv.Atoi(tokens[1])
	y, _ := strconv.Atoi(tokens[2])
	color := tokens[3]

	cmd := BucketFillCommand{
		x,
		y,
		color,
	}
	return &cmd
}
