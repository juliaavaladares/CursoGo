package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Arrays e Slices")

	var array [5]string
	array[0] = "Posição 1"
	fmt.Println(array)

	array2 := [5]string{"Posição 1", "Posição 1", "Posição 1", "Posição 1", "Posição 1"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{10,10, 11, 12, 13}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))

	fmt.Println(slice)
	slice = append(slice, 19)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("--------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}