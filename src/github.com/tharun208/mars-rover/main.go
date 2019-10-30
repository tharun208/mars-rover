package main

import (
    "fmt"
    "github.com/tharun208/mars-rover/roverutil"
)

func main() {
    var commands string
    fmt.Println("Enter the No Of Rovers")
    var noOfRovers int
    fmt.Scan(&noOfRovers)
    for noOfRovers != 0 {
        var rowPosition, columnPosition int
        var initialOrientation string
        fmt.Println("Enter the Initial X and Y Position")
        fmt.Scan(&rowPosition, &columnPosition)
        fmt.Println("Enter the Orientation")
        fmt.Scanln(&initialOrientation)
        rover := roverutil.Rover{}
        rover.SetXPosition(rowPosition)
        rover.SetYPosition(columnPosition)
        rover.SetPosition(initialOrientation)
        fmt.Println("Enter the Commands the rover needs to follow:")
        fmt.Scanln(&commands)
        rover.ProcessCommands(commands)
        fmt.Println("The Rover is at X ",rover.GetXPosition()," Y ",rover.GetYPosition()," facing ",rover.GetPosition())
        noOfRovers--
    }
}
