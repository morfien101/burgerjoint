package equipment

import (
	"github.com/morfien101/burgerJoint/logger"
	"github.com/morfien101/burgerJoint/menu"
)

// Type is a peice of equipment. The configuration will dictate what it is.
type Type struct {
	// StationCode will match food with equipment
	StationCode string `json:"station_code"`
	// MaxReadySlots will allow the equipment to preprep and hold food for orders
	// If equipment is over flowing a penilty on time is for preperation is added.
	MaxReadySlots float32 `json:"max_ready_slots"`
	// ConcurrentPrepSpaces is hold many spaces the equipment offers to work at the
	// same time. If grill can cook 4 burgers at a time then it would have a value
	// of 4.
	ConcurrentPrepSpaces int `json:"concurrent_preparation_spaces"`
	// Allow the equipment prep prepare food
	Preprep bool
	OrdersQ chan *menu.MenuItem
	logger  logger.Logger
}

func (e *Type) Start() {
	//
}
