package automation

import (
	"github.com/go-vgo/robotgo"
)

// TypeText types the specified text at the current cursor position
func TypeText(text string) error {
	robotgo.TypeStr(text)
	return nil
}

// PressKey presses a single key
func PressKey(key string) error {
	robotgo.KeyTap(key)
	return nil
}

// PressKeyCombo presses a key combination (e.g., "ctrl", "c")
func PressKeyCombo(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	if len(keys) == 1 {
		robotgo.KeyTap(keys[0])
		return nil
	}

	// Convert []string to []interface{} for robotgo.KeyTap
	modifiers := make([]interface{}, len(keys)-1)
	for i, key := range keys[:len(keys)-1] {
		modifiers[i] = key
	}
	robotgo.KeyTap(keys[len(keys)-1], modifiers...)
	return nil
}

// HoldKey holds down a key
func HoldKey(key string) error {
	robotgo.KeyToggle(key, "down")
	return nil
}

// ReleaseKey releases a held key
func ReleaseKey(key string) error {
	robotgo.KeyToggle(key, "up")
	return nil
}

// TypeWithDelay types text with a delay between characters (milliseconds)
func TypeWithDelay(text string, delay int) error {
	for _, char := range text {
		robotgo.TypeStr(string(char))
		robotgo.MilliSleep(delay)
	}
	return nil
}
