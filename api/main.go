package main

import (
	"github.com/Zakharii-Husar/e-commerce/api/internal/auto_generated"
	"github.com/gin-gonic/gin"
)

// Generate types (models)
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o ./internal/auto_generated/models_gen.go --package=auto_generated --generate=types ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

// Generate server routes for Gin
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o ./internal/auto_generated/routes_gen.go --package=auto_generated --generate=gin ../typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.yaml

func main() {
	router := gin.Default()

	auto_generated.RegisterHandlers(router, userImpl)

	router.Run(":8080")
}
