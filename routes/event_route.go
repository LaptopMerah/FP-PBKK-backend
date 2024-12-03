package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func Event(route *gin.Engine, eventController controller.EventController) {
	routes := route.Group("/api/event")
	{
		routes.POST("/", eventController.CreateEvent)
		routes.GET("/", eventController.GetAllEvents)
		routes.GET("/:id", eventController.GetEventByID)
		routes.PATCH("/:id", eventController.UpdateEvent)
		routes.DELETE("/:id", eventController.DeleteEvent)
	}
}
