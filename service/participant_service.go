package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
	"gorm.io/gorm"
)

type (
	ParticipantService interface {
		CreateParticipant(ctx context.Context, req dto.ParticipantCreateRequest) (dto.ParticipantResponse, error)
		GetAllParticipants(ctx context.Context) ([]dto.ParticipantResponse, error)
		GetParticipantByID(ctx context.Context, id uint) (dto.ParticipantResponse, error)
		UpdateParticipant(ctx context.Context, id uint, req dto.ParticipantUpdateRequest) (dto.ParticipantResponse, error)
		DeleteParticipant(ctx context.Context, id uint) error
	}

	participantService struct {
		participantRepo repository.ParticipantRepository
	}
)

func NewParticipantService(er repository.ParticipantRepository) ParticipantService {
	return &participantService{
		participantRepo: er,
	}
}

func (s *participantService) CreateParticipant (ctx context.Context, req dto.ParticipantCreateRequest) (dto.ParticipantResponse , error) {
	participant := entity.Participant{
		EventID:	req.EventID,
		Name:		req.Name,
		Email:		req.Email,
	}

	createdParticipant, err := s.participantRepo.Create(ctx, participant)
	if err != nil {
		return dto.ParticipantResponse{}, err
	}

	return dto.ParticipantResponse{
		ID:			createdParticipant.ID,
		EventID:	createdParticipant.EventID,
		Name:		createdParticipant.Name,
		Email:		createdParticipant.Email,
	}, nil
}

func (s *participantService) GetAllParticipants(ctx context.Context) ([]dto.ParticipantResponse, error) {
	participants, err := s.participantRepo.FindAll(ctx)
	if err != nil {
		return []dto.ParticipantResponse{}, err
	}

	var participantResponses []dto.ParticipantResponse
	for _, participant := range participants {
		participantResponses = append(participantResponses, dto.ParticipantResponse{
			ID:			participant.ID,
			EventID:	participant.EventID,
			Name:		participant.Name,
			Email:		participant.Email,
		})
	}

	return participantResponses, nil
}

func (s *participantService) GetParticipantByID(ctx context.Context, id uint) (dto.ParticipantResponse, error) {
	participant, err := s.participantRepo.FindByID(ctx, id)
	if err != nil {
		return dto.ParticipantResponse{}, err
	}

	return dto.ParticipantResponse{
		ID:			participant.ID,
		EventID:	participant.EventID,
		Name:		participant.Name,
		Email:		participant.Email,
	}, nil
}

func (s *participantService) UpdateParticipant(ctx context.Context, id uint, req dto.ParticipantUpdateRequest) (dto.ParticipantResponse, error) {
    existingParticipant, err := s.participantRepo.FindByID(ctx, id)
    if err != nil {
        return dto.ParticipantResponse{}, err
    }

    participant := entity.Participant{
        Model:   gorm.Model{ID: existingParticipant.ID, CreatedAt: existingParticipant.CreatedAt},
        EventID: req.EventID,
        Name:    req.Name,
        Email:   req.Email,
    }

    updatedParticipant, err := s.participantRepo.Update(ctx, id, participant)
    if err != nil {
        return dto.ParticipantResponse{}, err
    }

    return dto.ParticipantResponse{
        ID:      updatedParticipant.ID,
        EventID: updatedParticipant.EventID,
        Name:    updatedParticipant.Name,
        Email:   updatedParticipant.Email,
    }, nil
}


func (s *participantService) DeleteParticipant(ctx context.Context, id uint) error {
	err := s.participantRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}