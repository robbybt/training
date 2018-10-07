package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type training struct{}

var once sync.Once
var handler *training

func GetTrainingHandler() *training {
	once.Do(func() {
		handler = new(training)
	})
	return handler
}

func (t training) Handler(w http.ResponseWriter, r *http.Request) {
	//=========Dont change this code==================
	timeStart := time.Now()
	customerIDReq := r.FormValue("customerid")
	countProduct := r.FormValue("countproduct")
	count, err := strconv.ParseInt(countProduct, 10, 64)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	productIDs := t.loadCache(count)
	customerID, err := strconv.ParseInt(customerIDReq, 10, 64)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	logging := GetLogging()
	//=========Dont change this code==================

	// What you have Now :
	// customerID
	// productIDs

	ctx := r.Context()
	//Get Customer Data
	customer, err := GetCustomerData(ctx, customerID)
	logging.Save("GetCustomerData", BuildLogDetail(customerID, customer))
	if err != nil {
		w.Write([]byte("error GetCustomerData"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	//Get ALL Product
	products, err := GetProductData(ctx, productIDs)
	logging.Save("GetProductData", BuildLogDetail(productIDs, products))
	if err != nil {
		w.Write([]byte("error GetProductData"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	//Get Shop Data
	listShop := make([]int64, 0)
	for _, p := range products {
		listShop = append(listShop, p.ShopID)
	}
	shops, err := GetShopData(ctx, listShop)
	logging.Save("GetShopData", BuildLogDetail(listShop, shops))
	if err != nil {
		w.Write([]byte("error GetShopData"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	err = t.validationCustomer(ctx, customer)
	if err != nil {
		w.Write([]byte("error validationCustomer"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}
	err = t.validationProducts(ctx, products)
	if err != nil {
		w.Write([]byte("error validationProducts"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}
	err = t.validationShop(ctx, shops)
	if err != nil {
		w.Write([]byte("error validationShop"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	// Start Transactional db
	tx, err := beginTx()
	if err != nil {
		w.Write([]byte("error beginTx"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	//insert Cart
	carts := make([]*CartData, 0)
	for _, p := range products {
		cart := BuildCartData(customer, p)
		err = InsertCart(ctx, tx, cart) // will update cart id
		if err != nil {
			w.Write([]byte("error InsertCart"))
			SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
			return
		}
		carts = append(carts, cart)
	}

	//Insert Payment
	paymentData := BuildPaymentData(customer)
	err = InsertPayment(ctx, tx, paymentData) //will update payment id
	if err != nil {
		w.Write([]byte("error InsertPayment"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}

	ordersData := BuildOrderData(customer, carts)
	for _, o := range ordersData {
		//Insert Order
		err = InsertOrder(ctx, tx, o) //will update order id
		if err != nil {
			w.Write([]byte("error InsertOrder"))
			SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
			return
		}
		//Insert PaymentDetail
		err = InsertPaymentDetail(ctx, tx, BuildPaymentDetailData(o.OrderID, paymentData.PaymentID))
		if err != nil {
			w.Write([]byte("error InsertPaymentDetail"))
			SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
			return
		}
	}

	err = commit(tx)
	if err != nil {
		w.Write([]byte("error commit"))
		SetArraylistBelowFive(&FiveLastRequestDetail, logging.Detail)
		return
	}
	// End Transactional db

	err = publish(TopicPostOrderCreation)
	if err != nil {
		fmt.Println("error publish", TopicPostOrderCreation)
	}
	err = publish(TopicPostPaymentCreation)
	if err != nil {
		fmt.Println("error publish", TopicPostPaymentCreation)
	}

	w.Write([]byte(fmt.Sprint("success", time.Since(timeStart))))
}

func (training) loadCache(count int64) (result []int64) {
	for i := int64(0); i < count; i++ {
		result = append(result, i)
	}
	return
}

func (training) validationShop(ctx context.Context, shops []*ShopData) (err error) {
	for _, s := range shops {
		fmt.Println("validationShop...", s.ShopID)
		if s.ShopID == 1999 {
			err = errors.New("validationShop error")
			return
		}
		<-time.Tick(FastResponse)
	}
	return
}

func (training) validationProducts(ctx context.Context, products []*ProductData) (err error) {
	for _, p := range products {
		fmt.Println("validationProducts...", p.ProductID)
		if p.ProductID == 1999 {
			err = errors.New("validationProducts error")
			return
		}
		<-time.Tick(FastResponse)
	}
	return
}

func (training) validationCustomer(ctx context.Context, customer *CustomerData) (err error) {
	fmt.Println("validationCustomer...")
	<-time.Tick(SlowResponse)
	if customer.CustomerID == 999 {
		err = errors.New("validationCustomer error")
	}
	return
}

//just mock begintx
func beginTx() (*sqlx.Tx, error) {
	fmt.Println("beginTx...")
	return nil, nil
}

//just mock commit
func commit(tx *sqlx.Tx) (err error) {
	fmt.Println("commit...")
	return
}

func publish(topic TopicNsq) (err error) {
	fmt.Println("publish...", topic)
	<-time.Tick(SlowResponse)
	return
}
