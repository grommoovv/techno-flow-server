package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"strconv"
)

func (h *Handler) CreateEquipment(c *gin.Context) {
	var equipmentDto entities.EquipmentCreateDto

	if err := c.BindJSON(&equipmentDto); err != nil {
		ResponseError(c, "failed to bind equipment dto", err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Equipment.CreateEquipment(equipmentDto)

	if err != nil {
		ResponseError(c, "failed to create equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user created successfully", map[string]interface{}{"equipment_id": id})
}

func (h *Handler) GetAllEquipment(c *gin.Context) {
	equipment, err := h.services.Equipment.GetAllEquipment()

	if err != nil {
		ResponseError(c, "failed to fetch equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "equipment fetched successfully", equipment)
}

func (h *Handler) GetAvailableEquipmentByDate(c *gin.Context) {

	var dto entities.GetAvailableEquipmentByDateDto

	if err := c.BindJSON(&dto); err != nil {
		ResponseError(c, "failed to bind equipment dto", err.Error(), http.StatusBadRequest)
		return
	}

	equipment, err := h.services.Equipment.GetAvailableEquipmentByDate(dto)

	if err != nil {
		ResponseError(c, "failed to fetch available equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	if equipment == nil {
		ResponseError(c, "failed to fetch available equipment", "no available equipment", http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "available equipment fetched successfully", equipment)
}

func (h *Handler) GetEquipmentById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	equipment, err := h.services.Equipment.GetEquipmentById(id)

	if err != nil {
		ResponseError(c, "failed to fetch equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "equipment fetched successfully", equipment)
}

func (h *Handler) GetEquipmentUsageHistoryById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.services.Equipment.GetEquipmentUsageHistoryById(id)

	if err != nil {
		ResponseError(c, "failed to fetch equipment reservation dates", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "equipment usage history fetched successfully", history)
}

func (h *Handler) DeleteEquipment(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.services.Equipment.DeleteEquipment(id)
	if err != nil {
		ResponseError(c, "failed to delete equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "equipment deleted successfully", map[string]interface{}{"equipment_id": id})
}

func (h *Handler) UpdateEquipment(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	var equipmentUpdateDto entities.EquipmentUpdateDto

	if err := c.BindJSON(&equipmentUpdateDto); err != nil {
		ResponseError(c, "failed to bind equipment dto", err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.services.Equipment.UpdateEquipment(id, equipmentUpdateDto)

	if err != nil {
		ResponseError(c, "failed to update equipment", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user updated successfully", map[string]interface{}{"equipment_id": id})
}
