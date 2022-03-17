// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	str1 := "AAA"
	str2 := "aaa"
	str3 := "aAa"
	result1 := str1 == str2
	result2 := str2 == str3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2)
}
