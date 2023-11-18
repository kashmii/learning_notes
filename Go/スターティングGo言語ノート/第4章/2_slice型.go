// slice型
// 可変長配列のようなもの

// * 配列は固定長 sliceは可変長

// 関数len 要素数を取得するメソッド
// 関数cap 容量を取得するメソッド

// スライスはmakeだけでなく、配列と同じように作ることもできる
s1 := make([]int, 8)
s2 := []int{2, 3, 5, 7, 11, 13}

// append 要素を追加できる

s2 = append(s2, 21)

// 容量の自動拡張
// sliceではappendによって容量が不足した場合、自動的に拡張される仕組みを持っている

// copy 先の要素を先頭から塗り替えていく
copy(コピー先slice, コピー元slice)

// 完全スライス式
a[low : high : max]

max を置くことでsliceの容量をコントロールできる

// 可変長引数をint型のスライスにする
func sum(s ...int) int {} // sum(1, 2, 3)とか sum(4, 6, 8)とするとそれぞれスライスになる

【Goのsliceは、関数の引数に柔軟性をもたらす】

// sliceの参照渡し
スライスの値を変更したい場合、ポインタを使わなくてもスライスを渡すだけでOK

例：
func moveByCommand(c rune, s []int32) {
    switch c {
    case 'N':
        s[1]++
    case 'E':
        s[0]++
    case 'W':
        s[0]--
    case 'S':
        s[1]--
    }
}