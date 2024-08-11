import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/adminModels.tsp";


using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Admin service",
  })
@route("/api/admin")
namespace AdminService {
    @patch
    @route("/suspend/user/{user_id}") op suspendUser(@path user_id: string): 
    {
      @statusCode statusCode: 204;
    } 
    |
    {
      @statusCode: 400;
      @body error: InvalidRequestError;
    } | 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    } | {
      @statusCode statusCode: 404;
      @body error: NotFoundError;
    };

    @patch
    @route("/suspend/product/{product_id}") op suspendProduct(@path product_id: string): 
    {
      @statusCode statusCode: 204;
    } 
    |
    {
      @statusCode: 400;
      @body error: InvalidRequestError;
    } | 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    } | {
      @statusCode statusCode: 404;
      @body error: NotFoundError;
    };

    @get
    @route("/transactions") op getAllTransactions():
    {
      @body transactionResponse: TransactionResponse[];
    } 
    |
    {
      @statusCode: 400;
      @body error: InvalidRequestError;
    } | 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    } | {
      @statusCode statusCode: 404;
      @body error: NotFoundError;
    };
} 