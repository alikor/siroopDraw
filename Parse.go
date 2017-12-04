package main

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute(canvas *Canvas)
}

type QuiteCommand struct {
}

func NewQuiteCommand() *QuiteCommand {
	cmd := QuiteCommand{}
	return &cmd
}

func (QuiteCommand) Execute(canvas *Canvas) {
	os.Exit(0)
}

type UnknownCommand struct {
	line string
}

func NewUnknownCommand(line string) *UnknownCommand {
	cmd := UnknownCommand{line}
	return &cmd
}

func (c *UnknownCommand) Execute(canvas *Canvas) {
	fmt.Printf("Unknown command: \"%q\"\n", c.line)
}

func Parse(line string) Command {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "C":
		return NewCanvasCommand(tokens)
	case "L":
		return NewLineCommand(tokens)
	case "R":
		return NewRectangleCommand(tokens)
	case "B":
		return NewBucketFillCommand(tokens)
	case "Q":
		return NewQuiteCommand()

	default:
		return NewUnknownCommand(line)
	}
}
