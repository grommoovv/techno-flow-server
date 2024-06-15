package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"strconv"
)

func (h *Handler) CreateReport(c *gin.Context) {
	const op = "report/Handler.CreateReport"
	var reportDto entities.ReportCreateDto

	if err := c.BindJSON(&reportDto); err != nil {
		ResponseError(c, "failed to bind report dto", err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Report.CreateReport(c, reportDto)

	if err != nil {
		ResponseError(c, "failed to create report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report created successfully", map[string]interface{}{"report_id": id})
}

func (h *Handler) GetAllReports(c *gin.Context) {
	const op = "report/Handler.GetAllReports"
	reports, err := h.services.Report.GetAllReports(c)

	if err != nil {
		ResponseError(c, "failed to fetch reports", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "reports fetched successfully", reports)
}

func (h *Handler) GetReportById(c *gin.Context) {
	const op = "report/Handler.GetReportById"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	report, err := h.services.Report.GetReportById(c, id)

	if err != nil {
		ResponseError(c, "failed to fetch report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report fetched successfully", report)
}

func (h *Handler) GetReportsByUserId(c *gin.Context) {
	const op = "report/Handler.GetReportsByUserId"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	reports, err := h.services.Report.GetReportsByUserId(c, id)

	if err != nil {
		ResponseError(c, "failed to fetch reports", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "reports fetched successfully", reports)
}

func (h *Handler) DeleteReport(c *gin.Context) {
	const op = "report/Handler.DeleteReport"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.Report.DeleteReport(c, id)
	if err != nil {
		ResponseError(c, "failed to delete report", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "report deleted successfully", map[string]interface{}{
		"report_id": id,
	})
}

func (h *Handler) UpdateReport(c *gin.Context) {}
