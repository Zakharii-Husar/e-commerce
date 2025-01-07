// Package auth_server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package auth_server

import (
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /api/auth/recover_password)
	RecoverPassword(c *gin.Context)

	// (POST /api/auth/sign_in)
	SignIn(c *gin.Context)

	// (GET /api/auth/sign_out)
	SignOut(c *gin.Context)

	// (POST /api/auth/sign_up)
	SignUp(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// RecoverPassword operation middleware
func (siw *ServerInterfaceWrapper) RecoverPassword(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RecoverPassword(c)
}

// SignIn operation middleware
func (siw *ServerInterfaceWrapper) SignIn(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SignIn(c)
}

// SignOut operation middleware
func (siw *ServerInterfaceWrapper) SignOut(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SignOut(c)
}

// SignUp operation middleware
func (siw *ServerInterfaceWrapper) SignUp(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SignUp(c)
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

	router.POST(options.BaseURL+"/api/auth/recover_password", wrapper.RecoverPassword)
	router.POST(options.BaseURL+"/api/auth/sign_in", wrapper.SignIn)
	router.GET(options.BaseURL+"/api/auth/sign_out", wrapper.SignOut)
	router.POST(options.BaseURL+"/api/auth/sign_up", wrapper.SignUp)
}
