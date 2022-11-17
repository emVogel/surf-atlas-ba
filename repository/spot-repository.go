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
func (spotRepo *SpotRepository) AllSpots() (*[]model.Spot, *utils.HttpError) {
	var spots []model.Spot

	spotRepo.db.Find(&spots)

	if len(spots) == 0 {
		return nil, utils.NewHttpError(400, "Bad Request")
	}
	return &spots, &utils.HttpError{}
}