package main

import (
	"testing"
)

func TestCommands(t *testing.T) {
	//define a struct affected directly
	commands := []struct {
		command   string
		arguments string
		result    string
	}{
		{"mute", "true", "ke 00 00"},
		{"mute", "false", "ke 00 01"},
		{"volume", "10", "kf 00 0a"},
		{"volume", "0", "kf 00 00"},
		{"volume", "100", "kf 00 64"},
	}

	initializeCommands()

	for _, tt := range commands {
		res, err := m[tt.command](tt.arguments)

		if err != nil {
			t.Errorf("Error occured")
		}
		if res != tt.result {
			t.Errorf("Error: got %s, expected %s", res, tt.result)
		}
	}
}
