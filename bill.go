package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func makeNewBill(name string) bill {

	bill := bill{
		name:  name,
		items: map[string]float64{"pie": 5.999, "cake": 3.99, "pizza": 10.99},
		tip:   0,
	}

	return bill
}

// format bill

func (b bill) format() string {

	fs := "Bill breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {

		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)

	return fs
}
