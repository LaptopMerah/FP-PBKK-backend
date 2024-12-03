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
	EventController interface {
		CreateEvent(ctx *gin.Context)
		GetAllEvents(ctx *gin.Context)
		GetEventByID(ctx *gin.Context)
		UpdateEvent(ctx *gin.Context)
		DeleteEvent(ctx *gin.Context)
	}

	eventController struct {
		eventService service.EventService
	}
)

func NewEventController(es service.EventService) EventController {
	return &eventController{
		eventService: es,
	}
}

func (c *eventController) CreateEvent(ctx *gin.Context) {
	var req dto.EventCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Failed to bind request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.eventService.CreateEvent(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Event created successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *eventController) GetAllEvents(ctx *gin.Context) {
	result, err := c.eventService.GetAllEvents(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get events", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Events retrieved successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *eventController) GetEventByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid event ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.eventService.GetEventByID(ctx.Request.Context(), uint(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Event retrieved successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *eventController) UpdateEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid event ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EventUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Failed to bind request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.eventService.UpdateEvent(ctx.Request.Context(), uint(id), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Event updated successfully", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *eventController) DeleteEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid event ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.eventService.DeleteEvent(ctx.Request.Context(), uint(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete event", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Event deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}