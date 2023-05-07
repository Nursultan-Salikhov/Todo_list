package handler

import (
	"Todo_list/internal/model"
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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllTasks(c *gin.Context) {
	tasks, err := h.service.Task.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getCompletedTasks(c *gin.Context) {}

func (h *Handler) getUncompletedTasks(c *gin.Context) {}

func (h *Handler) updateTask(c *gin.Context) {}

func (h *Handler) deleteTask(c *gin.Context) {}
