package src

import "fmt"

// Client info
type Buyer struct {
	ID    string
	phone string
	email string
}

// Item in basket
type Purchase struct {
	ID      int
	product string
	price   float32
}

// Purchase list of Items
// type Purchase []Item

// SamplePurchase returns sample Purchase fiiled with Items
func SamplePurchase() (p *[]Purchase) {
	p = &[]Purchase{
		Purchase{1, "product1", 99.99},
		Purchase{2, "product2", 99.99},
		Purchase{3, "product3", 99.99},
		Purchase{4, "product4", 99.99},
		Purchase{5, "product5", 99.99},
		Purchase{6, "product6", 99.99},
		Purchase{7, "product7", 99.99},
		Purchase{8, "product8", 99.99},
		Purchase{9, "product9", 99.99},
		Purchase{10, "product10", 99.99},
	}
	return p
}

func PurchaseString(purchases []Purchase) (result string) {
	var total float32 = 0
	for _, val := range purchases {
		result += fmt.Sprintf("%v-%.2f\n", val.product, val.price)
		total += val.price
	}
	result += fmt.Sprintf("Total:%.2f", total)
	return result
}
