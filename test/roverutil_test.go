package test

import (
	"github.com/tharun208/mars-rover/roverutil"
	"testing"
)

func TestSetXPosition(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.SetXPosition(4)
	if rover.RowPosition != 4 {
		t.Errorf("X Position was incorrect, got: %d, want: %d.", rover.RowPosition, 4)
	}
}

func TestSetYPosition(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.SetYPosition(3)
	if rover.ColumnPosition != 3 {
		t.Errorf("Y Position was incorrect, got: %d, want: %d.", rover.ColumnPosition, 4)
	}
}

func TestSetPosition(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.SetPosition("E")
	if rover.Orientation != "E" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.Orientation, "E")
	}
}
func TestTurnRoverLeft(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.TurnRoverLeft()
	if rover.GetPosition() != "W" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "W")
	}
}

func TestTurnRoverRight(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.TurnRoverRight()
	if rover.GetPosition() != "E" {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "E")
	}
}

func TestProcessCommands(t *testing.T) {
	rover := roverutil.Rover{RowPosition: 1, ColumnPosition: 2, Orientation: "N"}
	rover.ProcessCommands("LM")
	if rover.GetPosition() != "W" && rover.GetYPosition() != 2 {
		t.Errorf("Orientation was incorrect, got: %s, want: %s.", rover.GetPosition(), "W")
		t.Errorf("Y Position was incorrect, got: %d, want: %d.", rover.GetYPosition(), 2)
	}
}