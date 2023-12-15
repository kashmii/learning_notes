** Siglens **

star: 55 (2023/11/22)
      68 (2023/11/30)

SigLens はオープンソースの Observability ソリューションで、Splunk や Elastic より 100 倍効率的です。

10,000人以上のエンジニアにObservabilityツールを提供してきた経験から、私たちはいくつかのことを学びました：

    開発者は、ログ、メトリクス、トレースのために様々なツールを飛び越えなければならない。
    Splunk、DataDog、NewRelic は非常に高価である。
    ElasticSearchはマシンの数が多すぎて、クラスタのメンテナンスが大変。
    Grafana Loki はクエリのパフォーマンスが遅い。

監視領域における数十年の経験を武器に、私たちは、外部依存ゼロでログ、メトリクス、トレースに特化した観測可能なDBを一から構築することに着手しました。ラップトップで実行できる単一のバイナリで、1日8TBを処理します。

特徴
1. 複数の取り込みフォーマット：Open Telemetry、Elastic、Splunk HEC、Loki
2. 複数のクエリ言語：Splunk SPL、SQL、Loki LogQL
3. シンプルなアーキテクチャ

* SigLensとElasticsearchの比較
SigLens が Elasticsearch よりも 8 倍速いというブログをご覧ください。

* SigLensとClickHouseの比較
SigLensがClickHouseより4倍～37倍高速であるこのブログをご覧ください。

* SigLensとSplunk、Elastic、Lokiの比較
Splunk、Elastic、Grafana Lokiに必要な3000のEC2インスタンスに対し、SigLensはわずか32のEC2インスタンスで24時間、1PB/日の速度でデータを取り込んだというブログをご覧ください。

(README 以上)

===
### Siglens のライバル

Splunk ログ収集ソフト
Elastic
DataDog クラウド環境やオンプレミス環境での監視、トレーシング、ログ管理などを行うためのクラウドベースの監視プラットフォーム
NewRelic パフォーマンスデータをリアルタイムで収集し、可視化・分析・アラートの機能を提供するサービス
Grafana Loki
ClickHouse






===

Elasticsearchとは

Elasticsearchは、検索エンジン、構造化ログ検索、全文検索、分析エンジンとして使用されている。オープンソースの分散型RESTful検索・アナリティクスエンジンで、Javaで実装されており、オープンソースライセンスのApache License 2.0の下で公開されている。Apache Luceneの分散型検索エンジンをベースにしており、構造化されたドキュメントをJSON形式で格納することができる。Elasticsearchは、ドキュメントのインデックス作成、検索、分析を行う。

rdbとの呼び方の違い

RDB	Elasticsearch
データベース	インデックス
テーブル	マッピングタイプ
カラム（列）	フィールド
レコード(行)	ドキュメント

===

ClickHouseとは

ClickHouseは、Yandexが開発したオープンソースの分散型カラム指向データベース管理システムである。カラム指向データベース管理システムは、データをカラム単位で格納するデータベース管理システムである。ClickHouseは、Yandexの内部で使用されている。ClickHouseは、データウェアハウス、オンラインアナリティクス処理（OLAP）のために設計されている。データの読み取りと書き込みの両方を高速に処理することができる。スケーラブルなデータウェアハウスとして設計されている。

集計など、大量のデータをreadする処理に最適化されたDBです。
https://qiita.com/tomo0/items/b730f479ba1b2c534f1a
