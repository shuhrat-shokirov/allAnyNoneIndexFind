package main

import "fmt"

var data = []account{
	{id: 1, name: "Misha", balance: 10_000},
	{id: 2, name: "Masha", balance: 15_000},
}

func Example_FunctionAllForTrue() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance >= 10_000
	}))
	// Output: true
}
func Example_FunctionAllForFalse() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance > 10_000
	}))
	// Output: false
}

func Example_FunctionAnyForTrue() {
	fmt.Println(any(data, func(item account) bool {
		return item.balance > 10_000
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
		return item.balance < 10_000
	}))
	// Output: true
}
func Example_FunctionNoneFalse() {
	fmt.Println(none(data, func(item account) bool {
		return item.balance <= 10_000
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
		return item.balance > 15_000
	}))
	// Output: -1
}

func Example_FindHasResult() {
	result, ok := find(data, func(item account) bool {
		return item.name == "Masha"
	})
	if ok != nil {
		fmt.Errorf(ok.Error())
	} else {
		fmt.Println(result)
	}
	// Output: {2 Masha 15000}
}

func Example_FindWithError() {
	result, ok := find(data, func(item account) bool {
		return item.name == "Alex"
	})
	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Println(result)
	}
	// Output: Not found
}
