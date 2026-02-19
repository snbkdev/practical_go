// Точные типы
package main

import "fmt"

// не рабочий формат
// type Numeric interface {
// 	int8 | int16 | int32 | int64 | float32 | float64
// }

// рабочий
type Numeric interface {
	~int8 | int16 | int32 | int64 | float32 | float64
}


type Smallint int8

func doubler[T Numeric](value T) T {
	return value * 2
}

func main() {
	// some code ...

	var five Smallint = 5
	fmt.Println(doubler(five))

//	var four Smallint = 4
//	fmt.Println(doubler(four)) // result ->  Smallint does not satisfy Numeric (possibly missing ~ for int8 in Numeric)
}