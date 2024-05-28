package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
	"time"
)

type SignInResponse struct {
	User        domain.User `json:"user"`
	AccessToken string      `json:"accessToken"`
}

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIp called")

	var signInDto domain.UserSignInDto

	if err := c.BindJSON(&signInDto); err != nil {
		ResponseError(c, "failed to bind sign-in dto", err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(signInDto)

	user, refreshToken, accessToken, err := h.services.Auth.SignIn(signInDto)
	if err != nil {
		ResponseError(c, "failed to sign-in", err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
	})

	//c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	signInResponse := SignInResponse{
		User:        user,
		AccessToken: accessToken,
	}

	ResponseSuccess(c, "user signed in successfully", signInResponse)
}

func (h *Handler) signOut(c *gin.Context) {
	ResponseError(c, "resource unavailable", "resource unavailable", http.StatusInternalServerError)
}
