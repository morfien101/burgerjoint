package burgerpalace

import (
	"github.com/morfien101/burgerJoint/customers"
)

const (
	// MaxCustomersInQueue is the capacity of our shop front
	MaxCustomersInQueue = 100
)

// DoorsOpen will take a slice of customers and feed them into a queue. Like a door to a store.
func (bs *BurgerStore) DoorsOpen(horde []*customers.Customer) <-chan customers.CustomerActions {
	customerQ := make(chan customers.CustomerActions, MaxCustomersInQueue)
	go func() {
		bs.logger.Print("Forming a queue inside.")
		for _, customer := range horde {
			customerQ <- customer
		}
		close(customerQ)
	}()
	return customerQ
}
