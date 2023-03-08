package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	s := NewService("product")
	s.Create("Nike Shoe")
}
