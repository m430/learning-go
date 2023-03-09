package decorator

import "fmt"

// 抽象
type Pizza interface {
	getFlavor() string
}

// 培根披萨
type BaconPizza struct{}

func (p *BaconPizza) getFlavor() string {
	return "bacon"
}

// 芝士口味
type Cheese struct {
	pizza Pizza
}

func (c *Cheese) getFlavor() string {
	flavor := fmt.Sprintf("%s %s", c.pizza.getFlavor(), "cheese")
	fmt.Printf("=====%s=====\n", flavor)
	return flavor
}

type Spicy struct {
	pizza Pizza
}

func (s *Spicy) getFlavor() string {
	flavor := fmt.Sprintf("%s %s", s.pizza.getFlavor(), "spicy")
	fmt.Printf("=====%s=====\n", flavor)
	return flavor
}
