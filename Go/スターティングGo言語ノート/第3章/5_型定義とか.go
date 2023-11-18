// Goは「静的型付け言語」
package main

import "fmt"

func main() {
	var n, m int
	n, m = 100, 600

	// 暗黙的な定義：　変数の代入ではない
	// これは関数の外では不可
	i := 35
	s := "eeee"

	// 複数の変数定義をするならこれもあり
	var (
		g  = 1
		w  = "popop"
		bz = true
	)

	fmt.Printf("%d, %b, %x, %#v\n", 67, 16, 13, [...]int{1, 2, 3, 55, 75})
	fmt.Println(n, m, i, s)
	fmt.Println(s, g, w, bz)
}

// 書式指定子