package main

func Index(elements [] int, predicate func(int) bool) int {
	for indexSlice, element := range elements {
		if predicate(element) {
			return indexSlice
		}
	}
	return -1
}
func All(elements [] int, predicate func(int) bool) bool {
	return Index(elements, func(i int) bool {
		return !predicate(i)
	}) == -1
}
func Any(elements [] int, predicate func(int) bool) bool {
	return Index(elements, predicate) != -1
}
func None(elements [] int, predicate func(int) bool) bool {
	return Index(elements, predicate) == -1
}
func Find(elements [] int, predicate func(int) bool) int {
	return elements [Index(elements, predicate)]
}
func main() {}
