// handlers.go (package: sellers)
package sellers

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/auto_generated"
	"github.com/gin-gonic/gin"
)

type SellersHandler struct {
	Handler auto_generated.ServerInterface
}

func (h *SellersHandler) GetInventory(c *gin.Context) {
	h.Handler.GetInventory(c)
}
