package main

import (
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

	r.Run("localhost:4040");
}