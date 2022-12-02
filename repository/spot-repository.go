package repository

import (
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

/*
* get all spots from db
*/
func (spotRepo *SpotRepository) AllSpots() ([]model.Response, *utils.HttpError) {
	var spots []model.Response

	spotRepo.db.Raw("SELECT id, name, alternative_name, wind, swell, province, bottom, access, location, description, crowd, best_season, type, tide,  ST_AsGeoJSON(geom) as geom FROM spots").Scan(&spots)

	if len(spots) == 0 {
		return nil, utils.NewHttpError(400, "Bad Request")
	}
	return spots, &utils.HttpError{}
}