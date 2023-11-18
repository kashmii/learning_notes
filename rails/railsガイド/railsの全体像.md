https://atmarkit.itmedia.co.jp/ait/articles/1103/09/news112.html

## Railitie

* Railitieというライブラリが他の数多くの独立的なライブラリを組み合わせてフレームワークを形成します

## Action Dispatch、Action Controller、Action Pack

### Action Pack
* Action DispatchとAction Controllerをまとめたもの
* リクエストとレスポンスの間のやりとりを処理し、ユーザーが操作できるウェブアプリケーションを構築するための必須コンポーネント

###　Action Dispatch
* リクエストの処理やルーティング、ミドルウェアを提供します
* リクエストがRailsアプリケーションに到達すると、Action Dispatchはルーティングを処理し、適切なコントローラーアクションにリクエストをディスパッチする

### Action Controller
* リクエストを受け取り、モデルの操作をトリガーし、ビューの表示に関するロジックを処理する

## Rack
* Webサーバとアプリケーションフレームワークとの間の共通のインターフェイスを提供します。Railsの他にも、Sinatra、Merbなど、さまざまなWebアプリケーションフレームワークがこのRackに準拠して実装されています。

##　ActiveSupport
* Rubyをより便利にするための機能拡張を寄せ集めたライブラリです。よく知られている機能としては、Object、String、Array、Hash、DateといったRubyの基本的なクラスに便利なメソッドを追加しています。

##　Bundler
* Railsアプリケーションで使用するgemライブラリの管理を行うためのライブラリです。アプリケーションに必要なgemライブラリは、すべてBundlerで管理およびインストールします