package main

import "fmt"

func main() {
	a, b := 10, 5.5

	fmt.Println(a + 5)   // => 15
	fmt.Println(3.1 - b) // => -2.4
	fmt.Println(a * a)   // => 100
	fmt.Println(a / a)   // => 1
	fmt.Println(11 / 5)  // => 2

	fmt.Println(a * int(b))     // => 50
	fmt.Println(float64(a) * b) // => 55

	x := 10
	x++ // x is 11. Same as: x += 1
	x-- // x is 10. Same as: x -= 1

	a = 10
	a += 2 // => a is 12
	a -= 3 // => a is 9
	a *= 2 // => a is 18
	a /= 3 // => a is 6
	a %= 5 // => a is 1

	fmt.Println(5 == 6)   // => false
	fmt.Println(5 != 6)   // => true
	fmt.Println(10 > 10)  // => false
	fmt.Println(10 >= 10) // => true
	fmt.Println(5 < 5)    // => false
	fmt.Println(5 <= 5)   // => true

	fmt.Println(0 < 2 && 4 > 1) // => true
	fmt.Println(1 > 5 || 4 > 5) // => false
	fmt.Println(!(1 > 2))       // => true

}
