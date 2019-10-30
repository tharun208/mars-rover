package roverutil

type Rover struct {
    rowPosition int
    columnPosition int
    orientation string
}

func (r *Rover) SetXPosition(x int) {
 	r.rowPosition = x
}
func (r Rover) GetXPosition() int {
 	return r.rowPosition
}

func (r *Rover) SetYPosition(y int) {
 	r.columnPosition = y
}
func (r Rover) GetYPosition() int {
 	return r.columnPosition
}

func (r *Rover) SetPosition(x string) {
 	r.orientation = x
 }
func (r Rover) GetPosition() string {
 	return r.orientation
}

func (r *Rover) ProcessCommands(commands string)  {
    for _, command := range commands {
     if string(command) == "L" {
        r.turnRoverLeft()
     } else if string(command) == "R" {
        r.turnRoverRight()
     } else if string(command) == "M" {
        r.moveRover()
     }
    }
}

func (r *Rover) turnRoverLeft() {
    currentOrientation := r.GetPosition();
    switch currentOrientation {
    case "N":
      r.SetPosition("W")
    case "S":
      r.SetPosition("E")
    case "E":
       r.SetPosition("N")
    case "W":
       r.SetPosition("S")
    }
}

func (r *Rover) turnRoverRight() {
    currentOrientation := r.GetPosition()
    switch currentOrientation {
    case "N":
      r.SetPosition("E")
    case "S":
      r.SetPosition("W")
    case "E":
       r.SetPosition("S")
    case "W":
       r.SetPosition("N")
    }
}

func (r *Rover) moveRover() {
    if r.GetPosition() == "N" {
        r.SetYPosition(r.GetYPosition() + 1)
    } else if r.GetPosition() == "E" {
        r.SetXPosition(r.GetXPosition() + 1)
    } else if r.GetPosition() == "W" {
        r.SetXPosition(r.GetXPosition() - 1)
    } else if r.GetPosition() =="S" {
        r.SetYPosition(r.GetYPosition() - 1)
    }
}