// implementation.go (package: sellers)
package sellers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SellersImpl implements the ServerInterface.
type SellersImpl struct{}

// GetSellersImpl creates a new SellersImpl instance.
func GetSellersImpl() *SellersImpl {
	return &SellersImpl{}
}

func (s *SellersImpl) GetInventory(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get inventory not implemented yet"})
}
