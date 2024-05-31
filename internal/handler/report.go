package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"strconv"
)

func (h *Handler) CreateReport(c *gin.Context) {
	var reportDto entities.ReportCreateDto

	if err := c.BindJSON(&reportDto); err != nil {
		ResponseError(c, "failed to bind report dto", err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Report.CreateReport(reportDto)

	if err != nil {
		ResponseError(c, "failed to create report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report created successfully", map[string]interface{}{"report_id": id})
}

func (h *Handler) GetAllReports(c *gin.Context) {
	reports, err := h.services.Report.GetAllReports()

	if err != nil {
		ResponseError(c, "failed to fetch reports", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "reports fetched successfully", reports)
}

func (h *Handler) GetReportById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	report, err := h.services.Report.GetReportById(id)

	if err != nil {
		ResponseError(c, "failed to fetch report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report fetched successfully", report)
}

func (h *Handler) DeleteReport(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.Report.DeleteReport(id)
	if err != nil {
		ResponseError(c, "failed to delete report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report deleted successfully", map[string]interface{}{
		"report_id": id,
	})
}

func (h *Handler) UpdateReport(c *gin.Context) {}
