// Copyright (c) 2017 Alex Pliutau

package snake

type keyboardKey rune

type keyboardEvent struct {
	key string
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

func listenToKeyboard(evChan chan keyboardEvent) {
	for {
		var key = ""
		evChan <- keyboardEvent{key: key}
	}
}
