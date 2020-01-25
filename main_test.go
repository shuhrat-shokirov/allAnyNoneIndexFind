package main

import "fmt"

var data = []account{
	{id: 1, name: "Misha", balance: 10_000},
	{id: 2, name: "Masha", balance: 14_000},
	{id: 3, name: "Alex", balance: 8_000},
	{id: 4, name: "Erzhan", balance: 15_000},
	{id: 5, name: "Sasha", balance: 12_000},
	{id: 6, name: "Nastya", balance: 11_000},
	{id: 7, name: "Petya", balance: 9_000},
	{id: 8, name: "Natasha", balance: 7_000},
}

func Example_FunctionAllForTrue() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance >= 7_000
	}))
	// Output: true
}
func Example_FunctionAllForFalse() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance >= 7_500
	}))
	// Output: false
}

func Example_FunctionAnyForTrue() {
	fmt.Println(any(data, func(item account) bool {
		return item.balance >= 7_500
	}))
	// Output: true
}
func Example_FunctionAnyForFalse() {
	fmt.Println(any(data, func(item account) bool {
		return item.balance > 15_000
	}))
	// Output: false
}

func Example_FunctionNoneForTrue() {
	fmt.Println(none(data, func(item account) bool {
		return item.balance < 7_000
	}))
	// Output: true
}
func Example_FunctionNoneFalse() {
	fmt.Println(none(data, func(item account) bool {
		return item.balance > 7_000
	}))
	// Output: false
}

func Example_IndexHasResult() {
	fmt.Println(index(data, func(item account) bool {
		return item.balance > 10_000
	}))
	// Output: 1
}
func Example_IndexHasNotResult() {
	fmt.Println(index(data, func(item account) bool {
		return item.balance <= 3_000
	}))
	// Output: -1
}

func Example_FindHasResult() {
	result, ok := find(data, func(item account) bool {
		return item.name == "Alex"
	})
	if ok != nil {
		fmt.Errorf(ok.Error())
	} else {
		fmt.Println(result)
	}
	// Output: {3 Alex 8000}
}

func Example_FindWithError() {
	result, ok := find(data, func(item account) bool {
		return item.balance < 3_000
	})
	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Println(result)
	}
	// Output: Not found
}
