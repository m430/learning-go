package observer

import "fmt"

type Item struct {
	Subject
	name    string
	inStock bool
}

func (i *Item) notifyAll() {
	for _, o := range i.observers {
		o.update(i)
	}
}

func NewItem(name string) *Item {
	return &Item{name: name}
}

func (i *Item) updateStock() {
	fmt.Printf("Item %s now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
