package handler

import (
	"Todo_list/internal/model"
	"Todo_list/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTask(c *gin.Context) {
	var input model.Task
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Task.Create(input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.ErrUnvalidatedData {
			statusCode = http.StatusBadRequest
		}

		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSONP(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllTasks(c *gin.Context) {}

func (h *Handler) getCompletedTasks(c *gin.Context) {}

func (h *Handler) getUncompletedTasks(c *gin.Context) {}

func (h *Handler) updateTask(c *gin.Context) {}

func (h *Handler) deleteTask(c *gin.Context) {}
