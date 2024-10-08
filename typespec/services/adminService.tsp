import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/adminModels.tsp";


using TypeSpec.Http;
using TypeSpec.Rest;

//#suppress "@typespec/http/no-service-found" ""
@route("/api/admin")
@service({
    title: "Admin service",
  })
namespace ECommerce.AdminService {
    @patch
    @route("/suspend/user/{user_id}") op suspendUser(@path user_id: string): 
    {
      @statusCode statusCode: 204;
    } 
    |
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

    @patch
    @route("/suspend/product/{product_id}") op suspendProduct(@path product_id: string): 
    {
      @statusCode statusCode: 204;
    } 
    |
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

    @get
    @route("/transactions") op getAllTransactions():
    {
      @body transactionResponse: TransactionResponse[];
    } 
    |
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
} 