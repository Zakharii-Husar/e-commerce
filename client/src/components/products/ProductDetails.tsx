import Error from "../reusable/Error";
import Loading from "../reusable/Loading";

import type { paths, components } from "../../types/types";

import { useState } from "react";

type ProductDetailsType = components["schemas"]["ProductDetailsResponse"];
type InvalidRequestError = components["schemas"]["InvalidRequestError"];
type UnauthorizedError = components["schemas"]["UnauthorizedError"];
type NotFoundError = components["schemas"]["NotFoundError"];

let productDetailsType: ProductDetailsType;

const ProductDetails: React.FC = () => {

  const [error, setError] = useState<null | InvalidRequestError | UnauthorizedError | NotFoundError>(null);
  const [loading, setLoading] = useState(false);

  if (loading) return <Loading />
  if (error !== null) return <Error code={error.code} message={error.message} />

  return (
    <h1>{productDetailsType.sellerUsername}</h1>
  )
}

export default ProductDetails


