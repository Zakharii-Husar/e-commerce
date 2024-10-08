import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/errorModels.tsp";
import "../models/userModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Users Service",
})
@route("/api/users")
namespace ECommerce.UsersService {
    
    @get
    @route("/")
    op getUsers(): {
        @body users: UserDetailsResponse[];
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

    @get
    @route("/search")
    op searchUsers(@query query: string): {
        @body users: UserDetailsResponse[];
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

    @get
    @route("/{user_id}")
    op getUserById(@path user_id: string): {
        @body user: UserDetailsResponse;
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

    @patch
    @route("/{user_id}")
    op updateUserDetails(@path user_id: string, @body update: UpdateUserRequest): {
        @body user: UserDetailsResponse;
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
    @route("/{user_id}")
    op deleteUser(@path user_id: string): {
        @statusCode statusCode: 204;
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
}
