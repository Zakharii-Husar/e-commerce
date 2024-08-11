import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/userModels.tsp";
import "../models/productModels.tsp";
import "../models/cartModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Cart Service",
})
@route("/api/cart")
namespace CartService {
    
    @get
    @route("/")
    op getCartItems(): {
        @body items: ProductOverviewResponse[];
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

    @post
    @route("/items/{product_id}")
    op addItemToCart(@path product_id: string): {
        @statusCode statusCode: 204;
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    }| {
        @statusCode statusCode: 500;
        @body error: InternalServerError;
      };

    @patch
    @route("/items/{product_id}")
    op updateItemQuantity(@path product_id: string, @body updateCartItemReq: UpdateCartItemRequest): {
        @statusCode statusCode: 204;
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    }| {
        @statusCode statusCode: 500;
        @body error: InternalServerError;
      };

    @delete
    @route("/items/{product_id}")
    op deleteItemFromCart(@path product_id: string): {
        @statusCode statusCode: 204;
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
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
