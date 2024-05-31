package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type DataResponse struct {
	Status     string       `json:"status"`
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	Error      *string      `json:"error"`
	Data       *interface{} `json:"data"`
	Timestamp  time.Time    `json:"timestamp"`
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	resp := DataResponse{
		Status:     "success",
		StatusCode: 200,
		Message:    strings.ToLower(message),
		Data:       &data,
		Timestamp:  time.Now().UTC(),
	}
	c.JSON(http.StatusOK, resp)
}

func ResponseError(c *gin.Context, message string, err string, code int) {
	err = strings.ToLower(err)
	resp := DataResponse{
		Status:     "error",
		StatusCode: code,
		Message:    strings.ToLower(message),
		Error:      &err,
		Timestamp:  time.Now().UTC(),
	}
	c.JSON(code, resp)
}
