package controller

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/Caknoooo/go-gin-clean-starter/utils"
	"github.com/gin-gonic/gin"
)

type (
	ParticipantController interface {
		CreateParticipant(ctx *gin.Context)
		GetAllParticipants(ctx *gin.Context)
		GetParticipantByID(ctx *gin.Context)
		UpdateParticipant(ctx *gin.Context)
		DeleteParticipant(ctx *gin.Context)
	}

	participantController struct {
		participantService service.ParticipantService
	}
)

func NewParticipantController(es service.ParticipantService) ParticipantController {
	return &participantController{
		participantService: es,
	}
}

func (c *participantController) CreateParticipant(ctx *gin.Context) {
	var req dto.ParticipantCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Failed to bind request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.participantService.CreateParticipant(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Participant created successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *participantController) GetAllParticipants(ctx *gin.Context) {
	result, err := c.participantService.GetAllParticipants(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get events", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Participants retrieved successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *participantController) GetParticipantByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid event ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.participantService.GetParticipantByID(ctx.Request.Context(), uint(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Participant retrieved successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *participantController) UpdateParticipant(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid event ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.ParticipantUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Failed to bind request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.participantService.UpdateParticipant(ctx.Request.Context(), uint(id), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update Participant", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Participant updated successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *participantController) DeleteParticipant(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid Participant ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.participantService.DeleteParticipant(ctx.Request.Context(), uint(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete Participant", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Participant deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}