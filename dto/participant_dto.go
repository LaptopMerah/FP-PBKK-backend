package dto

const (
	MESSAGE_FAILED_GET_DATA_PARTICIPANT = "Failed to get data from request"
	MESSAGE_SUCCESS_CREATE_PARTICIPANT  = "Participant created successfully"
	MESSAGE_FAILED_CREATE_PARTICIPANT   = "Failed to create participant"
	MESSAGE_SUCCESS_GET_PARTICIPANT     = "Participant retrieved successfully"
	MESSAGE_FAILED_GET_PARTICIPANT      = "Failed to retrieve participant"
	MESSAGE_SUCCESS_UPDATE_PARTICIPANT  = "Participant updated successfully"
	MESSAGE_FAILED_UPDATE_PARTICIPANT   = "Failed to update participant"
	MESSAGE_SUCCESS_DELETE_PARTICIPANT  = "Participant deleted successfully"
	MESSAGE_FAILED_DELETE_PARTICIPANT   = "Failed to delete participant"
)

type ParticipantCreateRequest struct {
	EventID uint   `json:"event_id" form:"event_id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required"`
}

type ParticipantResponse struct {
	ID        uint   `json:"id"`
	EventID   uint   `json:"event_id"`
	EventName string `json:"event_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type ParticipantUpdateRequest struct {
	EventID uint   `json:"event_id" form:"event_id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required"`
}

type ParticipantDeleteRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
