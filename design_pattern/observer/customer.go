package observer

import "fmt"

type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{id: id}
}

func (c *Customer) getID() string {
	return c.id
}

func (c *Customer) update(data any) {
	item := data.(*Item)
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, item.name)
}
