// handlers.go (package: sellers)
package sellers

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi"
	"github.com/gin-gonic/gin"
)

type SellersHandler struct {
	Handler oapi.ServerInterface
}

func (h *SellersHandler) GetInventory(c *gin.Context) {
	h.Handler.GetInventory(c)
}
