package main

import (
	"time"
)

//constant for waiting
const (
	FastResponse = time.Millisecond * 50
	SlowResponse = time.Millisecond * 100
)

//constant for topic nsq
const (
	TopicPostOrderCreation   TopicNsq = "TopicPostPaymentCreation"
	TopicPostPaymentCreation TopicNsq = "TopicPostPaymentCreation"
)
