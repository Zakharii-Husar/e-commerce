package wishlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WishlistImpl implements the ServerInterface.
type WishlistImpl struct{}

// GetWishlistImpl creates a new WishlistImpl instance.
func GetWishlistImpl() *WishlistImpl {
	return &WishlistImpl{}
}

func (w *WishlistImpl) GetWishlistItems(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get wishlist items not implemented yet"})
}

func (w *WishlistImpl) DeleteItemFromWishlist(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Delete item from wishlist not implemented yet"})
}

func (w *WishlistImpl) AddItemToWishlist(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Add item to wishlist not implemented yet"})
}
