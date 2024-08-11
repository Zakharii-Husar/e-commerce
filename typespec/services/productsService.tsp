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
    } | {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
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
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    @get
    @route("/{product_id}") 
    op getProductById(@path product_id: string): {
      @body product: ProductDetailsResponse;
    } | {
      @statusCode statusCode: 404;
      @body error: NotFoundError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    @put
    @route("/{product_id}") op updateProductById(@path product_id: string, @body updateProductReq: UpdateProductRequest): 
    {
      @statusCode statusCode: 204;
    }  | {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    } | {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    @delete
    @route("/{product_id}") op deleteProductById(@path product_id: string): {
      @statusCode statusCode: 204;
    } | {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    } | {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    //handle attachments
    @post
    @route("/{product_id}/attachement") op uploadProductAttachment(@path product_id: string, uploadProductImgReq: UploadProductImgRequest, 
      @header("Content-Type") contentType: 
      "image/jpeg" | 
      "image/jpg" |
       "image/png"
  ): {
    @body uploadProductImgResponse: UploadProductImgResponse;
  }  | {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  }| {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

    
    @get
    @route("/{product_id}/attachement/{attachement_id}") op getProductAttachmentById( @path product_id: string, @path attachement_id: string) : 
    {
      @header("Content-Type") contentType: "image/jpeg" | "image/jpg" | "image/png";
      @body imgData: GetProductImgResponse;
    } | {
      @statusCode statusCode: 404;
      @body error: NotFoundError;
    } | {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

  @delete
  @route("/{product_id}attachement/{attachement_id}") op deleteProductAttachmentById(@path product_id: string, @path attachement_id: string):  
  {
    @statusCode statusCode: 204;
  } | 
  {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  }| {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  //handle questions
  @get
  @route("/{product_id}/questions") op getProductQuestions(@path product_id: string):  {
    @statusCode statusCode: 204;
    @body questionResponse: QuestionResponse[] 
  } | {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  }| {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  @post
  @route("/{product_id}/questions") op postProductQuestion(@path product_id: string, @body postQuestionReq: QuestionRequest):  
  {
    @body questionResponse: QuestionResponse
  } | 
  {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  }| {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  @post 
  @route("/{product_id}/questions/{question_id}/respond") op postQuestionAnswer(@path product_id: string, @path question_id: string, postAnswerReq: AnswerToQuestionRequest):   
  {
    @body answerToQuestionResponse: AnswerToQuestionResponse;
  } | 
  {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  }| {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  //handle reviews
  @get
  @route("/{product_id}/reviews") op getProductReviews(@path product_id: string):  
  { 
    @body reviewResponse: ReviewResponse[] 
  }
  | {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  } | {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  @post
  @route("/{product_id}/reviews") op postProductReview(@path product_id: string, @body reviewReq: ReviewRequest):  
   {
    @body reviewResponse: ReviewResponse;
  } |
  {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  } | {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };

  @post
  @route("/{product_id}/reviews/{review_id}/respond") op respondToReview(@path product_id: string, @path review_id: string, @body answerToReviewReq: AnswerToReviewRequest):  
  {
    @body answerToReviewResponse: AnswerToReviewResponse;
  }|
  {
    @statusCode statusCode: 400;
    @body error: InvalidRequestError;
  } | 
  {
    @statusCode statusCode: 401;
    @body error: UnauthorizedError;
  } | {
    @statusCode statusCode: 404;
    @body error: NotFoundError;
  } | {
    @statusCode statusCode: 500;
    @body error: InternalServerError;
  };
} 