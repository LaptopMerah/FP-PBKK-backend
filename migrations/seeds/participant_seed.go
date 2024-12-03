package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func ListParticipantSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/participant.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listUser []entity.Participant
	if err := json.Unmarshal(jsonData, &listUser); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Participant{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Participant{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var participant entity.Participant
		err := db.Where(&entity.Participant{EventID: data.EventID}).First(&participant).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&participant, "event_id = ? AND email = ?", data.EventID, data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
