package main

type TopicNsq string

type ShopData struct {
	ShopID   int64
	ShopName string
	SellerID int64
	IsQA     bool
}

type ProductData struct {
	ProductID   int64
	ProductName string
	ShopID      int64
}

type CustomerData struct {
	CustomerID   int64
	CustomerName string
	IsQA         bool
}

type CartData struct {
	CartID   int64
	Customer *CustomerData
	Product  *ProductData
}

type PaymentData struct {
	PaymentID int64
	Customer  *CustomerData
}

type OrderData struct {
	OrderID  int64
	Customer *CustomerData
	CartIDs  []int64
}

type PaymentDetailData struct {
	PaymentID int64
	OrderID   int64
}

type LogDetail struct {
	Request  interface{}
	Response interface{}
}
