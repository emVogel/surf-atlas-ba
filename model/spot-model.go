package model

import "server-app/utils"

type GeoJsonSpots struct{
	Type string `json:"type"`
	Features []SpotFeature `json:"features"`
}

type SpotFeature struct {

	Type string
	Geometry 
	Properties Spot `json:"properties"`
}

type Geometry struct {
	Type string `json:"type"`
	CRS CoordinateType `json:"crs"`
	Coordinates []float64 `json:"coordinates"`
}

type CoordinateType struct{
	Type string `json:"type"`
	Properties map[string]string `json:"properties"`
}


type Spot struct{
	ID uint `json:"id"`;
	Name string `json:"name"`;
	AlternativeName string `json:"alternative_name"`;
	Province string `json:"province"`;
	Type string `json:"type"`;
	Tide string  `json:"tide"`;
	Wind string `json:"wind"`;
	Swell string `json:"swell"`;
	Bottom string `json:"bottom"`;
	Direction string `json:"direction"`
	Access string `json:"access"`;
	Location string `json:"location"`
	Description string `json:"description"`
	Crowd string `json:"crowd"`;
	BestSeason string`json:"best_season"`;
}

/**
* interface for the db queries
*/
type SpotDbModel interface {
	AllSpots() (*[]Spot, *utils.HttpError) 
}