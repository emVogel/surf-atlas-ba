package geojson

import (
	"encoding/json"
	"server-app/model"
	"server-app/utils"
)

type ResponseExecuter struct{
	resp model.Response
}

/**
* struct to handle the response for building the geojson 
*/
func newResonseExecuter(response model.Response) ResponseExecuter{
	return ResponseExecuter{
		resp: response,
	}
}

/**
* returns the geometry as astruct from the response
*/
func(executer *ResponseExecuter) GetGeometry() Geometry {
	// to convert the Geom filed into a struct, its marshalled by byte to later unmarshall it into Geom struct
	b := []byte(executer.resp.Geom)

	var geom Geometry

	json.Unmarshal(b, &geom)

	return geom
}

/**
* extracts the spot from the response 
*/
func(executer *ResponseExecuter) GetSpot() (model.Spot, utils.HttpError) {
	var spot model.Spot
	bi, mErr:= json.Marshal(&executer.resp)
	if(mErr != nil){
		return model.Spot{}, utils.NewHttpError(500, "server-error")
	}
	uErr := json.Unmarshal(bi, &spot)

	if (uErr !=nil) {
		return model.Spot{}, utils.NewHttpError(500, "server-error")
	}
	return spot, utils.HttpError{}
}

/**
* geojson collection containing the all features
*/
type GeoJsonFeatureCollection struct{
	Type string`json:"type"`
	Features []Feature`json:"features"`
}

/**
* GeoFeature containing geometry and properties
*/
type Feature struct {
	Type string`json:"type"`
	Geometry Geometry`json:"geometry"`
	Properties model.Spot`json:"properties"`
}


type Geometry struct {
	Type string `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}


func NewGeojsonFeatureCollection(features []Feature ) GeoJsonFeatureCollection{
	return GeoJsonFeatureCollection{
		Type: "FeatureCollection",
		Features: features,
	}
}

/**
* returns a new feature
*/
func NewFeature(geometry Geometry, properties  model.Spot) Feature {
	return Feature{
		Type: "Feature",
		Geometry: geometry,
		Properties: properties,
	}
}

/**
* build the geojson collection from the response
*/
func BuildGeojsonCollection(response []model.Response) (GeoJsonFeatureCollection, utils.HttpError) {
	var features []Feature 
	for index:= range response {
		resp := response[index]
	
		feature, err := GetFeatureFromResponse(resp)
		if(err != utils.HttpError{}) {
		return	GeoJsonFeatureCollection{}, err
		}

		features = append(features, feature)
	}

	return NewGeojsonFeatureCollection(features), utils.HttpError{}
}

func GetFeatureFromResponse(response model.Response) (Feature, utils.HttpError){
	exeCuter := newResonseExecuter(response);
	geom :=  exeCuter.GetGeometry()
	spot, err := exeCuter.GetSpot()
	
	if (err != utils.HttpError{}) {
		return Feature{}, err
	}
	 return NewFeature(geom, spot), utils.HttpError{}
}