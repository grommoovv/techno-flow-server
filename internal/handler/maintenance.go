package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateMaintenance(c *gin.Context) {}

func (h *Handler) GetAllMaintenance(c *gin.Context) {
	const op = "maintenance/Handler.GetAllMaintenance"
	maintenance, err := h.services.Maintenance.GetAll(c)

	if err != nil {
		ResponseError(c, "failed to fetch maintenance", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "maintenance fetched successfully", maintenance)
}

func (h *Handler) GetMaintenanceById(c *gin.Context) {
	const op = "maintenance/Handler.GetMaintenanceById"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	maintenance, err := h.services.Maintenance.GetById(c, id)

	if err != nil {
		ResponseError(c, "failed to fetch maintenance", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "maintenance fetched successfully", maintenance)
}

func (h *Handler) DeleteMaintenance(c *gin.Context) {}

func (h *Handler) UpdateMaintenance(c *gin.Context) {}
