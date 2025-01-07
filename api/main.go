package main

import (
	"github.com/Zakharii-Husar/e-commerce/api/db/database"
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/admin"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/auth"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/cart"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/orders"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/products"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/sellers"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/users"
	"github.com/Zakharii-Husar/e-commerce/api/internal/domains/wishlist"
	"github.com/gin-gonic/gin"
)

// Generate ORM from SQL Schema:
//go:generate sqlboiler psql -c ./config/sqlboiler.toml

// Generate types (models)
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/models.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

// Generate server interfaces
//go:generate mkdir -p ./generated/oapi/admin_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/admin_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.AdminService.yaml

//go:generate mkdir -p ./generated/oapi/auth_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/auth_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.AuthService.yaml

//go:generate mkdir -p ./generated/oapi/cart_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/cart_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.CartService.yaml

//go:generate mkdir -p ./generated/oapi/orders_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/orders_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.OrdersService.yaml

//go:generate mkdir -p ./generated/oapi/products_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/products_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.ProductsService.yaml

//go:generate mkdir -p ./generated/oapi/seller_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/seller_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.SellerService.yaml

//go:generate mkdir -p ./generated/oapi/users_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/users_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.UsersService.yaml

//go:generate mkdir -p ./generated/oapi/wishlist_server
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/wishlist_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.WishlistService.yaml

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
	database.InitDB()

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

	oapi.RegisterHandlers(router, combinedImpl)

	router.Run(":8080")
}
