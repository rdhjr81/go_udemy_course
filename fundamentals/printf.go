package main

import("fmt")

func main(){
	testval := 42

	fmt.Printf("binary value %b", testval)

	fmt.Printf("\nhex value %x", testval)

	//print as binary & hex
	a,b,c,d,e,f := 0,1,2,3,4,5


	fmt.Printf("\n %v \t %b \t %#X",a,a,a)
	fmt.Printf("\n %v \t %b \t %#X",b,b,b)
	fmt.Printf("\n %v \t %b \t %#X",f,f,f)
	fmt.Printf("\n %v \t %b \t %#X",c,c,c)
	fmt.Printf("\n %v \t %b \t %#X",d,d,d)
	fmt.Printf("\n %v \t %b \t %#X",e,e,e)
	fmt.Printf("\n %v \t %b \t %#X",f,f,f)
}