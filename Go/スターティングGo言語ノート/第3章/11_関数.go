// Goは「静的型付け言語」
package main

import "fmt"

func main() {
	func plus(x, y int) int {
		return x + y
	}
	fmt.Println(plus(4,8))
}