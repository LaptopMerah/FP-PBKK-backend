package dto

const (
	MESSAGE_FAILED_GET_DATA_EVENT = "Failed to get data from request"
	MESSAGE_SUCCESS_CREATE_EVENT  = "Event created successfully"
	MESSAGE_FAILED_CREATE_EVENT   = "Failed to create event"
	MESSAGE_SUCCESS_GET_EVENT     = "Event retrieved successfully"
	MESSAGE_FAILED_GET_EVENT      = "Failed to retrieve event"
	MESSAGE_SUCCESS_UPDATE_EVENT  = "Event updated successfully"
	MESSAGE_FAILED_UPDATE_EVENT   = "Failed to update event"
	MESSAGE_SUCCESS_DELETE_EVENT  = "Event deleted successfully"
	MESSAGE_FAILED_DELETE_EVENT   = "Failed to delete event"
)

type EventCreateRequest struct {
	EventName string `json:"event_name" form:"event_name" binding:"required"`
	Date      string `json:"date" form:"date" binding:"required"`
	Location  string `json:"location" form:"location" binding:"required"`
	Details   string `json:"details" form:"details" binding:"required"`
}

type EventResponse struct {
	ID        uint   `json:"id"`
	EventName string `json:"event_name"`
	Date      string `json:"date"`
	Location  string `json:"location"`
	Details   string `json:"details"`
}

type EventUpdateRequest struct {
	EventName string `json:"event_name" form:"event_name" binding:"required"`
	Date      string `json:"date" form:"date" binding:"required"`
	Location  string `json:"location" form:"location" binding:"required"`
	Details   string `json:"details" form:"details" binding:"required"`
}

type EventDeleteRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
