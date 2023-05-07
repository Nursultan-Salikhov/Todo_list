package handler

import (
	"Todo_list/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	task := router.Group("/task")
	{
		task.POST("/", h.createTask)
		task.GET("/", h.getAllTasks)
		task.PUT("/:id", h.updateTask)
		task.DELETE("/:id", h.deleteTask)

		completed := task.Group("completed")
		{
			completed.GET("/", h.getCompletedTasks)
		}

		uncompleted := task.Group("uncompleted")
		{
			uncompleted.GET("/", h.getUncompletedTasks)
		}
	}

	return router
}
