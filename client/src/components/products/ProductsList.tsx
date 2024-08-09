
import type { paths, components } from "../../types/types";
import ProductOverview from "./ProductOverview";

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type EndpointParams = paths["/api/products/{product_id}"]["parameters"];

const ProductsList: React.FC = () => {
  return (
    <h1>List Of products overview to display on home page, cart or category</h1>
  )
}

export default ProductsList