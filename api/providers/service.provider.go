package providers

import (
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/util"
)

func GetAllServices(dest *[]models.Service) error {
	// open database connection
	db, err := util.GetDBInstance()
	if err != nil {
		return err
	}

	// get services from database
	if err := db.Find(&dest).Error; err != nil {
		return err
	}

	return nil
}