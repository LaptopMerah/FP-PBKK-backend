package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

type (
	EventRepository interface {
		Create(ctx context.Context, event entity.Event) (entity.Event, error)
		FindAll(ctx context.Context) ([]entity.Event, error)
		FindByID(ctx context.Context, id uint) (entity.Event, error)
		Update(ctx context.Context, id uint, event entity.Event) (entity.Event, error)
		Delete(ctx context.Context, id uint) error
	}

	eventRepository struct {
		db *gorm.DB
	}
)

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) Create(ctx context.Context, event entity.Event) (entity.Event, error) {
	err := r.db.Create(&event).Error
	if err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) FindAll(ctx context.Context) ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Find(&events).Error
	if err != nil {
		return []entity.Event{}, err
	}

	return events, nil
}

func (r *eventRepository) FindByID(ctx context.Context, id uint) (entity.Event, error) {
	var event entity.Event
	err := r.db.Where("id = ?", id).First(&event).Error
	if err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) Update(ctx context.Context, id uint, event entity.Event) (entity.Event, error) {
	err := r.db.Model(&entity.Event{}).Where("id = ?", id).Updates(&event).Error
	if err != nil {
		return entity.Event{}, err
	}

	var updatedEvent entity.Event
	r.db.First(&updatedEvent, id)
	return updatedEvent, nil
}


func (r *eventRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).First(&entity.Event{}).Error
	if err != nil {
		return err
	}
	r.db.Where("id = ?", id).Delete(&entity.Event{})

	return nil
}