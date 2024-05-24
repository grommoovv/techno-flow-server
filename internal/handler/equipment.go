package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
	"strconv"
)

func (h *Handler) CreateEquipment(c *gin.Context) {
	var equipmentDto domain.EquipmentCreateDto

	if err := c.BindJSON(&equipmentDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Equipment.CreateEquipment(equipmentDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllEquipment(c *gin.Context) {
	equipment, err := h.services.Equipment.GetAllEquipment()

	if err != nil {
		fmt.Printf("error fetching equipment: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"equipment": equipment,
	})
}

func (h *Handler) GetEquipmentById(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Printf("GetEquipmentById Called, param: %s\n", paramId)

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equipment, err := h.services.Equipment.GetEquipmentById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"equipment": equipment,
	})
}

func (h *Handler) DeleteEquipment(c *gin.Context) {
	fmt.Printf("DeleteEquipment Called, param: %s\n", c.Param("id"))
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.Equipment.DeleteEquipment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"equipment_id": id,
	})
}

func (h *Handler) UpdateEquipment(c *gin.Context) {
	fmt.Printf("UpdateEquipment Called, param: %s\n", c.Param("id"))
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var equipmentUpdateDto domain.EquipmentUpdateDto

	if err := c.BindJSON(&equipmentUpdateDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errUpdate := h.services.Equipment.UpdateEquipment(id, equipmentUpdateDto)

	if errUpdate != nil {
		fmt.Printf("Error Updating User: %v\n", errUpdate)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
