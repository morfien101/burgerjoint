package burgerpalace

import (
	"github.com/morfien101/burgerJoint/computer"
	"github.com/morfien101/burgerJoint/customers"
	"github.com/morfien101/burgerJoint/logger"
)

// BurgerStore will hold everything about our store.
type BurgerStore struct {
	logger   logger.Logger
	computer *computer.Computer
}

// Menu describes what is available at our store
func (bs *BurgerStore) Menu() []string {
	return []string{"burger", "chips", "milk shake", "fried onions", "soda"}
}

// ServeCustomers serves the customers in the passed in queue
func (bs *BurgerStore) ServeCustomers(horde []*customers.Customer, stop chan interface{}) (done chan interface{}) {
	bs.logger.Print("Opening Tills.")
	tillQ, tillStop := TillPointWorkers(bs.logger, bs.computer)
	go func() {
		<-stop
		bs.logger.Print("Closing Tills")
		for _, stopFunc := range tillStop {
			stopFunc()
		}
	}()
	bs.logger.Print("Opening Doors.")
	customerQ := bs.DoorsOpen(horde)

	done = make(chan interface{}, 1)
	go bs.nextPlease(customerQ, tillQ, done, stop)
	return done
}

func (bs *BurgerStore) nextPlease(cq <-chan customers.CustomerActions, tillq chan<- customers.CustomerActions, done chan interface{}, stop chan interface{}) {
	defer func() {
		bs.logger.Print("All customers served")
		stop <- true
		done <- true
	}()
	for {
		select {
		case customer, ok := <-cq:
			if !ok {
				bs.logger.Print("No customers in queue")
				return
			}
			bs.logger.Print("Next Customer please")
			tillq <- customer
		}
	}
}
