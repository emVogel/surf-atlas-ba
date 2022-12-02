package geojson

import (
	"encoding/json"
	"server-app/model"
	"server-app/utils"
)

type ResponseExecuter struct{
	resp model.Response
}

func newResonseExecuter(response model.Response) ResponseExecuter{
	return ResponseExecuter{
		resp: response,
	}
}

/**
* returns the geomtry as astruct from the response
*/
func(executer *ResponseExecuter) GetGeometry() Geometry {
	b := []byte(executer.resp.Geom)

	var geom Geometry

	json.Unmarshal(b, &geom)

	return geom
}

func(executer *ResponseExecuter) GetSpot() (model.Spot) {
	var spot model.Spot
	bi, mErr:= json.Marshal(&executer.resp)
	if(mErr != nil){
		return model.Spot{}
	}
	uErr := json.Unmarshal(bi, &spot)

	if (uErr !=nil) {
		return model.Spot{}
	}
	return spot
}

type GeoJsonFeatureCollection struct{
	Type string`json:"type"`
	Features []Feature`json:"features"`
}

type Feature struct {
	Type string`json:"Type"`
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

func NewFeature(geometry Geometry, properties  model.Spot) Feature {
	return Feature{
		Type: "Feature",
		Geometry: geometry,
		Properties: properties,
	}
}

func BuildGeojsonCollection(response []model.Response) (GeoJsonFeatureCollection, *utils.HttpError) {
	var features []Feature 
	for index:= range response {
		resp := response[index]
		exeCuter := newResonseExecuter(resp);
		geom :=  exeCuter.GetGeometry()
		spot := exeCuter.GetSpot()
		
		if (spot == model.Spot{}) {
			return GeoJsonFeatureCollection{}, utils.NewHttpError(500, "server error")
		}
		 feature := NewFeature(geom, spot)

		features = append(features, feature)
	}

	return NewGeojsonFeatureCollection(features), &utils.HttpError{}
}