package rover

type Rover struct {
    rowPosition int
    columnPosition int
    orientation string
}

func (r *Rover) setXPosition(x int) {
 	r.rowPosition = x
}
func (r Rover) getXPosition() int {
 	return r.rowPosition
}

func (r *Rover) setYPosition(y int) {
 	r.columnPosition = y
}
func (r Rover) getYPosition() int {
 	return r.columnPosition
}

func (r *Rover) setPosition(x string) {
 	r.orientation = x
 }
func (r Rover) getPosition() string {
 	return r.orientation
}

func (r *Rover) processCommands(commands string)  {
    for _, command := range commands {
     if(string(command) == "L") {
        r.turnRoverLeft();
     } else if(string(command) == "R") {
        r.turnRoverRight();
     } else if(string(command) == "M") {
        r.moveRover();
     }
    }
}

func (r *Rover) turnRoverLeft() {
    currentOrientation := r.getPosition();
    switch(currentOrientation) {
    case "N":
      r.setPosition("W")
    case "S":
      r.setPosition("E")
    case "E":
       r.setPosition("N")
    case "W":
       r.setPosition("S")
    }
}

func (r *Rover) turnRoverRight() {
    currentOrientation := r.getPosition();
    switch(currentOrientation) {
    case "N":
      r.setPosition("E")
    case "S":
      r.setPosition("W")
    case "E":
       r.setPosition("S")
    case "W":
       r.setPosition("N")
    }
}

func (r *Rover) moveRover() {
    if(r.getPosition() == "N") {
        r.setYPosition(r.getYPosition() + 1);
    } else if(r.getPosition() == "E") {
        r.setXPosition(r.getXPosition() + 1);
    } else if(r.getPosition() == "W") {
        r.setXPosition(r.getXPosition() - 1);
    } else if(r.getPosition() =="S") {
        r.setYPosition(r.getYPosition() - 1);
    }
}