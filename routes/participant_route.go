package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func Participant(route *gin.Engine, participantController controller.ParticipantController) {
	routes := route.Group("/api/participant")
	{
		routes.POST("/", participantController.CreateParticipant)
		routes.GET("/", participantController.GetAllParticipants)
		routes.GET("/:id", participantController.GetParticipantByID)
		routes.PATCH("/:id", participantController.UpdateParticipant)
		routes.DELETE("/:id", participantController.DeleteParticipant)
	}
}
