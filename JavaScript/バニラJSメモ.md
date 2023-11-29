* 'use strict';
  * エラーモードの有効化, セキュリティの向上, 予約語制限の追加, this の挙動の変更 などが行われる

### fetch関数 Promiseについて

* fetch関数はPromiseを返す
  * fetch() とは
      組み込みのJavaScript関数で、HTTPリクエストを行い、サーバーとのデータの送受信を行う
      Promiseを返す非同期関数

  * Promise とは
      PromiseはJavaScriptオブジェクトで、非同期処理の結果やエラーを表現するための仕組みです。具体的には、Promiseのコンストラクターに渡される関数（executor function）が非同期処理を実行し、その結果を表すPromiseオブジェクトを返します。
      Promiseオブジェクトは、非同期処理の成功時にはresolve関数を、失敗時にはreject関数を呼び出します。Promiseオブジェクトは、その結果を表すだけでなく、その結果を利用する非同期処理を表すthenメソッドを持ちます。thenメソッドは、成功時と失敗時のコールバック関数を登録できます。
