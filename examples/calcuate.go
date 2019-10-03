package main

import (
	"fmt"
	"go-solvingmaze"
)

func main() {
	sm := solvingmaze.New("ABCDEFGHIJKLMNOP", "Los Angeles", 5, true)
	pack := sm.NewPackAndQuote()
	itemProps := make(map[string]interface{})
	itemProps["Name"] = "WidgetOne"
	itemProps["Sku"] = "W1S123"
	item, err := pack.AddItem(itemProps)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	dim := item.AddDimensions()
	dim.Length(25)
	dim.Width(12)
	dim.Height(22)
	dim2 := item.AddDimensions(map[string]interface{}{"Length": 12, "Width": 8, "Height": 2})
	dim2.Height(22)

	address := make(map[string]interface{})
	address["post"] = "48072"
	address["region"] = "MI"
	address["city"] = "Berkley"
	address["country"] = "US"
	address["residential"] = true

	_, err = pack.Destination(address)
	if err != nil {
		fmt.Printf("Error adding destination: %s", err)
	}

	dim = item.AddVoidDimensions(map[string]interface{}{"Length": 12, "Width": 8, "Height": 2})
	req, err := pack.Calculate()

	if err != nil {
		fmt.Printf("Failed to calculate shipping: %s", err)
	} else {
		fmt.Printf("\nRate: $%v\n", req.Services[0].Rate)
	}
}
