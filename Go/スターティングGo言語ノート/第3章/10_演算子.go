// Goは「静的型付け言語」
package main

import "fmt"

func main() {
	// 左シフト：　あふれた桁は切り捨て
	n := 3 << 4
	fmt.Println(n)
	
	// 右シフト：　下にあふれた桁は切り捨て
}