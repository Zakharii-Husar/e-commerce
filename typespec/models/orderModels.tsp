import "./productModels.tsp";

model OrderedItemModel extends ProductOverviewResponse{
    orderedQuantity: int16;
}

model OrderOverviewResponse {
    orderId: int64;
    orderStatus: string;
    orderPrice: float;
    createdAt: DateTimeKnownEncoding;
}

model OrderDetailsResponse extends OrderOverviewResponse {
    orderDestination: string;
    shippingRate: float;
    transactionId: string;
    itemes: OrderedItemModel[];
}

model CancelOrderResponse {
    success: boolean;
}

model UpdateOrderRequest {
    newStatus: string;
}

model UpdateOrderResponse {
    success: boolean;
}