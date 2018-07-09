package computer

import (
	"sync/atomic"
	"time"

	"github.com/morfien101/burgerJoint/logger"
)

// Computer is used to route orders and track them
type Computer struct {
	logger logger.Logger
	ticket int64
}

// NewComputer returns a computer
func NewComputer(logger logger.Logger) *Computer {
	return &Computer{
		logger: logger,
	}
}

func (c *Computer) nextTicket() int64 {
	return atomic.AddInt64(&c.ticket, 1)
}

// AcceptOrder will instruct the computer to send the orders to the machine workers.
func (c *Computer) AcceptOrder(o *Order) {
	// Currently we just sleep but the idea is that we will send a request for cooked food.
	time.Sleep(time.Millisecond * 10)
	o.Ready <- true
}

// Order is what the computer produced when the till staff input an order.
// Orders are distributed to machine workers.
type Order struct {
	Ticket       int64
	Items        []string
	Ready        chan interface{}
	FinishedFunc func()
}

// NewOrder will return a new order from an instances of computer.
func NewOrder(computerInstance *Computer, items []string) *Order {
	return &Order{
		Ticket: computerInstance.nextTicket(),
		Items:  items,
		Ready:  make(chan interface{}, 1),
	}
}
