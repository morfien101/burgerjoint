package main

import (
	"testing"

	"github.com/morfien101/burgerJoint/burgerpalace"
	"github.com/morfien101/burgerJoint/customers"
	"github.com/morfien101/burgerJoint/logger"
)

func TestBurgerPalace(t *testing.T) {
	logger := logger.NewLogger()
	logger.Toggle(true)
	bs := burgerpalace.NewBurgerPalace(logger)
	customers := customers.MakeCustomers(bs.Menu(), 1000)
	done := make(chan interface{})
	go bs.ServeCustomers(customers, done)
	<-done
	<-logger.Drain()
}
