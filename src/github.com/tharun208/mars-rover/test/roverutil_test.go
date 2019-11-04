package test

import (
	"github.com/tharun208/mars-rover/pkg/roverutil"
	"testing"
)

func TestSetXPosition(t *testing.T) {
	rover := roverutil.NewRover()
	rover.SetXPosition(4)
	if rover.GetXPosition() != 4 {
		t.Errorf("X Position was incorrect, got: %d, want: %d.", rover.GetXPosition(), 4)
	}
}

func TestSetYPosition(t *testing.T) {
	rover := roverutil.NewRover()
	rover.SetYPosition(3)
	if rover.GetYPosition() != 3 {
		t.Errorf("Y Position was incorrect, got: %d, want: %d.", rover.GetYPosition(), 4)
	}
}

func TestSetPosition(t *testing.T) {
	rover := roverutil.NewRover()
	rover.SetPosition("E")
	if rover.GetPosition() != "E" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "E")
	}
}

func TestTurnRoverLeft(t *testing.T) {
	rover := roverutil.NewRover()
	rover.TurnRoverLeft()
	if rover.GetPosition() != "W" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "W")
	}
}

func TestTurnRoverRight(t *testing.T) {
	rover := roverutil.NewRover()
	rover.TurnRoverRight()
	if rover.GetPosition() != "E" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "E")
	}
}

func TestProcessCommands(t *testing.T) {
	rover := roverutil.NewRover()
	rover.ProcessCommands("LM")
	if rover.GetPosition() != "W" && rover.GetYPosition() != 2 {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "W")
		t.Errorf("Y Position was incorrect, got: %d, want: %d.", rover.GetYPosition(), 2)
	}
}