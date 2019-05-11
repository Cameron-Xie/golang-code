package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	a, b, c, d := "a", "b", "c", "d"

	// string pointer slice
	strPtrSlice := StrPtrSlice{&d, &b, &c, &a}

	sortSortable(strPtrSlice)
}

func sortSortable(s sort.Interface) {
	sort.Sort(s)

	printSlice(s)
}

type StrPtrSlice []*string

func (s StrPtrSlice) Len() int {
	return len(s)
}

func (s StrPtrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StrPtrSlice) Less(i, j int) bool {
	return *s[i] < *s[j]
}

func slice(s interface{}) (iArr []interface{}) {
	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		panic("not a slice or array")
	}

	for i := 0; i < val.Len(); i++ {
		iArr = append(iArr, val.Index(i).Interface())
	}

	return
}

func getValue(input interface{}) reflect.Value {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func printSlice(s interface{}) {
	arr := slice(s)

	if len(arr) < 1 {
		return
	}

	fmt.Printf("\n|%8v|%8v|\n", "TYPE", "VALUE")

	for _, elem := range arr {
		fmt.Printf("|%8T|%8v|\n", elem, getValue(elem))
	}
}
