package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type DataResponse struct {
	Status    string       `json:"status"`
	Message   string       `json:"message"`
	Data      *interface{} `json:"data"`
	Error     *string      `json:"error"`
	Code      int          `json:"code"`
	Timestamp time.Time    `json:"timestamp"`
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	resp := DataResponse{
		Status:    "success",
		Message:   strings.ToLower(message),
		Data:      &data,
		Code:      200,
		Timestamp: time.Now().UTC(),
	}
	c.JSON(http.StatusOK, resp)
}

func ResponseError(c *gin.Context, message string, err string, code int) {
	err = strings.ToLower(err)
	resp := DataResponse{
		Status:    "error",
		Message:   strings.ToLower(message),
		Error:     &err,
		Code:      code,
		Timestamp: time.Now().UTC(),
	}
	c.JSON(http.StatusOK, resp)
}
