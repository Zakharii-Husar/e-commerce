// Package products_server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package products_server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/products/)
	GetProducts(c *gin.Context)

	// (POST /api/products/)
	CreateProduct(c *gin.Context)

	// (DELETE /api/products/{product_id})
	DeleteProductById(c *gin.Context, productId string)

	// (GET /api/products/{product_id})
	GetProductById(c *gin.Context, productId string)

	// (PUT /api/products/{product_id})
	UpdateProductById(c *gin.Context, productId string)

	// (POST /api/products/{product_id}/attachement)
	UploadProductAttachment(c *gin.Context, productId string)

	// (DELETE /api/products/{product_id}/attachement/{attachement_id})
	DeleteProductAttachmentById(c *gin.Context, productId string, attachementId string)

	// (GET /api/products/{product_id}/attachement/{attachement_id})
	GetProductAttachmentById(c *gin.Context, productId string, attachementId string)

	// (GET /api/products/{product_id}/questions)
	GetProductQuestions(c *gin.Context, productId string)

	// (POST /api/products/{product_id}/questions)
	PostProductQuestion(c *gin.Context, productId string)

	// (POST /api/products/{product_id}/questions/{question_id}/respond)
	PostQuestionAnswer(c *gin.Context, productId string, questionId string)

	// (GET /api/products/{product_id}/reviews)
	GetProductReviews(c *gin.Context, productId string)

	// (POST /api/products/{product_id}/reviews)
	PostProductReview(c *gin.Context, productId string)

	// (POST /api/products/{product_id}/reviews/{review_id}/respond)
	RespondToReview(c *gin.Context, productId string, reviewId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetProducts operation middleware
func (siw *ServerInterfaceWrapper) GetProducts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProducts(c)
}

// CreateProduct operation middleware
func (siw *ServerInterfaceWrapper) CreateProduct(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateProduct(c)
}

// DeleteProductById operation middleware
func (siw *ServerInterfaceWrapper) DeleteProductById(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteProductById(c, productId)
}

// GetProductById operation middleware
func (siw *ServerInterfaceWrapper) GetProductById(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductById(c, productId)
}

// UpdateProductById operation middleware
func (siw *ServerInterfaceWrapper) UpdateProductById(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateProductById(c, productId)
}

// UploadProductAttachment operation middleware
func (siw *ServerInterfaceWrapper) UploadProductAttachment(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UploadProductAttachment(c, productId)
}

// DeleteProductAttachmentById operation middleware
func (siw *ServerInterfaceWrapper) DeleteProductAttachmentById(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "attachement_id" -------------
	var attachementId string

	err = runtime.BindStyledParameterWithOptions("simple", "attachement_id", c.Param("attachement_id"), &attachementId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter attachement_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteProductAttachmentById(c, productId, attachementId)
}

// GetProductAttachmentById operation middleware
func (siw *ServerInterfaceWrapper) GetProductAttachmentById(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "attachement_id" -------------
	var attachementId string

	err = runtime.BindStyledParameterWithOptions("simple", "attachement_id", c.Param("attachement_id"), &attachementId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter attachement_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductAttachmentById(c, productId, attachementId)
}

// GetProductQuestions operation middleware
func (siw *ServerInterfaceWrapper) GetProductQuestions(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductQuestions(c, productId)
}

// PostProductQuestion operation middleware
func (siw *ServerInterfaceWrapper) PostProductQuestion(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostProductQuestion(c, productId)
}

// PostQuestionAnswer operation middleware
func (siw *ServerInterfaceWrapper) PostQuestionAnswer(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "question_id" -------------
	var questionId string

	err = runtime.BindStyledParameterWithOptions("simple", "question_id", c.Param("question_id"), &questionId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter question_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostQuestionAnswer(c, productId, questionId)
}

// GetProductReviews operation middleware
func (siw *ServerInterfaceWrapper) GetProductReviews(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductReviews(c, productId)
}

// PostProductReview operation middleware
func (siw *ServerInterfaceWrapper) PostProductReview(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostProductReview(c, productId)
}

// RespondToReview operation middleware
func (siw *ServerInterfaceWrapper) RespondToReview(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "review_id" -------------
	var reviewId string

	err = runtime.BindStyledParameterWithOptions("simple", "review_id", c.Param("review_id"), &reviewId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter review_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RespondToReview(c, productId, reviewId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/products/", wrapper.GetProducts)
	router.POST(options.BaseURL+"/api/products/", wrapper.CreateProduct)
	router.DELETE(options.BaseURL+"/api/products/:product_id", wrapper.DeleteProductById)
	router.GET(options.BaseURL+"/api/products/:product_id", wrapper.GetProductById)
	router.PUT(options.BaseURL+"/api/products/:product_id", wrapper.UpdateProductById)
	router.POST(options.BaseURL+"/api/products/:product_id/attachement", wrapper.UploadProductAttachment)
	router.DELETE(options.BaseURL+"/api/products/:product_id/attachement/:attachement_id", wrapper.DeleteProductAttachmentById)
	router.GET(options.BaseURL+"/api/products/:product_id/attachement/:attachement_id", wrapper.GetProductAttachmentById)
	router.GET(options.BaseURL+"/api/products/:product_id/questions", wrapper.GetProductQuestions)
	router.POST(options.BaseURL+"/api/products/:product_id/questions", wrapper.PostProductQuestion)
	router.POST(options.BaseURL+"/api/products/:product_id/questions/:question_id/respond", wrapper.PostQuestionAnswer)
	router.GET(options.BaseURL+"/api/products/:product_id/reviews", wrapper.GetProductReviews)
	router.POST(options.BaseURL+"/api/products/:product_id/reviews", wrapper.PostProductReview)
	router.POST(options.BaseURL+"/api/products/:product_id/reviews/:review_id/respond", wrapper.RespondToReview)
}
