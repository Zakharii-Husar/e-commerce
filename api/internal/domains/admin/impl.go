package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminImpl implements the ServerInterface.
type AdminImpl struct{}

// GetAdminImpl creates a new AdminImpl instance.
func GetAdminImpl() *AdminImpl {
	return &AdminImpl{}
}

func (a *AdminImpl) SuspendProduct(c *gin.Context, productId string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

func (a *AdminImpl) SuspendUser(c *gin.Context, userId string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

func (a *AdminImpl) GetAllTransactions(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}
