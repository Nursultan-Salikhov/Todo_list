package handler

import (
	"Todo_list/internal/model"
	"Todo_list/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) getCompletedTasks(c *gin.Context) {
	tasks, err := h.service.Task.GetCompleted()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getUncompletedTasks(c *gin.Context) {
	tasks, err := h.service.Task.GetUncompleted()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) updateTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input model.UpdateTask
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Task.Update(taskId, input)
	if err != nil {
		statusCode := http.StatusInternalServerError

		switch err {
		case service.ErrUnvalidatedData:
			statusCode = http.StatusBadRequest

		case service.ErrNotFound:
			statusCode = http.StatusNotFound
		}

		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Update success"})
}

func (h *Handler) deleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.service.Task.Delete(taskId)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.ErrNotFound {
			statusCode = http.StatusNotFound
		}

		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Delete success"})
}
