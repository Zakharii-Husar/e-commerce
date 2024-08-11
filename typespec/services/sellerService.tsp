import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/productModels.tsp";
import "../models/orderModels.tsp";
import "../models/sellerModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Seller Service",
})
@route("/api/sellers")
namespace SellerService {
    
    @get
    @route("/inventory")
    op getInventory(): {
        @body inventory: ProductDetailsResponse[];
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    };

    @get
    @route("/orders")
    op getOrders(): {
        @body orders: OrderDetailsResponse[];
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    };

    @put
    @route("/orders/{order_id}/update")
    op updateOrder(@path order_id: string, @body updateOrderReq: UpdateOrder): {
        @body response: UpdateOrder;
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    } | {
        @statusCode statusCode: 404;
        @body error: NotFoundError;
    };
}
