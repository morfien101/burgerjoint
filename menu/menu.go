package menu

import (
	"sync"
)

// MenuItem is a food item that is listed on the menu. It will also have details
// about how long it takes to prep and who preps it.
type MenuItem struct {
	sync.RWMutex
	Name        string
	PrepTime    int
	StationCode string
	SlotSize    float32
}

func (mi *MenuItem) String() string {
	mi.RLock()
	defer mi.RUnlock()
	return mi.Name
}

// GetPrepTime returns the number of unit of time taken to prep this item.
func (mi *MenuItem) GetPrepTime() int {
	mi.RLock()
	defer mi.RUnlock()
	return mi.PrepTime
}

// GetStationCode will return the code that specifies which stations should prepare this
// item.
func (mi *MenuItem) GetStationCode() string {
	mi.RLock()
	defer mi.RUnlock()
	return mi.StationCode
}

// GetSlotSize will return how many units this item will take up in the equipment
func (mi *MenuItem) GetSlotSize() float32 {
	mi.RLock()
	defer mi.RUnlock()
	return mi.SlotSize
}

// Menu holds a list of menu items that customers can pick from
type Menu struct {
	sync.RWMutex
	items []*MenuItem
}

// NewMenuItem will create a menu item and return a pointer to it.
func NewMenuItem(name, stationcode string, preptime int) *MenuItem {
	return &MenuItem{
		Name:        name,
		PrepTime:    preptime,
		StationCode: stationcode,
	}
}

// NewMenu is a menu that can be presented to customers.
func NewMenu() *Menu {
	return &Menu{
		items: make([]*MenuItem, 0),
	}
}

// AddItem will push a item into the list of items available on our menu
func (m *Menu) AddItem(item *MenuItem) {
	m.Lock()
	defer m.Unlock()
	m.items = append(m.items, item)
}

// Items will return the list of pointers to items that can be made.
func (m *Menu) Items() []*MenuItem {
	m.RLock()
	defer m.RUnlock()
	return m.items
}
