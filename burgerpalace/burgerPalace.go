package burgerpalace

import (
	"github.com/morfien101/burgerJoint/computer"
	"github.com/morfien101/burgerJoint/logger"
)

// This package will hold the logic for how our store "Burger Palace" will work.

// NewBurgerPalace will return a burger store ready to service customers.
func NewBurgerPalace(logger logger.Logger) *BurgerStore {
	return &BurgerStore{
		logger:   logger,
		computer: computer.NewComputer(logger),
	}
}
