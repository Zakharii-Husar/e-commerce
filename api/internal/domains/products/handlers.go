package products

import (
	"github.com/Zakharii-Husar/e-commerce/api/generated/oapi"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	Handler oapi.ServerInterface
}

func (h *ProductsHandler) GetProducts(c *gin.Context) {
	h.Handler.GetProducts(c)
}

func (h *ProductsHandler) CreateProduct(c *gin.Context) {
	h.Handler.CreateProduct(c)
}

func (h *ProductsHandler) DeleteProductById(c *gin.Context, productId string) {
	h.Handler.DeleteProductById(c, productId)
}

func (h *ProductsHandler) GetProductById(c *gin.Context, productId string) {
	h.Handler.GetProductById(c, productId)
}

func (h *ProductsHandler) UpdateProductById(c *gin.Context, productId string) {
	h.Handler.UpdateProductById(c, productId)
}

func (h *ProductsHandler) UploadProductAttachment(c *gin.Context, productId string) {
	h.Handler.UploadProductAttachment(c, productId)
}

func (h *ProductsHandler) GetProductAttachmentById(c *gin.Context, productId, attachmentId string) {
	h.Handler.GetProductAttachmentById(c, productId, attachmentId)
}

func (h *ProductsHandler) GetProductQuestions(c *gin.Context, productId string) {
	h.Handler.GetProductQuestions(c, productId)
}

func (h *ProductsHandler) PostProductQuestion(c *gin.Context, productId string) {
	h.Handler.PostProductQuestion(c, productId)
}

func (h *ProductsHandler) PostQuestionAnswer(c *gin.Context, productId, questionId string) {
	h.Handler.PostQuestionAnswer(c, productId, questionId)
}

func (h *ProductsHandler) GetProductReviews(c *gin.Context, productId string) {
	h.Handler.GetProductReviews(c, productId)
}

func (h *ProductsHandler) PostProductReview(c *gin.Context, productId string) {
	h.Handler.PostProductReview(c, productId)
}

func (h *ProductsHandler) RespondToReview(c *gin.Context, productId, reviewId string) {
	h.Handler.RespondToReview(c, productId, reviewId)
}

func (h *ProductsHandler) DeleteProductAttachmentById(c *gin.Context, productId, attachmentId string) {
	h.Handler.DeleteProductAttachmentById(c, productId, attachmentId)
}
