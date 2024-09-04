package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminImpl implements the ServerInterface.
type AdminImpl struct{}

// GetUserImpl creates a new AdminImpl instance.
func GetUserImpl() *AdminImpl {
	return &AdminImpl{}
}

func (u *AdminImpl) SuspendProduct(c *gin.Context, productId string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

func (u *AdminImpl) SuspendUser(c *gin.Context, userId string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}

func (u *AdminImpl) GetAllTransactions(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
}
