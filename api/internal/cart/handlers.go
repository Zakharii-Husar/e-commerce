package cart

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/auto_generated"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Handler auto_generated.ServerInterface
}

func (h *CartHandler) GetCartItems(c *gin.Context) {
	h.Handler.GetCartItems(c)
}

func (h *CartHandler) DeleteItemFromCart(c *gin.Context, productId string) {
	h.Handler.DeleteItemFromCart(c, productId)
}

func (h *CartHandler) UpdateItemQuantity(c *gin.Context, productId string) {
	h.Handler.UpdateItemQuantity(c, productId)
}

func (h *CartHandler) AddItemToCart(c *gin.Context, productId string) {
	h.Handler.AddItemToCart(c, productId)
}
