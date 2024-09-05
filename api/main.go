package main

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/admin"
	"github.com/Zakharii-Husar/e-commerce/api/internal/auth"
	"github.com/Zakharii-Husar/e-commerce/api/internal/auto_generated"
	"github.com/Zakharii-Husar/e-commerce/api/internal/cart"
	"github.com/Zakharii-Husar/e-commerce/api/internal/orders"
	"github.com/Zakharii-Husar/e-commerce/api/internal/products"
	"github.com/Zakharii-Husar/e-commerce/api/internal/sellers"
	"github.com/Zakharii-Husar/e-commerce/api/internal/users"
	"github.com/Zakharii-Husar/e-commerce/api/internal/wishlist"
	"github.com/gin-gonic/gin"
)

// Generate types (models)
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o ./internal/auto_generated/models_gen.go --package=auto_generated --generate=types ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

// Generate server routes for Gin
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o ./internal/auto_generated/routes_gen.go --package=auto_generated --generate=gin ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

type CombinedImpl struct {
	admin.AdminImpl
	auth.AuthImpl
	cart.CartImpl
	orders.OrdersImpl
	products.ProductsImpl
	sellers.SellersImpl
	users.UsersImpl
	wishlist.WishlistImpl
}

func main() {
	router := gin.Default()

	combinedImpl := &CombinedImpl{
		AdminImpl:    *admin.GetAdminImpl(),
		AuthImpl:     *auth.GetAuthImpl(),
		CartImpl:     *cart.GetCartImpl(),
		OrdersImpl:   *orders.GetOrdersImpl(),
		ProductsImpl: *products.GetProductsImpl(),
		SellersImpl:  *sellers.GetSellersImpl(),
		UsersImpl:    *users.GetUsersImpl(),
		WishlistImpl: *wishlist.GetWishlistImpl(),
	}

	auto_generated.RegisterHandlers(router, combinedImpl)

	router.Run(":8080")
}
