package main

import (
	"time"
)

//constant for waiting
const (
	FastResponse = time.Millisecond * 500
	SlowResponse = time.Millisecond * 1000
)

//constant for topic nsq
const (
	TopicPostOrderCreation   TopicNsq = "TopicPostPaymentCreation"
	TopicPostPaymentCreation TopicNsq = "TopicPostPaymentCreation"
)
