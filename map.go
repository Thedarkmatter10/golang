package main

import (
	"fmt"
)

func main() {

	// 1. What is a map? it is data str. which is used to for mapping
	// Map of Employee Salary
	// one key maps to one value

	sales := map[string]int{
		"Jan": 10,
		"Feb": 20,
		"Mar": 500,
		"Apr": 6000,
		"May": 1,
	}

	// 2. Data types of keys in a map

	// Can be: bool, int, float, struct, array, string, pointer, interfaces, channels etc
	// Cannot be: Slices, Maps, Functions **IMPORTANT**

	// m := map[[]int]string{}
	m := map[[3]int]string{}
	fmt.Printf("Value = %v, Type = %T\n", m, m)

	// 3. Make function
	// when we dont know what values will be in the map at the time of creating it
	map1 := make(map[string]int)
	fmt.Printf("Value = %v, Type = %T\n", map1, map1)

	// 4. Accesing the values form a map
	fmt.Printf("Value = %v, Type = %T\n", sales["Jan"], sales["Jan"])

	// 5. Updating values
	sales["Frankie"] = 40000
	fmt.Printf("Value = %v, Type = %T\n", sales["Jan"], sales["Jan"])
	fmt.Printf("Value = %v, Type = %T\n", sales, sales)

	// 6. deleting keys
	delete(sales, "Brook")
	fmt.Printf("Value = %v, Type = %T\n", sales, sales)

	// 7. Accessing keys not present in the map
	fmt.Printf("Value = %v, Type = %T\n", sales["Jan"], sales["Jan"])
	fmt.Printf("Value = %v, Type = %T\n", sales["g"], sales["g"])

	// 8. Optional Comma Ok syntax
	salesValue, ok := sales["Jan"]
	fmt.Printf("%v %v\n", salesValue, ok)
	salesValue, ok = sales["Jan"]
	fmt.Printf("%v %v\n", salesValue, ok)

	// 9. Check only for presence
	_, ok = sales["Jan"]
	fmt.Printf("%v %v\n", salesValue, ok)

	// 10. length
	fmt.Printf("Length = %v\n", len(sales))

	// 11. Maps are passed by reference
	salesRef := sales
	fmt.Printf("Value = %v, Type = %T\n", salesRef, salesRef)
	delete(salesRef, "Jan")
	fmt.Printf("Value = %v, Type = %T\n", salesRef, salesRef)
	fmt.Printf("Value = %v, Type = %T\n", sales, sales)

}
