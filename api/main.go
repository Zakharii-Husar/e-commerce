package main

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/admin"
	"github.com/Zakharii-Husar/e-commerce/api/internal/auth"
	"github.com/Zakharii-Husar/e-commerce/api/internal/cart"
	"github.com/Zakharii-Husar/e-commerce/api/internal/orders"
	"github.com/Zakharii-Husar/e-commerce/api/internal/products"
	"github.com/Zakharii-Husar/e-commerce/api/internal/sellers"
	"github.com/Zakharii-Husar/e-commerce/api/internal/users"
	"github.com/Zakharii-Husar/e-commerce/api/internal/wishlist"
)

//Generate ORM from SQL Schema:
//go:generate sqlboiler psql -c ./config/sqlboiler.toml

// Generate types (models)
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/codegen_types.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

// Generate types (models)
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./config/codegen_gin_server.yaml ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

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

	/*
		Host:     "localhost",
		Port:     5432,
		Username: "e_commerce_user",
		Password: "e_commerce_password",
		Database: "e_commerce",
		Type:     2,
	*/

	// router := gin.Default()

	// combinedImpl := &CombinedImpl{
	// 	AdminImpl:    *admin.GetAdminImpl(),
	// 	AuthImpl:     *auth.GetAuthImpl(),
	// 	CartImpl:     *cart.GetCartImpl(),
	// 	OrdersImpl:   *orders.GetOrdersImpl(),
	// 	ProductsImpl: *products.GetProductsImpl(),
	// 	SellersImpl:  *sellers.GetSellersImpl(),
	// 	UsersImpl:    *users.GetUsersImpl(),
	// 	WishlistImpl: *wishlist.GetWishlistImpl(),
	// }

	// auto_generated.RegisterHandlers(router, combinedImpl)

	// router.Run(":8080")
}
