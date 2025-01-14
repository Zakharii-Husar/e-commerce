// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package models

// Defines values for DateTimeKnownEncoding.
const (
	Rfc3339       DateTimeKnownEncoding = "rfc3339"
	Rfc7231       DateTimeKnownEncoding = "rfc7231"
	UnixTimestamp DateTimeKnownEncoding = "unixTimestamp"
)

// Defines values for ReviewRequestRate.
const (
	N1 ReviewRequestRate = 1
	N2 ReviewRequestRate = 2
	N3 ReviewRequestRate = 3
	N4 ReviewRequestRate = 4
	N5 ReviewRequestRate = 5
)

// AnswerToQuestionRequest defines model for AnswerToQuestionRequest.
type AnswerToQuestionRequest struct {
	Answer   string `json:"answer"`
	Username string `json:"username"`
}

// AnswerToQuestionResponse defines model for AnswerToQuestionResponse.
type AnswerToQuestionResponse = AnswerToQuestionRequest

// AnswerToReviewRequest defines model for AnswerToReviewRequest.
type AnswerToReviewRequest struct {
	Answer   string `json:"answer"`
	Username string `json:"username"`
}

// AnswerToReviewResponse defines model for AnswerToReviewResponse.
type AnswerToReviewResponse = AnswerToReviewRequest

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	Token string `json:"token"`
}

// CancelOrderResponse defines model for CancelOrderResponse.
type CancelOrderResponse struct {
	Success bool `json:"success"`
}

// CreateProductRequest defines model for CreateProductRequest.
type CreateProductRequest struct {
	ProductCategory    string   `json:"productCategory"`
	ProductDescription *string  `json:"productDescription,omitempty"`
	ProductHeight      *float32 `json:"productHeight,omitempty"`
	ProductLenght      *float32 `json:"productLenght,omitempty"`
	ProductMaterial    *string  `json:"productMaterial,omitempty"`
	ProductName        string   `json:"productName"`
	ProductPrice       float32  `json:"productPrice"`
	ProductQuantity    int      `json:"productQuantity"`
	ProductWeight      *float32 `json:"productWeight,omitempty"`
	ProductWidth       *float32 `json:"productWidth,omitempty"`
}

// DateTimeKnownEncoding Known encoding to use on utcDateTime or offsetDateTime
type DateTimeKnownEncoding string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetProductImgResponse defines model for GetProductImgResponse.
type GetProductImgResponse struct {
	Image []byte `json:"image"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError = ErrorResponse

// InvalidRequestError defines model for InvalidRequestError.
type InvalidRequestError = ErrorResponse

// NotFoundError defines model for NotFoundError.
type NotFoundError = ErrorResponse

// OrderDetailsResponse defines model for OrderDetailsResponse.
type OrderDetailsResponse = OrderOverviewResponse

// OrderOverviewResponse defines model for OrderOverviewResponse.
type OrderOverviewResponse struct {
	// CreatedAt Known encoding to use on utcDateTime or offsetDateTime
	CreatedAt   DateTimeKnownEncoding `json:"createdAt"`
	OrderId     int64                 `json:"orderId"`
	OrderPrice  float32               `json:"orderPrice"`
	OrderStatus string                `json:"orderStatus"`
}

// OrderedItemModel defines model for OrderedItemModel.
type OrderedItemModel = ProductOverviewResponse

// PassRecoverRequest defines model for PassRecoverRequest.
type PassRecoverRequest struct {
	UserNameOrEmail string `json:"userNameOrEmail"`
}

// ProductDetailsResponse defines model for ProductDetailsResponse.
type ProductDetailsResponse = ProductOverviewResponse

// ProductOverviewResponse defines model for ProductOverviewResponse.
type ProductOverviewResponse struct {
	ProductAttachements []string              `json:"productAttachements"`
	ProductCategory     string                `json:"productCategory"`
	ProductId           int64                 `json:"productId"`
	ProductName         string                `json:"productName"`
	ProductPrice        float32               `json:"productPrice"`
	Rating              ProductRatingResponse `json:"rating"`
	SellerUsername      string                `json:"sellerUsername"`
}

// ProductRatingResponse defines model for ProductRatingResponse.
type ProductRatingResponse struct {
	AverageRating float32 `json:"averageRating"`
	RatedTimes    int32   `json:"ratedTimes"`
}

// QuestionRequest defines model for QuestionRequest.
type QuestionRequest struct {
	Question string `json:"question"`
}

// QuestionResponse defines model for QuestionResponse.
type QuestionResponse struct {
	// DateTime Known encoding to use on utcDateTime or offsetDateTime
	DateTime   DateTimeKnownEncoding `json:"dateTime"`
	Question   string                `json:"question"`
	QuestionId int64                 `json:"questionId"`
	Username   string                `json:"username"`
}

// ReviewRequest defines model for ReviewRequest.
type ReviewRequest struct {
	Comment *string           `json:"comment,omitempty"`
	Rate    ReviewRequestRate `json:"rate"`
}

// ReviewRequestRate defines model for ReviewRequest.Rate.
type ReviewRequestRate float32

// ReviewResponse defines model for ReviewResponse.
type ReviewResponse = ReviewRequest

// SignInRequest defines model for SignInRequest.
type SignInRequest struct {
	Password        string `json:"password"`
	UserNameOrEmail string `json:"userNameOrEmail"`
}

// SignUpRequest defines model for SignUpRequest.
type SignUpRequest struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Email       string  `json:"email"`
	FullName    string  `json:"fullName"`
	Password    string  `json:"password"`
	PhoneNumber string  `json:"phoneNumber"`
	Province    string  `json:"province"`
	Street      string  `json:"street"`
	UnitNumber  *string `json:"unitNumber,omitempty"`
	UserName    string  `json:"userName"`
	ZipCode     string  `json:"zipCode"`
}

// TransactionResponse defines model for TransactionResponse.
type TransactionResponse struct {
	Amount            float32 `json:"amount"`
	OrderId           int64   `json:"orderId"`
	PayeeId           int64   `json:"payeeId"`
	PayeeUsername     string  `json:"payeeUsername"`
	PayerId           int64   `json:"payerId"`
	PayerUsername     string  `json:"payerUsername"`
	PaymentMethod     string  `json:"paymentMethod"`
	TransactionId     int64   `json:"transactionId"`
	TransactionStatus string  `json:"transactionStatus"`

	// TransactionTime Known encoding to use on utcDateTime or offsetDateTime
	TransactionTime DateTimeKnownEncoding `json:"transactionTime"`
	TransactionType string                `json:"transactionType"`
}

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError = ErrorResponse

// UpdateCartItemRequest defines model for UpdateCartItemRequest.
type UpdateCartItemRequest struct {
	Quantity *int32 `json:"quantity,omitempty"`
}

// UpdateOrderRequest defines model for UpdateOrderRequest.
type UpdateOrderRequest struct {
	NewStatus *string `json:"newStatus,omitempty"`
}

// UpdateOrderResponse defines model for UpdateOrderResponse.
type UpdateOrderResponse struct {
	Success bool `json:"success"`
}

// UpdateProductRequest defines model for UpdateProductRequest.
type UpdateProductRequest struct {
	ProductCategory    *string  `json:"productCategory,omitempty"`
	ProductDescription *string  `json:"productDescription,omitempty"`
	ProductName        *string  `json:"productName,omitempty"`
	ProductPrice       *float32 `json:"productPrice,omitempty"`
	ProductQuantity    *int     `json:"productQuantity,omitempty"`
	StockQuantity      *int32   `json:"stockQuantity,omitempty"`
}

// UpdateUserRequest defines model for UpdateUserRequest.
type UpdateUserRequest struct {
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	Email       *string `json:"email,omitempty"`
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Province    *string `json:"province,omitempty"`
	Street      *string `json:"street,omitempty"`
	UnitNumber  *string `json:"unitNumber,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	ZipCode     *string `json:"zipCode,omitempty"`
}

// UploadProductImgRequest defines model for UploadProductImgRequest.
type UploadProductImgRequest struct {
	Image []byte `json:"image"`
}

// UploadProductImgResponse defines model for UploadProductImgResponse.
type UploadProductImgResponse struct {
	AttachementId string `json:"attachementId"`
	Success       bool   `json:"success"`
}

// UserDetailsResponse defines model for UserDetailsResponse.
type UserDetailsResponse struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Email       string  `json:"email"`
	FullName    string  `json:"fullName"`
	PhoneNumber string  `json:"phoneNumber"`
	Province    string  `json:"province"`
	Street      string  `json:"street"`
	UnitNumber  *string `json:"unitNumber,omitempty"`
	UserId      string  `json:"userId"`
	UserName    string  `json:"userName"`
	ZipCode     string  `json:"zipCode"`
}

// RecoverPasswordJSONBody defines parameters for RecoverPassword.
type RecoverPasswordJSONBody struct {
	PassRecoverRequest PassRecoverRequest `json:"passRecoverRequest"`
}

// SignInJSONBody defines parameters for SignIn.
type SignInJSONBody struct {
	SignInReq SignInRequest `json:"signInReq"`
}

// SignUpJSONBody defines parameters for SignUp.
type SignUpJSONBody struct {
	SignUpReq SignUpRequest `json:"signUpReq"`
}

// PostQuestionAnswerJSONBody defines parameters for PostQuestionAnswer.
type PostQuestionAnswerJSONBody struct {
	PostAnswerReq AnswerToQuestionRequest `json:"postAnswerReq"`
}

// SearchUsersParams defines parameters for SearchUsers.
type SearchUsersParams struct {
	Query string `form:"query" json:"query"`
}

// RecoverPasswordJSONRequestBody defines body for RecoverPassword for application/json ContentType.
type RecoverPasswordJSONRequestBody RecoverPasswordJSONBody

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody SignInJSONBody

// SignUpJSONRequestBody defines body for SignUp for application/json ContentType.
type SignUpJSONRequestBody SignUpJSONBody

// UpdateItemQuantityJSONRequestBody defines body for UpdateItemQuantity for application/json ContentType.
type UpdateItemQuantityJSONRequestBody = UpdateCartItemRequest

// UpdateOrderJSONRequestBody defines body for UpdateOrder for application/json ContentType.
type UpdateOrderJSONRequestBody = UpdateOrderRequest

// CreateProductJSONRequestBody defines body for CreateProduct for application/json ContentType.
type CreateProductJSONRequestBody = CreateProductRequest

// UpdateProductByIdJSONRequestBody defines body for UpdateProductById for application/json ContentType.
type UpdateProductByIdJSONRequestBody = UpdateProductRequest

// PostProductQuestionJSONRequestBody defines body for PostProductQuestion for application/json ContentType.
type PostProductQuestionJSONRequestBody = QuestionRequest

// PostQuestionAnswerJSONRequestBody defines body for PostQuestionAnswer for application/json ContentType.
type PostQuestionAnswerJSONRequestBody PostQuestionAnswerJSONBody

// PostProductReviewJSONRequestBody defines body for PostProductReview for application/json ContentType.
type PostProductReviewJSONRequestBody = ReviewRequest

// RespondToReviewJSONRequestBody defines body for RespondToReview for application/json ContentType.
type RespondToReviewJSONRequestBody = AnswerToReviewRequest

// UpdateUserDetailsJSONRequestBody defines body for UpdateUserDetails for application/json ContentType.
type UpdateUserDetailsJSONRequestBody = UpdateUserRequest
