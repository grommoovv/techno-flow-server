package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
)

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIp called")

	var signInDto domain.UserSignInDto

	if err := c.BindJSON(&signInDto); err != nil {
		ResponseError(c, "failed to bind sign-in dto", err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(signInDto)

	user, err := h.services.Auth.SignIn(signInDto)
	if err != nil {
		ResponseError(c, "failed to sign-in", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "signed in successfully", map[string]interface{}{"signed_in_user": user})
}

func (h *Handler) signOut(c *gin.Context) {
	ResponseError(c, "resource unavailable", "resource unavailable", http.StatusInternalServerError)
}
