package wishlist

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/codegen"
	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	Handler codegen.ServerInterface
}

func (h *WishlistHandler) GetWishlistItems(c *gin.Context) {
	h.Handler.GetWishlistItems(c)
}

func (h *WishlistHandler) DeleteItemFromWishlist(c *gin.Context, productId string) {
	h.Handler.DeleteItemFromWishlist(c, productId)
}

func (h *WishlistHandler) AddItemToWishlist(c *gin.Context, productId string) {
	h.Handler.AddItemToWishlist(c, productId)
}
