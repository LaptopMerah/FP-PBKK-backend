package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

type (
	ParticipantRepository interface {
		Create(ctx context.Context, participant entity.Participant) (entity.Participant, error)
		FindAll(ctx context.Context) ([]entity.Participant, error)
		FindByID(ctx context.Context, id uint) (entity.Participant, error)
		Update(ctx context.Context, id uint, participant entity.Participant) (entity.Participant, error)
		Delete(ctx context.Context, id uint) error
	}

	participantRepository struct {
		db *gorm.DB
	}
)

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{
		db: db,
	}
}


func (r *participantRepository) Create(ctx context.Context, participant entity.Participant) (entity.Participant, error) {
	err := r.db.Create(&participant).Error
	if err != nil {
		return entity.Participant{}, err
	}

	return participant, nil
}

func (r *participantRepository) FindAll(ctx context.Context) ([]entity.Participant, error) {
	var participants []entity.Participant
	err := r.db.Find(&participants).Error
	if err != nil {
		return []entity.Participant{}, err
	}

	return participants, nil
}

func (r *participantRepository) FindByID(ctx context.Context, id uint) (entity.Participant, error) {
	var participant entity.Participant
	err := r.db.Where("id = ?", id).First(&participant).Error
	if err != nil {
		return entity.Participant{}, err
	}

	return participant, nil
}

func (r *participantRepository) Update(ctx context.Context, id uint, participant entity.Participant) (entity.Participant, error) {
    err := r.db.Model(&entity.Participant{}).Where("id = ?", id).Updates(&participant).Error
    if err != nil {
        return entity.Participant{}, err
    }

	var updatedParticipant entity.Participant
	r.db.First(&updatedParticipant, id)
	return updatedParticipant, nil
}



func (r *participantRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).First(&entity.Participant{}).Error
	if err != nil {
		return err
	}
	r.db.Where("id = ?", id).Delete(&entity.Participant{})

	return nil
}
