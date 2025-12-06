// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p Product) Info() string {

	return fmt.Sprintf("%s: %.2f", p.Name, p.Price)

}
func (p *Product) UpdatePrice(newPrice float64) {

	p.Price = newPrice

}
func (p *Product) ApplyDiscount(percent float64) {

	p.Price = p.Price * (1 - percent/100)
}

func main() {

	p := Product{ID: 1, Name: "Laptop", Price: 1000}
	fmt.Println(p.Info())
	p.UpdatePrice(800)
	fmt.Println(p.Price)
	p.ApplyDiscount(10)
	fmt.Println(p.Price)
	fmt.Println(p.Info())

}
