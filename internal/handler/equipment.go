package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
	"strconv"
)

func (h *Handler) CreateEquipment(c *gin.Context) {
	var equipmentDto domain.EquipmentCreateDto

	if err := c.BindJSON(&equipmentDto); err != nil {
		ResponseError(c, "failed to bind equipment dto", err.Error(), http.StatusInternalServerError)
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

	ResponseSuccess(c, "user fetched successfully", equipment)
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

	var equipmentUpdateDto domain.EquipmentUpdateDto

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
