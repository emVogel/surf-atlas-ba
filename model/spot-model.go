package model

import (
	"server-app/utils"
)

/**
* the response returned from the db
 */
type Response struct{
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
	Geom string`json:"geom"`;
}

/**
* definition for a spot
*/
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
	AllSpots() ([]Response, *utils.HttpError) 
}