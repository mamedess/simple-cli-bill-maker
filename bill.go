package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"torta": 10.22, "bolo": 14.55},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	i := 0
	fs := "Bill breakdown: \n\n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		i++
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)

		if i == len(b.items) {
			fs += fmt.Sprintf("\n")
		}

		total += v
	}

	//line
	fs += fmt.Sprintf("------------------------------------ \n\n")

	//tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

// update top
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// handle prompt input "a"
func (b *bill) handlePromptAdd(r *bufio.Reader) {
	name, _ := getInput("Item name: ", r)
	price, _ := getInput("Item price: ", r)

	p, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("The price must be a number")
		prompOptions(*b)
	}

	b.addItem(name, p)

	fmt.Println("Item added - ", name, price)
	prompOptions(*b)
}

// handle prompt input "t"
func (b *bill) handlePromptAddTip(r *bufio.Reader) {
	tip, _ := getInput("Enter tip amount: ", r)

	t, err := strconv.ParseFloat(tip, 64)

	if err != nil {
		fmt.Println("The tip must be a number")
		prompOptions(*b)
	}

	b.updateTip(t)
	fmt.Println("tip added - ", tip)
	prompOptions(*b)
}

// handle prompt input "s"
func (b *bill) handlePromptSave() {
	data := []byte(b.format())
	filename := "bills/" + b.name + ".txt"

	err := os.WriteFile(filename, data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("you saved the file - ", b.name)
}
