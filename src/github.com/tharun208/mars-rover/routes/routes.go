package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tharun208/mars-rover/pkg/roverutil"
	"net/http"
)

var rover roverutil.Rover
var RoverArray = make([]roverutil.Rover,100)
func InitiateRouter(c *gin.Context) {
	// var jsonObj roverutil.Rover
	var jsonObj map[string]int
	c.BindJSON(&jsonObj)
	rover = roverutil.NewRover()
	if jsonObj["x"] > 0 && jsonObj["y"] > 0 {
		rover.SetXPosition(jsonObj["x"])
		rover.SetYPosition(jsonObj["y"])
		c.JSON(http.StatusOK,
			gin.H{
				"orientation": rover.GetPosition(),
				"x": rover.GetXPosition(),
				"y": rover.GetYPosition(),
			})
	} else {
		c.JSON(http.StatusNotFound,
			gin.H{
				"error" : "Please Provide Valid X and Y Points",
			})
	}

}

func ExecuteCommand(c *gin.Context) {
	var jsonObj map[string]string
	c.BindJSON(&jsonObj)
	if jsonObj["message"] != "" {
		rover.ProcessCommands(jsonObj["message"])
		c.JSON(http.StatusOK,
			gin.H{
				"orientation": rover.GetPosition(),
				"x": rover.GetXPosition(),
				"y": rover.GetYPosition(),
			})
	} else {
		c.JSON(http.StatusNotFound,
			gin.H{
				"error" : "Please Provide Valid Instructions",
			})
	}
}

func TurnRoverRight(c *gin.Context) {
	rover.TurnRoverRight()
	c.JSON(http.StatusOK,
		gin.H{
			"orientation": rover.GetPosition(),
			"x": rover.GetXPosition(),
			"y": rover.GetYPosition(),
		})
}

func TurnRoverLeft(c *gin.Context) {
	rover.TurnRoverLeft()
	c.JSON(http.StatusOK,
		gin.H{
			"orientation": rover.GetPosition(),
			"x": rover.GetXPosition(),
			"y": rover.GetYPosition(),
		})
}

func MoveRover(c *gin.Context) {
	rover.MoveRover()
	c.JSON(http.StatusOK,
		gin.H{
			"orientation": rover.GetPosition(),
			"x": rover.GetXPosition(),
			"y": rover.GetYPosition(),
		})
}

func GetRoverPosition(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"orientation": rover.GetPosition(),
			"x": rover.GetXPosition(),
			"y": rover.GetYPosition(),
		})
}