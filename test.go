package main

import "fmt"

func main() {

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst

	fmt.Println(defaultBool)
	fmt.Println(customBool)

}
