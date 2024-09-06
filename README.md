# e-commerce

## DB Initialization

1. Install PostgreSQL on your machine.

2.  Log in to PostgreSQL as a superuser: 

``psql -U postgres``

3. Create the database:

``CREATE DATABASE ecommerce;``

4. Create the user: 

``CREATE USER ecommerce WITH PASSWORD 'ecommerce_password';``

5. Grant privileges to the user on the database:

``GRANT ALL PRIVILEGES ON DATABASE ecommerce TO ecommerce;``


6. Exit the psql prompt:

``/q``

7. Initialize the database (you must be in the project root directory to run the following command):

`` psql -h localhost -U postgres -d ecommerce -f ./api/data/init_db.sql``

## Typespec

### Install dependencies:

``cd typespec && npm install``

### Generate Open API Schema with [typespec.io](https://typespec.io/)

1. Run the TypeSpec compiler to compile TypeSpec files and generate the OpenAPI schema:

``npx tsp compile ./typespec``


### Generating TypeScript types with [openapi-typescript](https://www.npmjs.com/package/openapi-typescript)


1. Provide yaml schemas and output files and generate the files with TS types(run the following command from the project root directory):

````
npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.AdminService.yaml -o ./client/src/types/adminServiceTypes.ts 
 
npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.AuthService.yaml -o ./client/src/types/authServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.CartService.yaml -o ./client/src/types/CartServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.OrdersService.yaml -o ./client/src/types/ordersServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.ProductsService.yaml -o ./client/src/types/productsServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.SellerService.yaml -o ./client/src/types/sellerServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.UsersService.yaml -o ./client/src/types/usersServiceTypes.ts

npx openapi-typescript ./typespec/tsp-output/@typespec/openapi3/openapi.ECommerce.WishlistService.yaml -o ./client/src/types/wishlistServiceTypes.ts

````

2. Import types in React project:

````
import type { paths, components } from "../../types/types.ts";

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type EndpointParams = paths["/api/products/{product_id}"]["parameters"];

````

### Generating API boilerplate code with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)

1. Navigate to "api" directory:

``cd api``

2. Generate the code:

``go generate .``