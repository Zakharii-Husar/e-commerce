
import type { paths, components } from "../../types/types";

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type EndpointParams = paths["/api/products/{product_id}"]["parameters"];

const ProductDetails: React.FC = () => {
  return (
    <h1>Product Details</h1>
  )
}

export default ProductDetails


