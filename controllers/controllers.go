package controllers

import (
	"net/http"
	"server-app/geojson"
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

/*
* the controller for filterig spots by properties such as ?province=Costa do Morte
*/
func (ctrl *BaseController) GetSpotsByFilter(c *gin.Context) {

	query := c.Request.URL.Query()

	resp, dbErr := ctrl.dbSpotModel.FilterSpotsByProperties(query)

	if (dbErr != utils.HttpError{}) {
		c.JSON(http.StatusBadRequest, gin.H{ "status": dbErr.Status, "message": dbErr.Err})
		return
	}

	geojson, geoErr := geojson.BuildGeojsonCollection(resp)

	if (geoErr != utils.HttpError{}) {
		c.JSON(http.StatusBadRequest, gin.H{ "status": geoErr.Status, "message": geoErr.Err})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "message": "ok", "data": geojson})
}

/**
* controller to return all spots to the clients
*/
func(ctrl *BaseController) GetAllSpots(c *gin.Context) {
	
	resp, err := ctrl.dbSpotModel.AllSpots()
	
	geojson, err := geojson.BuildGeojsonCollection(resp)
	
	if (err != utils.HttpError{}) {
		c.JSON(http.StatusBadRequest, gin.H{ "status": err.Status, "message": err.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "message": "ok", "data": geojson})
}
