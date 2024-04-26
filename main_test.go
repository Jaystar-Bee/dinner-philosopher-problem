package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {

	sleepTime = 0 * time.Second
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	main()

	if len(orderOfFinish) != 5 {
		t.Errorf("Expected 5, got %d", len(orderOfFinish))
	}
}
