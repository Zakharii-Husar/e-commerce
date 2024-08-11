import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/productModels.tsp";
import "../models/errorModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Products service",
  })
@route("/api/products")
namespace ProductsService {
  //handle product
    @get
    @route("/") op getProducts():  
    {
      @body products: ProductOverviewResponse[]
    } | {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    };

    @post
    @route("/") op createProduct(@body createProductReq :CreateProductRequest):  
    { @statusCode statusCode: 201;
      @body productDetailsResponse: ProductDetailsResponse 
    }| {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    } | {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    };

    @get
    @route("/{product_id}") 
    op getProductById(@path product_id: string): {
      @body product: ProductDetailsResponse;
    } | {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    };

    @put
    @route("/{product_id}") op updateProductById(@path product_id: string, updateProductReq: UpdateProductRequest):  ProductDetailsResponse | InvalidRequestError | UnauthorizedError;

    @delete
    @route("/{product_id}") op deleteProductById(@path product_id: string):  void | InvalidRequestError | UnauthorizedError;

    //handle attachments
    @post
    @route("/{product_id}/attachement") op uploadProductAttachment(@path product_id: string, uploadProductImgReq: UploadProductImgRequest, 
      @header("Content-Type") contentType: 
      "image/jpeg" | 
      "image/jpg" |
       "image/png"
  ) : UploadProductImgResponse | InvalidRequestError | UnauthorizedError;

    
    @get
    @route("/{product_id}/attachement/{attachement_id}") op getProductAttachmentById( @path product_id: string, @path attachement_id: string) : GetProductImgResponse & {
      @header("Content-Type") contentType: 
      "image/jpeg" | 
      "image/jpg" |
       "image/png"
  } | NotFoundError | InvalidRequestError;

  @delete
  @route("/{product_id}attachement/{attachement_id}") op deleteProductAttachmentById(@path product_id: string, @path attachement_id: string):  void | InvalidRequestError | UnauthorizedError | NotFoundError;

  //handle questions
  @get
  @route("/{product_id}/questions") op getProductQuestions(@path product_id: string):  QuestionResponse[] | InvalidRequestError | NotFoundError;

  @post
  @route("/{product_id}/questions") op postProductQuestion(@path product_id: string, @body postQuestionReq: QuestionRequest):  QuestionResponse | InvalidRequestError | UnauthorizedError | NotFoundError;

  @post 
  @route("/{product_id}/questions/{question_id}/respond") op postQuestionAnswer(@path product_id: string, @path question_id: string, postAnswerReq: AnswerToQuestionRequest):   AnswerToQuestionResponse | InvalidRequestError | UnauthorizedError | NotFoundError;

  //handle reviews
  @get
  @route("/{product_id}/reviews") op getProductReviews(@path product_id: string):  ReviewResponse[] | InvalidRequestError  | NotFoundError;

  @post
  @route("/{product_id}/reviews") op postProductReview(@path product_id: string, @body reviewReq: ReviewRequest):  ReviewResponse | InvalidRequestError  | NotFoundError | UnauthorizedError;

  @post
  @route("/{product_id}/reviews/{review_id}/respond") op respondToReview(@path product_id: string, @path review_id: string, @body answerToReviewReq: AnswerToReviewRequest):  AnswerToReviewResponse | InvalidRequestError  | NotFoundError | UnauthorizedError;
} 