学習しているUdemyのコース
https://www.udemy.com/course/iac-with-terraform/

参考にしたzenn記事
https://zenn.dev/yutamiona/articles/b457ee3c9934d4

===

Terraform とは

- Infrastructure as Code でインフラを管理するツール
- コード化された構築設定に従って、自動で任意のクラウド上にデプロイしてくれるツール

Terraform を使った開発の流れ

  1. アカウント作成
  2. terraform でインフラをコード化、環境構築(CLIで実行)
  3. 足りない部分を手動で補完

tfstateファイル
    リソースを展開するときに、そのリソースの状態を記録するファイル（ログファイルのようなもの）

よく使われるコマンド
  terraform plan
  terraform apply
  terraform destroy

HCL2 とは
  terraformのソースコードのこと。独自の言語
  HashiCorp Language 2の略

  Jsonのような形式で記述する
    * コメントは # で始まる
    * コロンではなくイコールで値を指定する
    * ヒアドキュメントが使える
    ブロック構文が使える

  学習するブロックタイプ
    * provider
    * resource
    * data
    * locals
    * variable
    * output
    * terraform

データ型
    * string
    * number
    * bool

    * object
    * tuple ... 複数の値をまとめて扱うことができる型
                RDBでいう1レコードのようなもの
    * list
    * map
    * set

変数の上書き方法

terraform と provider はバージョンを合わせる必要があり、またバージョンを固定する必要がある
