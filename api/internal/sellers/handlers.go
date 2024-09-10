// handlers.go (package: sellers)
package sellers

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/codegen"
	"github.com/gin-gonic/gin"
)

type SellersHandler struct {
	Handler codegen.ServerInterface
}

func (h *SellersHandler) GetInventory(c *gin.Context) {
	h.Handler.GetInventory(c)
}
