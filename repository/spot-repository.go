package repository

import (
	"fmt"
	"server-app/model"
	"server-app/utils"

	"gorm.io/gorm"
)

/**
* holds the db instance with the repo Queries
 */
type SpotRepository struct {
	db *gorm.DB
}

/**
* returns a new SpotRepository, which holds the db-instance and all spots
*/
func NewSpotRepository(db *gorm.DB) *SpotRepository {
	return &SpotRepository{
		db: db,
	}
}

func(spotRepo *SpotRepository) FilterSpotsByProperties(query map[string][]string)([]model.Response, utils.HttpError) {
	fmt.Print("query", query)
	var spots [] model.Response
	var spotStruct model.Spot
	var queryString []string
	for key, value:= range query {
		
		ok, error := spotStruct.Validador(key, value[0])
		if !ok{
			return nil, error
		}
		str := key +" LIKE " + "'"+value[0]+"'"
		queryString = append(queryString, str )
	}
	var sqlString string
	for index := range queryString {
		
		if (index == 0) {
			sqlString =  queryString[index] + " " + " AND "
		}
		if (len(queryString) -1 != index && index >0) {
			sqlString = sqlString + " " + queryString[index] + " AND "
		} 
		if (len(queryString) -1 == index) {
			sqlString = sqlString + " "  + queryString[index]
		}
		
	}

	sql := fmt.Sprintf("SELECT id, name, alternative_name, wind, swell, province, bottom, access, location, description, direction, crowd, best_season, type, tide,  ST_AsGeoJSON(ST_Transform(geom, 4326)) as geom from spots WHERE %s", sqlString)
	fmt.Print(sqlString)
	spotRepo.db.Raw(sql).Scan(&spots)
	return spots, utils.HttpError{}
}

/*
* get all spots from db
*/
func (spotRepo *SpotRepository) AllSpots() ([]model.Response, utils.HttpError) {
	var spots []model.Response

	spotRepo.db.Table("spots").Select("id, name, alternative_name, wind, swell, province, bottom, access, location, description,direction, crowd, best_season, type, tide, best_conditions, ST_AsGeoJSON(ST_Transform(geom, 4326)) as geom").Scan(&spots)

	if len(spots) == 0 {
		return nil, utils.NewHttpError(400, "Bad Request")
	}
	return spots, utils.HttpError{}
}
 
func(spotRepo *SpotRepository) GetSpotById(id int) ([]model.Response, utils.HttpError) {
	var spots []model.Response

	spotRepo.db.Table("spots").Select("id, name, alternative_name, wind, swell, province, bottom, access, location, description,direction, crowd, best_season, type, tide, best_conditions, ST_AsGeoJSON(ST_Transform(geom, 4326)) as geom").First(&spots, id)

	if len(spots) == 0 {
		return nil, utils.NewHttpError(400, "Bad Request, no spot with such id")
	}
	return spots, utils.HttpError{}

}