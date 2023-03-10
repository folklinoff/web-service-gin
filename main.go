package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postGames(c *gin.Context) {
	var newGame Game

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

var publishers []Publisher = []Publisher{
	{Name: "Valve", Year: 1996},
	{Name: "Mojang", Year: 2009},
	{Name: "Bethesda Game Studios", Year: 2001},
}

var games []Game = []Game{
	{Name: "CS:GO", Genre: "FPS", Year: 2001, Publisher: &publishers[0]},
	{Name: "Minecraft", Genre: "Sandbox", Year: 2009, Publisher: &publishers[1]},
	{Name: "Skyrim", Genre: "Roleplay Fighting", Year: 2011, Publisher: &publishers[2]},
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.POST("/games", postGames)

	router.Run("localhost:8080")
}
