ポインタとは： メモリ上の特定の場所を指し示す変数

ポインタの主な目的： 変数のアドレスを格納すること

ポインタのメリット： 
メモリの効率的な使用とデータの共有です。
ポインタを使用することで、大きなデータ構造をコピーする代わりに、
メモリ上のデータを直接共有できます。これにより、プログラムのパフォーマンスが向上します
またメモリ上のデータを直接操作もできる




// 登場人物

* アスタリスク *int でポインタであることを表す／宣言する

& アンド 変数のアドレスを取得するときに使う ex| &x ...xのアドレス 正式名称はアドレス演算子

ポインタ変数 ptr := &x のptrのこと *ptrとすることで変数xの元の値を取得できる (つまり、変数のアドレス変数)

ポインタ型 ポインタ変数の型 printfを使うと *int と表される

// 注意

ポインタ型の変数がnilのときに *x 等で呼び出す（デリファレンス）とランタイムパニックになる
