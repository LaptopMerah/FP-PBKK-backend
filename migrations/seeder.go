package migrations

import (
	"github.com/Caknoooo/go-gin-clean-starter/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListEventSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListParticipantSeeder(db); err != nil {
		return err
	}

	return nil
}
