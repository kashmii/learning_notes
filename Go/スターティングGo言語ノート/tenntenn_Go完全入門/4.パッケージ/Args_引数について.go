package main

func main() {
// os.Args...0 番目は実行したファイルのパス

// flag.Parse の内部は os.Args[1:] ファイル名が覗かれる

// =====

// 標準入力と標準出力

// 標準入力: os.Stdin
// 標準出力: os.Stdout
// 標準エラー出力: os.Stderr

// =====

// プロクラムを終了

// os.Exit(code int)
// プログラムの呼び出し元に終了状態を伝えられる
// 0: 成功（デフォルト）

// log.Fatal
// os.Exit(1)で異常終了させる
// 終了コードがコントロールできないためあまり多用しない

// =====

// defer

// 関数終了時に実行される
// 引数の評価はdefer呼び出し時
// スタック形式で実行される（最後に呼び出したものが最初に実行）

// =====

// 入出力関連の便利パッケージ

// encoding
// strings
// bufio
// strconv
// unicode
}