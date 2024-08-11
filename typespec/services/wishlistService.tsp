import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/userModels.tsp";
import "../models/productModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Wishlist Service",
})
@route("/api/wishlist")
namespace WishlistService {
    
    @get
    @route("/")
    op getWishlistItems(): {
        @body wishlist: ProductOverviewResponse[];
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 500;
        @body error: InternalServerError;
      };

    @post
    @route("/items/{product_id}")
    op addItemToWishlist(@path product_id: string): {
        @statusCode statusCode: 201;
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    } | {
        @statusCode statusCode: 500;
        @body error: InternalServerError;
      };

    @delete
    @route("/items/{product_id}")
    op deleteItemFromWishlist(@path product_id: string): {
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
    } | {
        @statusCode statusCode: 500;
        @body error: InternalServerError;
      };
}
