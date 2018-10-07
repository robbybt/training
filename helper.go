package main

func BuildCartData(customer *CustomerData, product *ProductData) *CartData {
	return &CartData{
		Customer: customer,
		Product:  product,
	}
}

func BuildOrderData(customer *CustomerData, carts []*CartData) []*OrderData {
	modZeroCart := make([]int64, 0)
	modOneCart := make([]int64, 0)
	modTwoCart := make([]int64, 0)
	for _, c := range carts {
		if c.Product.ProductID%3 == 0 {
			modZeroCart = append(modZeroCart, c.CartID)
		} else if c.Product.ProductID%3 == 1 {
			modOneCart = append(modOneCart, c.CartID)
		} else {
			modTwoCart = append(modTwoCart, c.CartID)
		}
	}

	return []*OrderData{
		&OrderData{
			Customer: customer,
			CartIDs:  modZeroCart,
		},
		&OrderData{
			Customer: customer,
			CartIDs:  modOneCart,
		},
		&OrderData{
			Customer: customer,
			CartIDs:  modTwoCart,
		},
	}
}

func BuildPaymentData(customer *CustomerData) *PaymentData {
	return &PaymentData{
		Customer: customer,
	}
}

func BuildPaymentDetailData(orderID int64, paymentID int64) *PaymentDetailData {
	return &PaymentDetailData{
		OrderID:   orderID,
		PaymentID: paymentID,
	}
}

func BuildLogDetail(request interface{}, response interface{}) *LogDetail {
	return &LogDetail{
		Request:  request,
		Response: response,
	}
}
