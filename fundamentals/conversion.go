package main

import("fmt")

func main(){
	y := 42
	z := 42.0

	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 42.123

	fmt.Printf("%v of type %T \n", m, m)

	//wont work - cannot use m (variable of type float32) as float64 value in assignment
	// z = m
	// fmt.Printf("%v of type %T \n", z, z)

	//conversion from float64 to float32
	n64 := 42.123
	fmt.Printf("%v of type %T \n", n64, n64)

	var n64to32 = float32(n64)
	fmt.Printf("%v of type %T \n", n64to32, n64to32)

	//conversion from float32 to float64
	var n32 float32 = 42.2442
	fmt.Printf("%v of type %T \n", n32, n32)

	var n32to64  = float64(n32)
	fmt.Printf("%v of type %T \n", n32to64, n32to64)


}