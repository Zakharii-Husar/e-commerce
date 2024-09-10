package admin

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/codegen"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Handler codegen.ServerInterface
}

func (h *AdminHandler) SuspendProduct(c *gin.Context, productId string) {
	h.Handler.SuspendProduct(c, productId)
}

func (h *AdminHandler) SuspendUser(c *gin.Context, userId string) {
	h.Handler.SuspendUser(c, userId)
}

func (h *AdminHandler) GetAllTransactions(c *gin.Context) {
	h.Handler.GetAllTransactions(c)
}
