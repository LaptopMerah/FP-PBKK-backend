package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"gorm.io/gorm"
)

type (
	EventService interface {
		CreateEvent(ctx context.Context, req dto.EventCreateRequest) (dto.EventResponse, error)
		GetAllEvents(ctx context.Context) ([]dto.EventResponse, error)
		GetEventByID(ctx context.Context, id uint) (dto.EventResponse, error)
		UpdateEvent(ctx context.Context, id uint, req dto.EventUpdateRequest) (dto.EventResponse, error)
		DeleteEvent(ctx context.Context, id uint) error
	}

	eventService struct {
		eventRepo repository.EventRepository
	}
)

func NewEventService(er repository.EventRepository) EventService {
	return &eventService{
		eventRepo: er,
	}
}

func (s *eventService) CreateEvent(ctx context.Context, req dto.EventCreateRequest) (dto.EventResponse, error) {
	event := entity.Event{
		EventName:  req.EventName,
		Date: 		req.Date,
		Location:   req.Location,
		Details:    req.Details,
	}

	createdEvent, err := s.eventRepo.Create(ctx, event)
	if err != nil {
		return dto.EventResponse{}, err
	}

	return dto.EventResponse{
		ID:			createdEvent.ID,
		EventName:	createdEvent.EventName,
		Date:		createdEvent.Date,
		Location:	createdEvent.Location,
		Details:	createdEvent.Details,
	}, nil
}

func (s *eventService) GetAllEvents(ctx context.Context) ([]dto.EventResponse, error) {
	events, err := s.eventRepo.FindAll(ctx)
	if err != nil {
		return []dto.EventResponse{}, err
	}

	var eventResponses []dto.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, dto.EventResponse{
			ID:			event.ID,
			EventName:	event.EventName,
			Date:		event.Date,
			Location:	event.Location,
			Details:	event.Details,
		})
	}

	return eventResponses, nil
}

func (s *eventService) GetEventByID(ctx context.Context, id uint) (dto.EventResponse, error) {
	event, err := s.eventRepo.FindByID(ctx, id)
	if err != nil {
		return dto.EventResponse{}, err
	}

	return dto.EventResponse{
		ID:			event.ID,
		EventName:	event.EventName,
		Date:		event.Date,
		Location:	event.Location,
		Details:	event.Details,
	}, nil
}

func (s *eventService) UpdateEvent(ctx context.Context, id uint, req dto.EventUpdateRequest) (dto.EventResponse, error) {
	existingEvent, err := s.eventRepo.FindByID(ctx, id)
	if err != nil {
		return dto.EventResponse{}, err
	}

	event := entity.Event{
		Model:      gorm.Model{ID: existingEvent.ID, CreatedAt: existingEvent.CreatedAt},
		EventName:  req.EventName,
		Date:       req.Date,
		Location:   req.Location,
		Details:    req.Details,
	}

	updatedEvent, err := s.eventRepo.Update(ctx, id, event)
	if err != nil {
		return dto.EventResponse{}, err
	}

	return dto.EventResponse{
		ID:         updatedEvent.ID,
		EventName:  updatedEvent.EventName,
		Date:       updatedEvent.Date,
		Location:   updatedEvent.Location,
		Details:    updatedEvent.Details,
	}, nil
}


func (s *eventService) DeleteEvent(ctx context.Context, id uint) error {
	err := s.eventRepo.Delete(ctx, id)
	if err != nil {
		return  err
	}

	return nil

}
