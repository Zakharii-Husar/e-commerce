import Error from "../reusable/Error";
import Loading from "../reusable/Loading";
import { GET_PRODUCT_BY_ID } from "../../APIEndpoints";

import type { paths, components } from "../../types/types";

import { useState } from "react";

type GetProductResponse200 = paths["/api/products/{product_id}"]["get"]["responses"]["200"]
type GetProductResponse404 = paths["/api/products/{product_id}"]["get"]["responses"]["404"]

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type NotFoundError = components["schemas"]["NotFoundError"];

const fetchProductById = async (id: string): Promise<ProductDetailsType | NotFoundError> => {
  const link = GET_PRODUCT_BY_ID(id);

  const response = await fetch(link, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });


  if (response.status == 200) {
    const data: ProductDetailsType = await response.json();
    return data;
  }
  if (response.status == 404) {
    const data: NotFoundError = await response.json();
    return data;
  }
  
  return await response.json();

}

let productDetailsType: ProductDetailsType;

const ProductDetails: React.FC = () => {

  const [error, setError] = useState<null | NotFoundError>(null);
  const [loading, setLoading] = useState(false);

  if (loading) return <Loading />
  if (error !== null) return <Error code={error.code} message={error.message} />

  return (
    <h1>{productDetailsType.sellerUsername}</h1>
  )
}

export default ProductDetails


