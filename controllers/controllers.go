package controllers

import (
	"net/http"
	"server-app/geojson"
	"server-app/model"
	"server-app/utils"

	"strconv"

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
		status := model.HttpResponseStatus{Status: dbErr.Status, Message: dbErr.Err};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	geojson, geoErr := geojson.BuildGeojsonCollection(resp)

	if (geoErr != utils.HttpError{}) {
		status := model.HttpResponseStatus{Status: geoErr.Status, Message: geoErr.Err};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	status := model.HttpResponseStatus{Status: http.StatusOK, Message: "ok"}
	c.JSON(http.StatusOK, gin.H{ "response_status": status, "data": geojson})
}

/**
* controller to return all spots to the clients
*/
func(ctrl *BaseController) GetAllSpots(c *gin.Context) {
	resp, err := ctrl.dbSpotModel.AllSpots()
	
	
	if (err != utils.HttpError{}) {
		status := model.HttpResponseStatus{Status: err.Status, Message: err.Err};

		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	geojson, err := geojson.BuildGeojsonCollection(resp)

	 status := model.HttpResponseStatus{Status: http.StatusOK, Message: "ok"}
	c.JSON(http.StatusOK, gin.H{ "response_status": status, "data": geojson})
}

func (ctrl * BaseController) GetSpotById(c *gin.Context) {
	id, isValid := c.GetQuery("id")

	if(!isValid) {
		status := model.HttpResponseStatus{Status: 403, Message: "no id provided"};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	intId, err := 	strconv.Atoi(id)

	if (err != nil) {
		status := model.HttpResponseStatus{Status: 500, Message: "server error"};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	resp, dbErr := ctrl.dbSpotModel.GetSpotById(intId)

	if (dbErr != utils.HttpError{}) {
		status := model.HttpResponseStatus{Status: dbErr.Status, Message: dbErr.Err};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}


	geojson, geoErr := geojson.BuildGeojsonCollection(resp)

	if (geoErr != utils.HttpError{}) {
		status := model.HttpResponseStatus{Status: geoErr.Status, Message: geoErr.Err};
		c.JSON(http.StatusBadRequest, gin.H{ "response_status": status})
		return
	}

	status := model.HttpResponseStatus{Status: http.StatusOK, Message: "ok"}
	c.JSON(http.StatusOK, gin.H{ "response_status": status, "data": geojson})


	
}