package menu

import "testing"

func TestNewMenu(t *testing.T) {
	items := []struct {
		itemName        string
		itemPrepTime    int
		itemStationCode string
		itemSlotSize    float32
	}{
		{
			itemName:        "chips",
			itemPrepTime:    10,
			itemStationCode: "fryer",
			itemSlotSize:    0.2,
		},
		{
			itemName:        "burger",
			itemPrepTime:    20,
			itemStationCode: "grill",
			itemSlotSize:    1,
		},
	}

	menu := NewMenu()
	for _, item := range items {
		menu.AddItem(NewMenuItem(item.itemName, item.itemStationCode, item.itemPrepTime))
	}

	for _, i := range items {
		found := false
		match := false
		for _, mi := range menu.Items() {
			if mi.String() == i.itemName {
				found = true
				if mi.GetPrepTime() == i.itemPrepTime && mi.GetStationCode() == i.itemStationCode {
					match = true
				}
			}
		}
		if !found {
			t.Fail()
			t.Logf("Could not find %s in menu items.", i.itemName)
			if !match {
				t.Logf("The items attributes did not match")
			}
		}
	}

	if len(menu.Items()) != len(items) {
		t.Logf("Menu items are missing. Have %d, want %d.", len(menu.Items()), len(items))
		t.Fail()
	}
}
