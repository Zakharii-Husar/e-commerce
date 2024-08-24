/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/api/admin/suspend/product/{product_id}": {
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
        patch: operations["suspendProduct"];
        trace?: never;
    };
    "/api/admin/suspend/user/{user_id}": {
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
        patch: operations["suspendUser"];
        trace?: never;
    };
    "/api/admin/transactions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["getAllTransactions"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
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
        TransactionResponse: {
            /** Format: int64 */
            transactionId: number;
            transactionType: string;
            paymentMethod: string;
            /** Format: int64 */
            orderId: number;
            /** Format: int64 */
            payerId: number;
            /** Format: int64 */
            payeeId: number;
            payerUsername: string;
            payeeUsername: string;
            amount: number;
            transactionStatus: string;
            transactionTime: components["schemas"]["DateTimeKnownEncoding"];
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
    suspendProduct: {
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
    suspendUser: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                user_id: string;
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
    getAllTransactions: {
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
                    "application/json": components["schemas"]["TransactionResponse"][];
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
