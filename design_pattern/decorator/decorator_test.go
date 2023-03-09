package decorator

import (
	"learning-go/utils"
	"testing"
)

func TestOOPDecorator(t *testing.T) {
	pizza := &BaconPizza{}

	// 添加芝士口味
	pizzaWithCheese := &Cheese{pizza: pizza}
	// 添加辣味
	pizzaWithCheeseAndSpicy := &Spicy{pizza: pizzaWithCheese}

	if pizzaWithCheeseAndSpicy.getFlavor() != "bacon cheese spicy" {
		t.Errorf("pizza decorator failed")
	}
}

func TestFuncDecorator(t *testing.T) {
	s := "Hello, World"

	t.Run("func decorator", func(t *testing.T) {
		s = ToMD5(ToUpper(Just))(s)
		if !utils.IsUpper(s) || len(s) != 32 {
			t.Errorf("func decorator is failed")
		}
	})

	t.Run("func decorator with pipline", func(t *testing.T) {
		fn := Pipline(Just, ToMD5, ToUpper)
		s := fn(s)
		if !utils.IsUpper(s) || len(s) != 32 {
			t.Errorf("func decorator with pipline is failed")
		}
	})
}
