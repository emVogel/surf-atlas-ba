package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"server-app/utils"

	"github.com/lib/pq"
)
type Swell struct {
	Direction string`json:"direction"`
	Size string`json:"size"`
}

type BestConditions struct {
	Swell Swell`json:"swell"`
	Wind string`json:"wind"`
}

/**
* the response returned from the db
 */
type Response struct{
	ID uint `json:"id"`;
	Name string `json:"name"`;
	AlternativeName string `json:"alternative_name"`;
	Province string `json:"province"`;
	Type string `json:"type"`;
	Tide string  `json:"tide"`;	Wind pq.StringArray `gorm:"type:text[]" json:"wind"`;
	Swell string `json:"swell"`;
	Bottom string `json:"bottom"`;
	Direction string `json:"direction"`
	Access string `json:"access"`;
	Location string `json:"location"`
	Description string `json:"description"`
	Crowd string `json:"crowd"`;
	BestSeason string`json:"best_season"`;
	Geom string`json:"geom"`;
	BestConditions json.RawMessage`json:"best_conditions"`;
}

/**
* definition for a spot
*/
type RawSpot struct{
	ID uint `json:"id"`;
	Name string `json:"name"`;
	AlternativeName string `json:"alternative_name"`;
	Province string `json:"province"`;
	Type string `json:"type"`;
	Tide string  `json:"tide"`;
	Wind []string `json:"wind"`;
	Bottom string `json:"bottom"`;
	Direction string `json:"direction"`
	Access string `json:"access"`;
	Location string `json:"location"`
	Description string `json:"description"`
	Crowd string `json:"crowd"`;
	BestSeason string`json:"best_season"`;
	BestConditions BestConditions `json:"best_conditions"`;
}

type Spot struct{
	RawSpot
	Swell []string `json:"swell"`;

}

func(spot *Spot) Validador(key string, value string) (bool, utils.HttpError) {
	structureVal:= reflect.ValueOf(spot).Elem()
	fmt.Println("a struct value of Elem", structureVal)
	isSpotStructProperty := utils.ValidateFilterKey(key, structureVal)
	
	 if !isSpotStructProperty {
			return false, utils.NewHttpError(403, "Not allowed")
	 }
	 isValidValue := utils.ValidateFilterValue(value)
	 if !isValidValue {
		return false, utils.NewHttpError(403, "Not allowed")
	 }
	 return true, utils.HttpError{}
	}

	type HttpResponseStatus struct {
		Status int `json:"status"`;
		Message string `json:"message"`;
	}

/**
* interface for the db queries
*/
type SpotDbModel interface {
	AllSpots() ([]Response, utils.HttpError) 
	FilterSpotsByProperties(map[string][]string) ([]Response, utils.HttpError) 
	GetSpotById(id int) ([]Response, utils.HttpError)
}