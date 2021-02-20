package main

import "fmt"

func main() {
	practiceArray()
	practiceSlice()
	practiceMap()
}

func practiceArray() {
	array := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println("array is: ", array)
	fmt.Println("array two is: ", array[0])

	for i, v := range array {
		fmt.Printf("数组索引：%d, 对应值：%s\n", i, v)
	}
}

func practiceSlice() {
	array := [5]string{"A", "B", "C", "D", "E"}
	// 基于数组生成切片
	slice := array[2:5]
	fmt.Println("slice is：", slice)

	// 切片声明
	slice1 := make([]string, 4)
	slice2 := []string{"a", "b", "c"}
	slice3 := append(slice2, "f")
	fmt.Println("slice1 is：", slice1, slice2, slice3)
}

func practiceMap() {
	person := make(map[string]int)
	// person := map[string]int{}
	person["mt"] = 28
	person["sy"] = 32

	for k, v := range person {
		fmt.Println("key is： ", k, ", value is：", v)
	}

	fmt.Println("person is：", person)
	delete(person, "sy")
	fmt.Println("person is：", person)
}
