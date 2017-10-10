package main

type keyboardEventType int

type keyboardKey rune

const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       string
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
		evChan <- keyboardEvent{eventType: MOVE, key: key}
	}
}
