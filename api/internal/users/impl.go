// implementation.go (package: users)
package users

import (
	"net/http"

	"github.com/Zakharii-Husar/e-commerce/api/generated/codegen"
	"github.com/gin-gonic/gin"
)

// UsersImpl implements the ServerInterface.
type UsersImpl struct{}

// GetUsersImpl creates a new UsersImpl instance.
func GetUsersImpl() *UsersImpl {
	return &UsersImpl{}
}

func (u *UsersImpl) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get users not implemented yet"})
}

func (u *UsersImpl) SearchUsers(c *gin.Context, params codegen.SearchUsersParams) {
	// Implement the actual search logic here, for now just return a not implemented response.
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Search users not implemented yet"})
}

func (u *UsersImpl) DeleteUser(ctx *gin.Context, userId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Delete user not implemented yet"})
}

func (u *UsersImpl) GetUserById(ctx *gin.Context, userId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get user by ID not implemented yet"})
}

func (u *UsersImpl) UpdateUserDetails(ctx *gin.Context, userId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Update user details not implemented yet"})
}
