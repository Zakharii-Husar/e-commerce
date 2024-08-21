import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

import "./services/adminService.tsp";
import "./services/authService.tsp";
import "./services/ordersService.tsp";
import "./services/productsService.tsp";
import "./services/sellerService.tsp";
import "./services/usersService.tsp"; 
import "./services/wishlistService.tsp";
import "./services/cartService.tsp";


using TypeSpec.Http;
using TypeSpec.Rest;


@server("https://ecommerce.com", "Single server endpoint for the eCommerce service")
@service({
  title: "Cart Service",
})
@route("/")
namespace ECommerce {
  @service({ title: "My Combined API"})
  @route("/")
  namespace CombinedAPI {
    using AdminService;
    using AuthService;
    using CartService;
    using OrdersService;
    using ProductsService;
    using SellerService;
    using UsersService;
     using WishlistService;
  }
  
}