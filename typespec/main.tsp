import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

import "./services/adminService.tsp";
import "./services/authService.tsp";
import "./services/ordersService.tsp";
import "./services/productsService.tsp";
import "./services/usersService.tsp"; 
import "./services/wishlistService.tsp";
import "./services/cartService.tsp";
import "./services/sellerService.tsp";


using TypeSpec.Http;
using TypeSpec.Rest;

@server("https://ecommerce.com", "Single server endpoint for the eCommerce service")
@service
@route("/")
namespace ECommerce {}