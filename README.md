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

1. Edit tspconfig.yaml file to change API schema format and output directory( ``ecommers/typespec/tsp-output/@typespec/openapi3/openapi.yaml`` by default).

2. Run the TypeSpec compiler to compile TypeSpec files and generate the OpenAPI schema:

``npx tsp compile .``

### Generating TypeScript types with [openapi-typescript](https://www.npmjs.com/package/openapi-typescript)

1. Provide yaml schema and output file and generate the file with TS types(run the following command from the project root directory or adjust path accordingly):

``npx openapi-typescript typespec/tsp-output/@typespec/openapi3/openapi.yaml -o ./client/src/types/types.ts``

2. Import types in React project:

````
import type { paths, components } from "../../types/types.ts";

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type EndpointParams = paths["/api/products/{product_id}"]["parameters"];

````
