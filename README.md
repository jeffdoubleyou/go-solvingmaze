# go-solvingmaze
SolvingMaze API Library for Go

[![GoDoc](https://godoc.org/github.com/jeffdoubleyou/go-solvingmaze?status.svg)](https://godoc.org/github.com/jeffdoubleyou/go-solvingmaze)

## Import
`import github.com/jeffdoubleou/go-solvingmaze`

## Usage
`package main

import (
	"fmt"
	"go-solvingmaze"
)

func main() {
	sm := solvingmaze.New("ABCDEFGHIJKLMNOP", "Los Angeles", 5, true)
	// Create a new quote
	quote := sm.NewPackAndQuote()

	// Define an item to add to the quote
	itemProps := make(map[string]interface{})
	itemProps["Name"] = "Widget"
	itemProps["Sku"] = "W11122"

	// Add the item to the quote
	item, err := quote.AddItem(itemProps)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// Add dimensions for the item
	dimension := item.AddDimensions(map[string]interface{}{"Length": 12, "Width": 8, "Height": 5})

	// Change the dimension height
	dimension.Height(10)

	// Add a second, empty dimension for the item
	dimension2 := item.AddDimensions()

	// Set dimensions individually
	dimension2.Length(24)
	dimension2.Width(12)
	dimension2.Height(8)

	// Define the destination address
	address := make(map[string]interface{})
	address["post"] = "48072"
	address["region"] = "MI"
	address["city"] = "Berkley"
	address["country"] = "US"
	address["residential"] = true

	// Add the destination to the quote
	_, err = quote.Destination(address)
	if err != nil {
		fmt.Printf("Error adding destination: %s", err)
	}

	// Add a void dimension
	void := item.AddVoidDimensions(map[string]interface{}{"Length": 12, "Width": 8, "Height": 2})

	// Change void height
	void.Height(4)

	// Send calculation request
	calc, err := quote.Calculate()

	if err != nil {
		fmt.Printf("Failed to calculate shipping: %s", err)
	} else {
		fmt.Printf("\nRate: $%v\n", calc.Services[0].Rate)
	}
}`




