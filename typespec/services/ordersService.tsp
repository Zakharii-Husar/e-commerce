import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/orderModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Orders Service",
})
@route("/api/orders")
namespace OrdersService {
    
    @get
    @route("/")
    op getOrders(): {
        @statusCode statusCode: 200;
        @body orders: OrderOverviewResponse[];
    } | {
        @statusCode statusCode: 400;
        @body error: InvalidRequestError;
    } | {
        @statusCode statusCode: 401;
        @body error: UnauthorizedError;
    };

    @get
    @route("/items/{order_id}")
    op getOrderById(@path order_id: string): {
        @statusCode statusCode: 200;
        @body order: OrderDetailsResponse;
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

    @patch
    @route("/items/{order_id}/cancel")
    op cancelOrder(@path order_id: string): {
        @statusCode statusCode: 200;
        @body response: CancelOrderResponse;
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

    @patch
    @route("/items/{order_id}/update")
    op updateOrder(@path order_id: string, @body updateOrderReq: UpdateOrderRequest): {
        @statusCode statusCode: 200;
        @body response: UpdateOrderResponse;
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
