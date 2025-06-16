package automation

import (
	"github.com/go-vgo/robotgo"
)

// Click performs a mouse click at the specified coordinates
func Click(x, y int) error {
	robotgo.Move(x, y)
	robotgo.Click()
	return nil
}

// MoveMouse moves the mouse cursor to the specified coordinates
func MoveMouse(x, y int) error {
	robotgo.Move(x, y)
	return nil
}

// DoubleClick performs a double click at the specified coordinates
func DoubleClick(x, y int) error {
	robotgo.Move(x, y)
	robotgo.Click("left", true)
	return nil
}

// RightClick performs a right click at the specified coordinates
func RightClick(x, y int) error {
	robotgo.Move(x, y)
	robotgo.Click("right")
	return nil
}

// GetMousePos returns the current mouse position
func GetMousePos() (int, int) {
	x, y := robotgo.GetMousePos()
	return x, y
}
