import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "../models/authModels.tsp";
import "../models/errorModels.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

#suppress "@typespec/http/no-service-found" ""

@service({
    title: "Auth service",
  })
@route("/api/auth")
namespace ECommerce.AuthService {
    @post
    @route("/sign_up") op signUp(signUpReq: SignUpRequest): 
    {
      @body authResponse: AuthResponse;
    }
    | 
    {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    } | 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    @post
    @route("/sign_in") op signIn(signInReq: SignInRequest): 
    AuthResponse |
    {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    } | 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };


    @get
    @route("/sign_out") op signOut(): 
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
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };

    @post
    @route("/recover_password") op recoverPassword(passRecoverRequest: PassRecoverRequest): 
    {
      @statusCode statusCode: 204;
    } |
    {
      @statusCode statusCode: 400;
      @body error: InvalidRequestError;
    }| 
    {
      @statusCode statusCode: 401;
      @body error: UnauthorizedError;
    }| {
      @statusCode statusCode: 500;
      @body error: InternalServerError;
    };
} 

