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
	TopicPostOrderCreation   TopicNsq = "TopicPostOrderCreation"
	TopicPostPaymentCreation TopicNsq = "TopicPostPaymentCreation"
)

//constant for context
const (
	KeyTimeStart ContextKey = "timeStart"
)
