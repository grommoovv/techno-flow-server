package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
		return
	}
}

func loggingMiddleware(c *gin.Context) {
	log.Printf("%s: [%s] - %s ", time.Now().Format(time.RFC3339), c.Request.Method, c.Request.RequestURI)
	c.Next()
}

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		ResponseError(c, "", "empty authorization header", http.StatusUnauthorized)
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		ResponseError(c, "invalid authorization header", "", http.StatusUnauthorized)
		c.Abort()
		return
	}

	userId, err := h.services.Token.ParseAccessToken(headerParts[1])

	if err != nil {
		ResponseError(c, "", err.Error(), http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Set(userCtx, userId)
}
