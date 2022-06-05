package main

import "fmt"

type Weight int

func (w Weight) isHeavy() bool {
	return w > 40
}

type Item struct {
	title  string
	price  int
	weight Weight
}

func (i *Item) newTitle(new_title string) {
	i.title = new_title
}

func initItem(title string, price, weight int) Item {
	// эта функция нужна для инициализации типа с приватными полями
	return Item{
		title:  title,
		price:  price,
		weight: Weight(weight),
	}
}

func WorkingStructure() {
	itemToy := initItem("ball", 2000, 10)      // function for init instance
	itemBox := Item{"boxing gloves", 10000, 2} // init instance
	fmt.Printf("%+v\n", itemToy)
	fmt.Printf("%+v\n", itemBox)
	itemToy.newTitle("basketball ball")
	fmt.Printf("new title %+v\n", itemToy)

	itemToy.weight = Weight(50)
	fmt.Println("is item heavy?: ", itemToy.weight.isHeavy())

}
