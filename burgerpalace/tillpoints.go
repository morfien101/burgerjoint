package burgerpalace

import (
	"strings"

	"github.com/morfien101/burgerJoint/computer"
	"github.com/morfien101/burgerJoint/logger"

	"github.com/morfien101/burgerJoint/customers"
)

const (
	// TillWorkers is how many open tills we have
	TillWorkers = 4
	// TillWorkerMaxWip is how many orders they can be working on at any given time.
	TillWorkerMaxWip = 3
)

// TillPointWorkers are the first point of contact for our customers.
// The workers will return a queue that can be used to feed in customers. Like that
// of a queue in front of a store.
// The stop functions can be used to stop the workers. They will not process anymore orders
// once the stop funcs have been called. Think of this as light weigh context.
func TillPointWorkers(logger logger.Logger, computer *computer.Computer) (chan<- customers.CustomerActions, []func()) {
	customerQueue := make(chan customers.CustomerActions, TillWorkers)
	stopFuncs := make([]func(), TillWorkers)
	for i := 1; i <= TillWorkers; i++ {
		logger.Printf("Starting Till worker %d\n", i)
		worker := newTillWorker(logger, i, computer, customerQueue)
		stopFuncs[i-1] = worker.start()
	}
	return customerQueue, stopFuncs
}

type tillWorker struct {
	customerQueue <-chan customers.CustomerActions
	finished      bool
	logger        logger.Logger
	id            int
	currentOrders map[int64]*computer.Order
	maxWIP        int
	computer      *computer.Computer
}

func newTillWorker(
	logger logger.Logger,
	id int,
	tillcomputer *computer.Computer,
	tillQueue <-chan customers.CustomerActions) *tillWorker {
	return &tillWorker{
		id:            id,
		computer:      tillcomputer,
		customerQueue: tillQueue,
		logger:        logger,
		currentOrders: make(map[int64]*computer.Order),
	}
}

func (tw *tillWorker) start() func() {
	tw.finished = false
	go tw.work()
	return func() {
		tw.finished = true
	}
}

func (tw *tillWorker) work() {
	tw.logger.Printf("tillWorker[%d]: Starting work.\n", tw.id)
	finished := func() {
		tw.logger.Print("worker is finished")
	}
	defer finished()

	customersAvailble := true
	for {
		if !customersAvailble && len(tw.currentOrders) == 0 {
			tw.logger.Printf("tillWorker[%d]: has no more work to complete\n", tw.id)
			tw.finished = true
		}
		// This is used to signal that the till worker is complete or should abandon
		// their post.
		if tw.finished {
			return
		}

		// Check if any order is ready
		tw.logger.Printf("tillWorker[%d] check state of current orders\n", tw.id)
		for _, order := range tw.currentOrders {
			select {
			case <-order.Ready:
				close(order.Ready)
				delete(tw.currentOrders, order.Ticket)
				tw.serveFood(order)
			default:
				continue
			}
		}

		tw.logger.Printf("tillWorker[%d]: Waiting for order\n", tw.id)
		select {
		case customerOrder, ok := <-tw.customerQueue:
			if !ok {
				customersAvailble = false
				continue
			}
			tw.logger.Printf(
				"tillWorker[%d]: processing %s\n",
				tw.id,
				strings.Join(customerOrder.Order(), ", "),
			)
			o := tw.inputOrder(customerOrder.Order())
			o.FinishedFunc = customerOrder.StartWait()
			tw.currentOrders[o.Ticket] = o
		}
	}
}

func (tw *tillWorker) inputOrder(items []string) *computer.Order {
	o := computer.NewOrder(tw.computer, items)
	tw.computer.AcceptOrder(o)
	return o
}

func (tw *tillWorker) serveFood(o *computer.Order) {
	o.FinishedFunc()
	return
}
