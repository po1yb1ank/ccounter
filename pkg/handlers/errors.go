package handlers

import "errors"

var (
	ErrorFailedToReset     = errors.New("failed to reset counter")
	ErrorFailedToIncrement = errors.New("failed to increment counter")
	ErrorFailedToDecrement = errors.New("failed to decrement counter")
	ErrorFailedToGet       = errors.New("failed to get counter value")

	ErrorWSConn = errors.New("failed to set WS connection")
)
