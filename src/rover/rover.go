package rover

type Rover struct {
    rowPosition int
    columnPosition int
    orientation string
}

func (r *Rover) SetXPosition(x int) {
 	r.rowPosition = x
}
func (r Rover) getXPosition() int {
 	return r.rowPosition
}

func (r *Rover) SetYPosition(y int) {
 	r.columnPosition = y
}
func (r Rover) getYPosition() int {
 	return r.columnPosition
}

func (r *Rover) SetPosition(x string) {
 	r.orientation = x
 }
func (r Rover) getPosition() string {
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
    currentOrientation := r.getPosition();
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
    currentOrientation := r.getPosition()
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
    if r.getPosition() == "N" {
        r.SetYPosition(r.getYPosition() + 1)
    } else if r.getPosition() == "E" {
        r.SetXPosition(r.getXPosition() + 1)
    } else if r.getPosition() == "W" {
        r.SetXPosition(r.getXPosition() - 1)
    } else if r.getPosition() =="S" {
        r.SetYPosition(r.getYPosition() - 1)
    }
}