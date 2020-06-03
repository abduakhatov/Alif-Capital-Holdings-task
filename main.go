package main

import (
	"fmt"
	"log"

	"./src"
)

func main() {
	notifyType, buyer := src.GetInput()
	msg := src.PurchaseString(*src.SamplePurchase())
	fmt.Println(msg)

	result, err := buyer.Notify(msg, notifyType)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

}
