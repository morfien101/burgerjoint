package customers

import (
	"testing"
	"time"
)

func testMenu() []string {
	return []string{"burger", "chips", "milk shake", "fried onions", "soda"}
}

func TestMakeCustomers(t *testing.T) {
	menu := testMenu()

	customerSlice := MakeCustomers(menu, 10)
	if len(customerSlice) < 10 {
		t.Fatal()
	}

	for index, customer := range customerSlice {
		if len(customer.Order()) < 1 {
			t.Logf("Customer %d does not have an order", index)
			t.Fail()
		} else {
			for _, order := range customer.Order() {
				if order == "" {
					t.Logf("Customer %d has a blank order.", index)
					t.Fail()
				}
			}
		}
	}
}

func TestCustomerWaitTimer(t *testing.T) {
	customer := MakeCustomers(testMenu(), 1)[0]
	stop := customer.StartWait()
	time.Sleep(time.Millisecond * 10)
	stop()
	if customer.TimeWaiting() > time.Millisecond*10 {
		t.Logf("Customer TimeWaiting returned a number smaller than 10ms. Got %s", customer.TimeWaiting())
		t.Fail()
	}
}

func BenchmarkMakeCustomers100(b *testing.B) {
	menu := testMenu()
	MakeCustomers(menu, 100)
}

func BenchmarkMakeCustomers1000(b *testing.B) {
	menu := testMenu()
	MakeCustomers(menu, 1000)
}
func BenchmarkMakeCustomers10000(b *testing.B) {
	menu := testMenu()
	MakeCustomers(menu, 10000)
}
