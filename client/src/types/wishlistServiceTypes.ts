/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/api/wishlist/": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getWishlistItems"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/api/wishlist/items/{product_id}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post: operations["addItemToWishlist"];
        delete: operations["deleteItemFromWishlist"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
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
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
    getWishlistItems: {
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
                    "application/json": components["schemas"]["ProductOverviewResponse"][];
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
    addItemToWishlist: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                product_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description The request has succeeded and a new resource has been created as a result. */
            201: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
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
    deleteItemFromWishlist: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                product_id: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description There is no content to send for this request, but the headers may be useful.  */
            204: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
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
