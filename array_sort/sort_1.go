package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	a, b, c, d := "a", "b", "c", "d"

	// string slice
	strSlice := []string{d, b, c, a}

	printSlice(sortInterface(strSlice))

	// string pointer slice
	strPtrSlice := []*string{&d, &b, &c, &a}

	printSlice(sortInterface(strPtrSlice))

	// string pointer slice
	strArray := [4]string{d, b, c, a}

	printSlice(sortInterface(strArray))

	// int slice
	intSlice := []int{4, 2, 3, 1}

	printSlice(sortInterface(intSlice))
}

func sortInterface(s interface{}) []interface{} {
	arr := slice(s)

	sort.Slice(arr, func(i, j int) bool {
		return compare(arr[i], arr[j]) == -1
	})

	return arr
}

// 1 = Greater
// 0 = Equal
// -1 = Less
func compare(x, y interface{}) int {
	xVal := getValue(x)
	yVal := getValue(y)

	if xVal.Kind() != yVal.Kind() {
		panic("can't compare apple to orange.")
	}

	if xVal.Kind() == reflect.String {
		return stringCompare(xVal.Interface().(string), yVal.Interface().(string))
	}

	if xVal.Kind() == reflect.Int {
		return intCompare(xVal.Interface().(int), yVal.Interface().(int))
	}

	panic("type not supported.")
}

func getValue(input interface{}) reflect.Value {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func stringCompare(x, y string) int {
	if x == y {
		return 0
	}

	if x > y {
		return 1
	}

	return -1
}

func intCompare(x, y int) int {
	if x == y {
		return 0
	}

	if x > y {
		return 1
	}

	return -1
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

func getTypeName(i interface{}) {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fmt.Println(val.Type().String())
}
