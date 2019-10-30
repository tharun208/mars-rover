package main
import (
    "fmt"
    "rover"
    )
func main() {
    var commands string
    fmt.Println("Enter the No Of Rovers");
    var noOfRovers int;
    fmt.Scan(&noOfRovers);
    for noOfRovers != 0 {
    var rowPosition, columnPosition int;
    var initialOrientation string;
    fmt.Println("Enter the Initial X and Y Position")
    fmt.Scan(&rowPosition, &columnPosition);
    fmt.Println("Enter the Orientation");
    fmt.Scanln(&initialOrientation);
    rover := rover.Rover{}
    rover.setXPosition(rowPosition);
    rover.setYPosition(columnPosition);
    rover.setPosition(initialOrientation);
    fmt.Println("Enter the Commands the rover needs to follow:");
    fmt.Scanln(&commands);
    rover.processCommands(commands);
    fmt.Println("The Rover ",rover);
    noOfRovers--;
    }
}