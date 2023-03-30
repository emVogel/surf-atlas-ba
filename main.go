package main

import (
	"fmt"
	"net/http"
	"server-app/controllers"
	"server-app/database"
	"server-app/repository"

	"github.com/gin-gonic/gin"
)




func main() {
	r :=gin.Default()

	db := database.InitDataBase()

	repo:= repository.NewSpotRepository(db)

	controllers:= controllers.NewBaseController(repo)

	r.GET("/spots", controllers.GetAllSpots)

	r.GET("/filterspots", controllers.GetSpotsByFilter)

	r.GET("/", func(c *gin.Context) {
		paramPairs := c.Request.URL.Query()
		for key, values := range paramPairs {
			fmt.Printf("key = %v, value(s) = %v\n", key, values[0])
		}

		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s %s", firstname, lastname, paramPairs)
	})

	r.Run("0.0.0.0:4040");
}