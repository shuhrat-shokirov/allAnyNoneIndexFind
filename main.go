package main

type account struct {
	id      int64
	name string
	balance int
}

type accountError struct {
	message string
}

func (a accountError) Error() string {
	return a.message
}

func all(items []account, predicate func(item account) bool) bool {
	return index(items, func(item account) bool {
		return !predicate(item)
	}) == -1
}

func any(items []account, predicate func(item account) bool) bool {
	if index(items, predicate) != -1{
		return true
	}
	return false
}

func none(items []account, predicate func(item account) bool) bool {
	if index(items, predicate) == -1 {
		return true
	}
	return false
}

func index(items []account, predicate func(item account) bool) int {
	for index, item := range items {
		if predicate(item) {
			return index
		}
	}
	return -1
}
func find(items []account, predicate func(item account) bool) (account, error) {
	increment := index(items, predicate)
	if increment != -1 {
		return items[increment], nil
	}
	return account{}, accountError{"Not found"}

}
func main() {

}
