package controllers

import (
	"net/http"
	"server-app/model"
	"server-app/utils"

	"github.com/gin-gonic/gin"
)

/**
* the struct holds the syntax to access the dbModel for the controllers
 */
type BaseController struct {
	dbSpotModel model.SpotDbModel
}

/**
* the BaseController contains the controller with access to the dbModel
*/
func NewBaseController(dbSpotModel model.SpotDbModel) *BaseController {
	return &BaseController{
		dbSpotModel: dbSpotModel,
	}
}

/**
* controller to return all spots to the clients
*/
func(ctrl *BaseController) GetAllSpots(c *gin.Context) {
	
	spots, err := ctrl.dbSpotModel.AllSpots()

	if err != (&utils.HttpError{}) {
		c.JSON(http.StatusBadRequest, gin.H{ "status": err.Status, "message": err.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "message": "ok", "data": spots})
}
