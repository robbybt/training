package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetProductData(ctx context.Context, requestAPI []int64) ([]*ProductData, error) {
	responseAPI := make([]*ProductData, 0)
	fmt.Println("GetProductData...")
	//APICALL start
	for _, id := range requestAPI {
		temp := new(ProductData)
		<-time.Tick(FastResponse)
		temp.ProductName = fmt.Sprint("productname ", id)
		temp.ProductID = id
		temp.ShopID = 1000 + id
		responseAPI = append(responseAPI, temp)
	}
	//APICALL end
	return responseAPI, nil
}

func GetShopData(ctx context.Context, requestAPI []int64) ([]*ShopData, error) {
	result := make([]*ShopData, 0)
	fmt.Println("GetShopData...")
	//APICALL start
	for _, id := range requestAPI {
		temp := new(ShopData)
		<-time.Tick(FastResponse)
		temp.ShopName = fmt.Sprint("shopname ", id)
		temp.SellerID = 1000 + id
	}
	//APICALL end
	return result, nil
}

func GetCustomerData(ctx context.Context, requestAPI int64) (*CustomerData, error) {
	fmt.Println("GetCustomerData...")
	//APICALL start
	responseAPI := new(CustomerData)
	<-time.Tick(SlowResponse)
	responseAPI.CustomerID = requestAPI
	responseAPI.CustomerName = fmt.Sprint("CustomerName ", requestAPI)
	responseAPI.IsQA = false
	//APICALL end
	return responseAPI, nil
}

func InsertCart(ctx context.Context, tx *sqlx.Tx, cart *CartData) (err error) {
	fmt.Println("InsertCart...", cart.Product.ProductID)
	<-time.Tick(SlowResponse)
	cart.CartID = cart.Customer.CustomerID + cart.Product.ProductID
	return
}

func InsertPayment(ctx context.Context, tx *sqlx.Tx, payment *PaymentData) (err error) {
	fmt.Println("InsertPayment...")
	<-time.Tick(SlowResponse)
	payment.PaymentID = payment.Customer.CustomerID + 1000
	return
}

func InsertOrder(ctx context.Context, tx *sqlx.Tx, order *OrderData) (err error) {
	fmt.Println("InsertOrder...")
	<-time.Tick(SlowResponse)
	order.OrderID = order.Customer.CustomerID + 10000
	return
}

func InsertPaymentDetail(ctx context.Context, tx *sqlx.Tx, paymentDetail *PaymentDetailData) (err error) {
	fmt.Println("InsertPaymentDetail...")
	<-time.Tick(SlowResponse)
	return
}
