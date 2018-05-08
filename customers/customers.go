package customers

import (
	"math/rand"
	"time"
)

// CustomerActions holds the actions that a customer can perform
type CustomerActions interface {
	Order() []string
	StartWait() func()
	TimeWaiting() time.Duration
}

// Customer describes a customer who has an order and a time waiting in store.
type Customer struct {
	order       []string
	timeWaiting time.Duration
}

// Order will return the customers order
func (c *Customer) Order() []string {
	return c.order
}

// StartWait will start a timer and return a stop timer function
func (c *Customer) StartWait() func() {
	t1 := time.Now()
	return func() {
		c.timeWaiting = time.Since(t1)
	}
}

// TimeWaiting returns how long the customer spent waiting for something.
// This is populated by calling the StartWait function on the customer.
func (c *Customer) TimeWaiting() time.Duration {
	return c.timeWaiting
}

// MakeCustomers will return a list of customers to visit a store.
func MakeCustomers(menu []string, howMany int64) []*Customer {
	rand.Seed(time.Now().UnixNano())
	customers := make([]*Customer, howMany)
	var counter int64
	for counter < howMany {
		customers[counter] = newCustomer(menu)
		counter++
	}

	return customers
}

func newCustomer(menu []string) *Customer {
	c := new(Customer)
	c.getOrderReady(menu)
	return c
}

func (c *Customer) getOrderReady(menu []string) {
	itemsAvailable := len(menu)
	itemsToOrder := rand.Intn(5) + 1
	c.order = make([]string, itemsToOrder)
	for i := 0; i < itemsToOrder; i++ {
		item := rand.Intn(itemsAvailable)
		c.order[i] = menu[item]
	}
}
