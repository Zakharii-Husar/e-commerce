package auth

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi/auth_server"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Handler auth_server.ServerInterface
}

func (h *AuthHandler) RecoverPassword(c *gin.Context) {
	h.Handler.RecoverPassword(c)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	h.Handler.SignIn(c)
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	h.Handler.SignOut(c)
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	h.Handler.SignUp(c)
}
