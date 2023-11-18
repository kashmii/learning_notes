import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// 標準入力
// https://zenn.dev/pikarich/articles/8f1c7fe4d04e93

// 数字を取得する

var i int
fmt.Scanf("%v", &i)

var s string
fmt.Scanf("%v", &s)

// 入力の数が決まっておらず、半角スペース区切りの場合
sc := bufio.NewScanner(os.Stdin)
sc.Scan()
inputs := strings.Split(sc.Text(), " ")

var ps []int
for _, input := range inputs {
	p, _ := strconv.Atoi(input)
	ps = append(ps, p)
}


// if文

if k <= i {
	fmt.Println("Thank you")
} else {
	fmt.Println(k - i)
}

// 配列の各要素をforループで取得

arr := [...]int{100, 200, 300}
sum := 0
for _, val := range arr {
	sum += val
}
fmt.Println(sum)  //=> 600

// スライスの作り方
s := []int{10, 20, 30, 40, 50}

// 文字列にそれが含まれていたら文字列から削除する
if strings.Contains(s, val) {
	s = strings.Replace(s, val, "", -1)
}

// 文字列を指定した回数だけ繰り返す
base_i, _ := strconv.Atoi(strings.Repeat("1", len(s)))

// 数値のスライスをそーとする
// int型のスライスをソートするときには、sort.Intsを使う。（昇順）
// sort.Ints関数は、sort.Sort(sort.IntSlice(s))を呼んでいるだけ。
sort.Ints(numbers)
// 逆順
sort.Sort(sort.Reverse(sort.IntSlice(numbers)))



// 文字列 順番を逆にするメソッド
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}