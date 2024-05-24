package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
	"strconv"
)

func (h *Handler) CreateReport(c *gin.Context) {
	fmt.Println("CreateReport called")

	var reportDto domain.ReportCreateDto

	if err := c.BindJSON(&reportDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Report.CreateReport(reportDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllReports(c *gin.Context) {
	reports, err := h.services.Report.GetAllReports()

	if err != nil {
		fmt.Printf("error fetching reports: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"reports": reports,
	})
}

func (h *Handler) GetReportById(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Printf("GetReportById called, param: %s\n", paramId)

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report, err := h.services.Report.GetReportById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"report": report,
	})
}

func (h *Handler) DeleteReport(c *gin.Context) {
	fmt.Printf("DeleteReport called, param: %s\n", c.Param("id"))
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.Report.DeleteReport(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"report_id": id,
	})
}

func (h *Handler) UpdateReport(c *gin.Context) {}
