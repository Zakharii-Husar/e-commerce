import Error from "../../reusable/Error";
import Loading from "../../reusable/Loading";
import { GET_PRODUCT_BY_ID } from "../../../APIEndpoints";

import type { paths, components } from "../../../types/types";

import { useState, useEffect } from "react";
import styled from "styled-components";
import productDetailsData from "../../../mock_data/productDetailsData";

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
  // const link = GET_PRODUCT_BY_ID(id);

  // const response = await fetch(link, {
  //   method: 'GET',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  // });


  // if (response.status === 200) {
  //   const result: ProductDetailsType = await response.json();
  //   return { status: 200, data: result, error: null };
  // }
  // if (response.status === 404) {
  //   const result: NotFoundError = await response.json();
  //   return { status: 404, data: null, error: result };
  // }

  // if (response.status === 500) {
  //   const result: InternalServerError = await response.json();
  //   return { status: 500, data: null, error: result };
  // }

  return { status: 200, data: productDetailsData, error: null };
}


const ProductDetails: React.FC = () => {

  const [error, setError] = useState<ApiError>(null);
  const [loading, setLoading] = useState(false);
  const [productDetails, setProductDetails] = useState<ProductDetailsType | null>(null);

  useEffect(() => {
    const getData = async () => {
      setLoading(true);
      const result = await fetchProductById("12");
      // if (result.status === 200) {
      //   setError(null);
      //   setProductDetails(result.data)
      // } else {
      //   setError(result.error);
      //   setProductDetails(null);
      // }
      
      setLoading(false);

      setProductDetails(result.data)
    }

    getData();

  }, [])

  if (loading) return <Loading />
  if (error !== null) return <Error code={error.code} message={error.message} />


    return (
      <Container>
        <ProductHeader>
          {productDetails?.productName}
        </ProductHeader>
        <div>Category: {productDetails?.productCategory}</div>
        <div>Description: {productDetails?.productDescription ?? "-"}</div>
        <div>Height: {productDetails?.productHeight ?? "-"}</div>
        <div>Material: {productDetails?.productMaterial ?? "-"}</div>
        <div>Price: {productDetails?.productPrice}</div>
        <div>Weigth: {productDetails?.productWeight ?? "-"}</div>
        <div>Width: {productDetails?.productWidth ?? "-"}</div>
        <div>Seller Username: {productDetails?.sellerUsername}</div>
        <div>Quantuty available: {productDetails?.stockQuantity}</div>
      </Container>
    )

}

export default ProductDetails;

const Container = styled.div`
  display: flex;
  width: 100%;
`

const ProductHeader = styled.h1`
  display: flex;
  width: 100%;
`


