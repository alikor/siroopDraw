package main

import (
	"testing"
)

const (
	border = 2
	rows   = 4
	cols   = 20
)

func TestNewCanvas(t *testing.T) {
	result := NewCanvas(cols, rows)

	if result.height != (rows + border) {
		t.Errorf("canvas is not the correct height")
	}

	if result.width != (cols + border) {
		t.Errorf("canvas is not the correct width")
	}

	if len(result.data) != ((cols + border) * (rows + border)) {
		t.Errorf("Array should be large enough to hold all elements ")
	}
}

func TestSetElementWhenSettingElementGetterShouldReturnCorrectValue(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	canvas.SetElement(1, 1, "X")
	el, _ := canvas.Element(1, 1)

	if el != "X" {
		t.Error("Canvas element was not set")
	}

}

func TestElementShouldReturnNillIfOutOfRange(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	_, error := canvas.Element(-1, -1)

	if error == nil {
		t.Error("canvas should return an error if element out of range", error)
	}

}

func TestElementShouldReturnErrorIfRangeColIsToLarge(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	_, error := canvas.Element(cols+2, 0)

	if error == nil {
		t.Error("canvas should return an error if col value is to large", error)
	}

}

func TestElementShouldReturnErrorIfRangeColIsToSmall(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	_, error := canvas.Element(-1, 0)

	if error == nil {
		t.Error("canvas should return an error if col value is to small", error)
	}

}

func TestElementShouldReturnErrorIfRangeRowIsToLarge(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	_, error := canvas.Element(0, rows+2)

	if error == nil {
		t.Error("canvas should return an error if row value is to large", error)
	}

}

func TestElementShouldReturnErrorIfRangeRowIsToSmall(t *testing.T) {
	canvas := NewCanvas(cols, rows)

	_, error := canvas.Element(0, -1)

	if error == nil {
		t.Error("canvas should return an error if row value is to small", error)
	}

}

func TestNeighbourElementsShouldReturnTopBottomLeftRightElements(t *testing.T) {

	canvas := NewCanvas(3, 3)
	results := canvas.NeighbourElements(2, 2)

	if len(results) != 4 {
		t.Fatal("Did not contain all 4 Neighbour elements")
	}

	topElement := Point{2, 1}
	if results[0] != topElement {
		t.Error("Did not return top element")
	}

	bottomElement := Point{2, 3}
	if results[1] != bottomElement {
		t.Error("Did not return bottom element")
	}

	leftElement := Point{1, 2}
	if results[2] != leftElement {
		t.Error("Did not return left element", results)
	}

	rightElement := Point{3, 2}
	if results[3] != rightElement {
		t.Error("Did not return right element")
	}
}

func TestNeighbourElementsShouldReturnBottomLeRightElementsOnly(t *testing.T) {

	canvas := NewCanvas(3, 3)
	results := canvas.NeighbourElements(0, 0)

	if len(results) != 2 {
		t.Fatal("Did not contain all 2 Neighbour elements")
	}

	bottomElement := Point{0, 1}
	if results[0] != bottomElement {
		t.Error("Did not return bottom element")
	}

	rightElement := Point{1, 0}
	if results[1] != rightElement {
		t.Error("Did not return right element")
	}
}

func TestNeighbourElementsShouldReturnTopLeftElementsOnly(t *testing.T) {

	canvas := NewCanvas(3, 3)
	results := canvas.NeighbourElements(4, 4)

	if len(results) != 2 {
		t.Fatal("Did not contain all 2 Neighbour elements", canvas)
	}

	topElement := Point{4, 3}
	if results[0] != topElement {
		t.Error("Did not return Top element", results)
	}

	rightElement := Point{3, 4}
	if results[1] != rightElement {
		t.Error("Did not return right element")
	}
}

func TestNeighbourElementsShouldReturnTopRightElementsOnly(t *testing.T) {

	canvas := NewCanvas(3, 3)
	results := canvas.NeighbourElements(0, 4)

	if len(results) != 2 {
		t.Fatal("Did not contain all 2 Neighbour elements", canvas)
	}

	topElement := Point{0, 3}
	if results[0] != topElement {
		t.Error("Did not return top element", results)
	}

	rightElement := Point{1, 4}
	if results[1] != rightElement {
		t.Error("Did not return right element")
	}
}

func TestNeighbourElementsShouldReturnLeftBottomElementsOnly(t *testing.T) {

	canvas := NewCanvas(3, 3)
	results := canvas.NeighbourElements(4, 0)

	if len(results) != 2 {
		t.Fatal("Did not contain all 2 Neighbour elements", canvas)
	}

	bottomElement := Point{4, 1}
	if results[0] != bottomElement {
		t.Error("Did not return bottom element")
	}

	leftElement := Point{3, 0}
	if results[1] != leftElement {
		t.Error("Did not return left element", results)
	}

}
