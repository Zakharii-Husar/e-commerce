// handlers.go (package: orders)
package orders

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi/orders_server"
	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {
	Handler orders_server.ServerInterface
}

func (h *OrdersHandler) GetOrders(c *gin.Context) {
	h.Handler.GetOrders(c)
}

func (h *OrdersHandler) GetOrderById(c *gin.Context, orderId string) {
	h.Handler.GetOrderById(c, orderId)
}

func (h *OrdersHandler) CancelOrder(c *gin.Context, orderId string) {
	h.Handler.CancelOrder(c, orderId)
}

func (h *OrdersHandler) UpdateOrder(c *gin.Context, orderId string) {
	h.Handler.UpdateOrder(c, orderId)
}
