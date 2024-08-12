import Error from "../reusable/Error";
import Loading from "../reusable/Loading";
import { GET_PRODUCT_BY_ID } from "../../APIEndpoints";

import type { paths, components } from "../../types/types";

import { useState, useEffect } from "react";

type GetProductResponse200 = paths["/api/products/{product_id}"]["get"]["responses"]["200"]
type GetProductResponse404 = paths["/api/products/{product_id}"]["get"]["responses"]["404"]

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type NotFoundError = components["schemas"]["NotFoundError"];
type InternalServerError = components["schemas"]["InternalServerError"];

type ApiError = NotFoundError | InternalServerError | null;

interface ApiResponse<T> {
  status: number;
  data: T | null;
  error: ApiError;
}


const fetchProductById = async (id: string): Promise<ApiResponse<ProductDetailsType>> => {
  const link = GET_PRODUCT_BY_ID(id);

  const response = await fetch(link, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });


  if (response.status === 200) {
    const result: ProductDetailsType = await response.json();
    return { status: 200, data: result, error: null };
  }
  if (response.status === 404) {
    const result: NotFoundError = await response.json();
    return { status: 404, data: null, error: result };
  }

  if (response.status === 500) {
    const result: InternalServerError = await response.json();
    return { status: 500, data: null, error: result };
  }

  return await response.json();
}

let productDetailsType: ProductDetailsType;

const ProductDetails: React.FC = () => {

  const [error, setError] = useState<ApiError>(null);
  const [loading, setLoading] = useState(false);
  const [productDetails, setProductDetails] = useState<ProductDetailsType | null>(null);

  useEffect(() => {
    const getData = async () => {
      setLoading(true);
      const result = await fetchProductById("12");
      if (result.status === 200) {
        setError(null);
        setProductDetails(result.data)
      } else {
        setError(result.error);
        setProductDetails(null);
      }
      setLoading(false);
    }

    getData();

  }, [])

  if (loading) return <Loading />
  if (productDetails !== null) return (
    <h1>{productDetails.sellerUsername}</h1>
  )
  if (error !== null) return <Error code={error.code} message={error.message} />

  return <></>

}

export default ProductDetails;


