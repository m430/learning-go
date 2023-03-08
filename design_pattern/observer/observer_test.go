package observer

import "testing"

func TestObserver(t *testing.T) {
	// nike商品
	shirtItem := NewItem("Nike Shirt")

	// 客户
	c1 := NewCustomer("11@qq.com")
	c2 := NewCustomer("22@qq.com")

	// 两个客户订阅了商品
	shirtItem.addObserver(c1)
	shirtItem.addObserver(c2)

	// 当商品上架新的库存后，会自动通知客户c1和c2
	shirtItem.updateStock()
}
