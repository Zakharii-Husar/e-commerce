// handlers.go (package: sellers)
package sellers

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi/seller_server"
	"github.com/gin-gonic/gin"
)

type SellersHandler struct {
	Handler seller_server.ServerInterface
}

func (h *SellersHandler) GetInventory(c *gin.Context) {
	h.Handler.GetInventory(c)
}
