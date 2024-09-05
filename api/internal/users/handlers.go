// handlers.go (package: users)
package users

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/auto_generated"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	Handler auto_generated.ServerInterface
}

func (h *UsersHandler) GetUsers(c *gin.Context) {
	h.Handler.GetUsers(c)
}

func (h *UsersHandler) SearchUsers(c *gin.Context) {
	// Extract query parameters into SearchUsersParams struct
	params := auto_generated.SearchUsersParams{
		// Assuming there's a "query" parameter in the search query.
		Query: c.Query("query"),
	}
	h.Handler.SearchUsers(c, params)
}

func (h *UsersHandler) DeleteUser(c *gin.Context, userId string) {
	h.Handler.DeleteUser(c, userId)
}

func (h *UsersHandler) GetUserById(c *gin.Context, userId string) {
	h.Handler.GetUserById(c, userId)
}

func (h *UsersHandler) UpdateUserDetails(c *gin.Context, userId string) {
	h.Handler.UpdateUserDetails(c, userId)
}
