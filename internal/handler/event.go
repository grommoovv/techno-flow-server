package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"strconv"
)

func (h *Handler) CreateEvent(c *gin.Context) {
	var eventDto entities.EventCreateDto

	if err := c.BindJSON(&eventDto); err != nil {
		ResponseError(c, "failed to bind event dto", err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.services.Event.CreateEvent(eventDto)

	if err != nil {
		ResponseError(c, "failed to create event", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "event created successfully", map[string]interface{}{"event_id": id})
}

func (h *Handler) GetAllEvents(c *gin.Context) {
	events, err := h.services.Event.GetAllEvents()

	if err != nil {
		ResponseError(c, "failed to fetch events", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "events fetched successfully", events)
}

func (h *Handler) GetEventById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	event, err := h.services.Event.GetEventById(id)

	if err != nil {
		ResponseError(c, "failed to fetch event", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "event fetched successfully", event)
}

func (h *Handler) GetEventsByUserId(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.services.Event.GetEventsByUserId(id)

	if err != nil {
		ResponseError(c, "failed to fetch events", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "events fetched successfully", events)
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.Event.DeleteEvent(id)
	if err != nil {
		ResponseError(c, "failed to delete event", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "event deleted successfully", map[string]interface{}{"event_id": id})
}

func (h *Handler) UpdateEvent(c *gin.Context) {}
