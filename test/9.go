package main

import "fmt"

type Bag struct {
	items []int
}

func (b *Bag) insert(id int) *Bag {
	b.items = append(b.items, id)
	return b
}

func (b *Bag) getItem() []int {
	return b.items
}

func main() {
	b := new(Bag)
	b.insert(222)
	items := b.insert(11).insert(22).getItem()
	fmt.Println(items)
}
