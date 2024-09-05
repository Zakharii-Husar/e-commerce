package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductsImpl implements the ServerInterface.
type ProductsImpl struct{}

// GetProductsImpl creates a new ProductsImpl instance.
func GetProductsImpl() *ProductsImpl {
	return &ProductsImpl{}
}

func (p *ProductsImpl) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get products not implemented yet"})
}

func (p *ProductsImpl) CreateProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Create product not implemented yet"})
}

func (p *ProductsImpl) DeleteProductById(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Delete product not implemented yet"})
}

func (p *ProductsImpl) GetProductById(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get product by ID not implemented yet"})
}

func (p *ProductsImpl) UpdateProductById(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Update product not implemented yet"})
}

func (p *ProductsImpl) UploadProductAttachment(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Upload product attachment not implemented yet"})
}

func (p *ProductsImpl) GetProductAttachmentById(ctx *gin.Context, productId, attachmentId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get product attachment not implemented yet"})
}

func (p *ProductsImpl) GetProductQuestions(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get product questions not implemented yet"})
}

func (p *ProductsImpl) PostProductQuestion(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Post product question not implemented yet"})
}

func (p *ProductsImpl) PostQuestionAnswer(ctx *gin.Context, productId, questionId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Post question answer not implemented yet"})
}

func (p *ProductsImpl) GetProductReviews(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Get product reviews not implemented yet"})
}

func (p *ProductsImpl) PostProductReview(ctx *gin.Context, productId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Post product review not implemented yet"})
}

func (p *ProductsImpl) RespondToReview(ctx *gin.Context, productId, reviewId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Respond to review not implemented yet"})
}

func (p *ProductsImpl) DeleteProductAttachmentById(ctx *gin.Context, productId, attachmentId string) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Delete product attachment not implemented yet"})
}
