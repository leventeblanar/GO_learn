package main

import (
	"fmt"
	"errors"
)

type Item struct {
	name		string
	price		float64
	quantity	int
}

type ShoppingCart struct {
	items		[]Item
	discount	float64
}

// AddItem(item Item) - termék hozzáadása
func (s *ShoppingCart) AddItem(item Item) {
	s.items = append(s.items, item)
	fmt.Printf("%s hozzáadva a kosárhoz.\n", item.name)
} 

// RemoveItem(name string) error - termék törlése név alapján (error ha nincs meg)
func (s *ShoppingCart) RemoveItem(name string) error {
	for i, element := range s.items {
		if element.name == name {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return nil
		}
	}
	return errors.New("a termék nem található")
}

// UpdateQuantity(name string, newQuantity int) error - mennyiség módosítása (error ha nincs meg vagy negatív)
func (s *ShoppingCart) UpdateQuantity(name string, newQuantity int) error {
	if newQuantity <= 0 {
		return errors.New("a mennyiség nem lehet 0 vagy negatív")
	}
	for i := range s.items {
		if s.items[i].name == name {
			s.items[i].quantity = newQuantity
			return nil
		}
	}
	return errors.New("a termék nem található")
}
// ApplyDiscount(percent float64) error - kedvezmény alkalmazása (error ha nem 0-100 között)
func (s *ShoppingCart) ApplyDiscount(percent float64) error {
	if percent <= 0 || percent > 100 {
		return errors.New("a kedvezmény mennyisége 0-100 között kell legyen")
	}
	s.discount = percent
	return nil
}

// GetTotal() float64 - végösszeg kedvezménnyel
func (s ShoppingCart) GetTotal() float64 {
	var totalPrice float64
	for i := range s.items {
		totalPrice = totalPrice + (s.items[i].price * float64(s.items[i].quantity)) 
	}
	totalPrice = totalPrice * (1 - s.discount/100)
	return totalPrice
}
// GetItemCount() int - hány féle termék van
func (s ShoppingCart) GetItemCount() int {
	itemCount := len(s.items)
	return itemCount
}

// ListItems() - kilistázza az összes terméket
func (s ShoppingCart) ListItems() {
	for i, item := range s.items {
		fmt.Printf("%d. %s - %.2f$ - %ddb\n", i+1, item.name, item.price, item.quantity)
	}
}

func main() {
	var shoppingCart1 ShoppingCart

	shoppingCart1.AddItem(Item{name: "alma", price: 3, quantity: 6})
	shoppingCart1.AddItem(Item{name: "kenyér", price: 2.5, quantity: 2})
	shoppingCart1.AddItem(Item{name: "szalonna", price: 40, quantity: 1})

	shoppingCart1.ListItems()

	if err := shoppingCart1.UpdateQuantity("alma", 10); err != nil {
		fmt.Println("Mennyiség update Hiba: ", err)
	} else {
		fmt.Println("Mennyiség sikeresen updatelve")
	}

	if err := shoppingCart1.ApplyDiscount(15); err != nil {
		fmt.Println("Kedvezmény hiba: ", err)
	} else {
		fmt.Println("Kedvezmény hozzáadva.")
	}

	fmt.Printf("Végösszeg: %.2f$\n", shoppingCart1.GetTotal())

	if err := shoppingCart1.RemoveItem("körte"); err != nil {
		fmt.Println("Remove hiba: ", err)
	} else {
		fmt.Println("Sikeresen eltávolítva.")
	}

	
	if err := shoppingCart1.ApplyDiscount(150); err != nil {
		fmt.Println("Kedvezmény hiba: ", err)
	} else {
		fmt.Println("Kedvezmény hozzáadva.")
	}
}