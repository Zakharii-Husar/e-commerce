/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/api/orders/": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getOrders"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/orders/items/{order_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getOrderById"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/orders/items/{order_id}/cancel": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch: operations["cancelOrder"];
        trace?: never;
    };
    "/api/orders/items/{order_id}/update": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch: operations["updateOrder"];
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        CancelOrderResponse: {
            success: boolean;
        };
        /**
         * @description Known encoding to use on utcDateTime or offsetDateTime
         * @enum {string}
         */
        DateTimeKnownEncoding: "rfc3339" | "rfc7231" | "unixTimestamp";
        ErrorResponse: {
            /** Format: int32 */
            code: number;
            message: string;
        };
        InternalServerError: {
            /** @enum {number} */
            code: 500;
            /** @enum {string} */
            message: "Something went wrong. Try again later";
        } & WithRequired<components["schemas"]["ErrorResponse"], "code" | "message">;
        InvalidRequestError: {
            /** @enum {number} */
            code: 400;
            /** @enum {string} */
            message: "Invalid Request";
        } & WithRequired<components["schemas"]["ErrorResponse"], "code" | "message">;
        NotFoundError: {
            /** @enum {number} */
            code: 404;
            /** @enum {string} */
            message: "Not Found";
        } & WithRequired<components["schemas"]["ErrorResponse"], "code" | "message">;
        OrderDetailsResponse: {
            orderDestination: string;
            shippingRate: number;
            transactionId: string;
            itemes: components["schemas"]["OrderedItemModel"][];
        } & components["schemas"]["OrderOverviewResponse"];
        OrderOverviewResponse: {
            /** Format: int64 */
            orderId: number;
            orderStatus: string;
            orderPrice: number;
            createdAt: components["schemas"]["DateTimeKnownEncoding"];
        };
        OrderedItemModel: {
            /** Format: int16 */
            orderedQuantity: number;
        } & components["schemas"]["ProductOverviewResponse"];
        ProductDetailsResponse: {
            /** Format: int32 */
            stockQuantity: number;
            productDescription?: string;
            productMaterial?: string;
            productWeight?: number;
            productWidth?: number;
            productLenght?: number;
            productHeight?: number;
        } & components["schemas"]["ProductOverviewResponse"];
        ProductOverviewResponse: {
            /** Format: int64 */
            productId: number;
            productName: string;
            productPrice: number;
            productCategory: string;
            productAttachements: string[];
            sellerUsername: string;
            rating: components["schemas"]["ProductRatingResponse"];
        };
        ProductRatingResponse: {
            /** Format: int32 */
            ratedTimes: number;
            averageRating: number;
        };
        UnauthorizedError: {
            /** @enum {number} */
            code: 401;
            /** @enum {string} */
            message: "Unauthorized";
        } & WithRequired<components["schemas"]["ErrorResponse"], "code" | "message">;
        UpdateOrderRequest: {
            newStatus?: string;
        };
        UpdateOrderResponse: {
            success: boolean;
        };
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
    getOrders: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description The request has succeeded. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OrderOverviewResponse"][];
                };
            };
            /** @description The server could not understand the request due to invalid syntax. */
            400: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InvalidRequestError"];
                };
            };
            /** @description Access is unauthorized. */
            401: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UnauthorizedError"];
                };
            };
            /** @description Server error */
            500: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InternalServerError"];
                };
            };
        };
    };
    getOrderById: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                order_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description The request has succeeded. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OrderDetailsResponse"];
                };
            };
            /** @description The server could not understand the request due to invalid syntax. */
            400: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InvalidRequestError"];
                };
            };
            /** @description Access is unauthorized. */
            401: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UnauthorizedError"];
                };
            };
            /** @description The server cannot find the requested resource. */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NotFoundError"];
                };
            };
            /** @description Server error */
            500: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InternalServerError"];
                };
            };
        };
    };
    cancelOrder: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                order_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description The request has succeeded. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CancelOrderResponse"];
                };
            };
            /** @description The server could not understand the request due to invalid syntax. */
            400: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InvalidRequestError"];
                };
            };
            /** @description Access is unauthorized. */
            401: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UnauthorizedError"];
                };
            };
            /** @description The server cannot find the requested resource. */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NotFoundError"];
                };
            };
            /** @description Server error */
            500: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InternalServerError"];
                };
            };
        };
    };
    updateOrder: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                order_id: string;
            };
            cookie?: never;
        };
        requestBody: {
            content: {
                "application/json": components["schemas"]["UpdateOrderRequest"];
            };
        };
        responses: {
            /** @description The request has succeeded. */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UpdateOrderResponse"];
                };
            };
            /** @description The server could not understand the request due to invalid syntax. */
            400: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InvalidRequestError"];
                };
            };
            /** @description Access is unauthorized. */
            401: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["UnauthorizedError"];
                };
            };
            /** @description The server cannot find the requested resource. */
            404: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["NotFoundError"];
                };
            };
            /** @description Server error */
            500: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["InternalServerError"];
                };
            };
        };
    };
}
type WithRequired<T, K extends keyof T> = T & {
    [P in K]-?: T[P];
};
