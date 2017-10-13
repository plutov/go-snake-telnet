// Copyright (c) 2017 Alex Pliutau

package snake

type keyboardKey rune

// KeyboardEvent type
type KeyboardEvent struct {
	Key string
}

func keyToDirection(k string) direction {
	switch k {
	case "a":
		return LEFT
	case "s":
		return DOWN
	case "d":
		return RIGHT
	case "w":
		return UP
	default:
		return 0
	}
}
