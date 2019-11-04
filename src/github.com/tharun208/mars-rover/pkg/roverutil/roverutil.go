package roverutil

type Rover struct {
    RowPosition    int
    ColumnPosition int
    Orientation    string
}

func (r *Rover) SetXPosition(x int) {
 	r.RowPosition = x
}
func (r Rover) GetXPosition() int {
 	return r.RowPosition
}

func (r *Rover) SetYPosition(y int) {
 	r.ColumnPosition = y
}
func (r Rover) GetYPosition() int {
 	return r.ColumnPosition
}

func (r *Rover) SetPosition(x string) {
 	r.Orientation = x
 }
func (r Rover) GetPosition() string {
 	return r.Orientation
}

func (r *Rover) ProcessCommands(commands string)  {
    for _, command := range commands {
     if string(command) == "L" {
        r.TurnRoverLeft()
     } else if string(command) == "R" {
        r.TurnRoverRight()
     } else if string(command) == "M" {
        r.moveRover()
     }
    }
}

func (r *Rover) TurnRoverLeft() {
    currentOrientation := r.GetPosition()
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

func (r *Rover) TurnRoverRight() {
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