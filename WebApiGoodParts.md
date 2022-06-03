# Web API the Good Parts

<https://www.oreilly.co.jp/books/9784873116860/>

## 1章 Web APIとは何か

- Web APIがいろんなシーンで利用されてきた
- 公開APIはセキュリティ強化が重要

### 1.3 何をAPIで公開すべきか

- システムのコア機能

### 1.3.1 API公開のリスク

### 1.3.2 API公開で得られること

### 1.4 Web APIを美しく設計する重要性

### 1.5 Web APIを美しくするには

- 仕様に従う
- 仕様がないものはデファクトスタンダードに従う

### 1.6 RESTとWebAPI

- これが出たのが2014年でRESTという言葉の意味を注意深く扱う必要があったのだろう

### 1.7 対象となる開発者の数とAPI設計思想

- LSUDs (large set of unknown developers) / SSKDs (small set of known developrs)

### 1.8 まとめ

- Web APIを今すぐ公開仕様
- 美しく設計しよう
- RESTという言葉にこだわりすぎない

## 2章 エンドポイント設計とリクエスト形式(P.19 - P.64)

### 2.1.1 モバイルアプリケーション向けに必要なAPIの機能

1. 画面遷移図を作る
2. 機能一覧を作る
3. 画面遷移図と機能一覧を比較して過不足がないかチェックする

### 2.2 APIエンドポイントの考え方

#### 2.2.1 エンドポイント設計の基本

- 覚えやすく、どんな機能を持つURIなのかがひと目でわかる
  - 短く入力しやすいURI
  - 人間が読んで理解できる
    - 無理に省略しない・よく使われる英単語を選ぶ・複数系・過去形 重要
  - 大文字・小文字が混在してない
  - 改造しやすい
  - サーバ側のアーキテクチャが反映されていない
  - ルールが統一されている

### 2.3 HTTPメソッドとエンドポイント P.29

`GET` `POST` `PUT` `DELETE` `PATCH` `HEAD`

#### 2.3.1 GET

- GETメソッドでも、未読/既読や最終アクセス日付など、GETによりリソースが参照されたことを記録する場合もあるらしい

#### 2.3.2 POST

#### 2.3.3 PUT

- 既存リソースの上書き変更
- 既存リソースの一部変更はPATCHを使う

##### 2.3.5.1 X-HTTP-Method-Overrideヘッダ

- HTMLのFormの仕様でGETとPOSTのみサポートされているが、他のメソッドを利用したい場合などにこのヘッダを使う
  - 利用例）POSTでデータを送信しておいて、X-HTTP-Method-Oberrideヘッダに、DELETEを指定する
  - `_method` を利用する場合もある

### 2.4 APIエンドポイントの設計（P.32）

![スクリーンショット 2022-05-05 10.13.18.png](:/39a3a62ccfa24801bd8e15fa93b26068)

- 基本の考え方：`あるデータの集合` と `個々のデータ` をエンドポイントとして表現し、それに対してHTTPメソッドで操作していく

個々のデータの例
![スクリーンショット 2022-05-05 10.17.15.png](:/1899937252a24ad0a85786ac46b35239)

一覧データの例
![スクリーンショット 2022-05-05 10.17.42.png](:/2a47fe85481a41c8aa46c14af548d9f8)

- 一覧データは、`list`　という言葉をつける場合とそうでない場合に分けれる

友達（ソーシャルグラフ）関連API
![スクリーンショット 2022-05-05 10.19.55.png](:/6ea73e3a6a4d4045b7b070cc8a819c94)

#### 2.4.1 リソースいアクセスするためのエンドポイント設計の注意点

- 複数形の名詞を使う(リソースは「集合」を意味するので複数形が望ましい)
- 利用する単語に注意する
- スペースやエンコードを必要とする文字を使わない
- 単語をつなげる必要がある場合はハイフンを利用する

#### 2.4.2 利用する単語に気をつける

- 辞書を見よう
- 困ったら参考になるAPIを見よう
  - [ProgrammableWeb](https://www.programmableweb.com/)

### 2.5 検索とクエリパラメータ P.42

- 2.5.1 取得数と取得位置はlimit/offsetの組み合わせがおすすめ
  - アイテム数を扱うので利用者の自由度が高い
  - page/per_pageはページ単位で扱うので自由度が下がる
- 2.5.2 相対位置を取得する方法の問題
  - offsetやpageは先頭から探索するのでパフォーマンスが下がる
  - 更新頻度が高いデータは不整合が生じやすい
- 2.5.3 絶対位置でデータを取ろう
  - `このIDより前のもの` や`この時刻より古いもの`のような指定をする
  - Twitterの`max_id`がこれ
- 2.5.4 絞り込みのパラメータは２パターン
  - クエリパラメータに絞り込みむ要素名をつける
    - 複数要素を指定したい場合に利用されがち
  - 指定要素が1つの場合は、`q` を使いがち
- 2.5.4.1 URIにsearchという単語を入れるのはあり？ P.47
  - searchは動詞なので設計の原則からは外れる
  - わかりやすさを優先するなら、検索するためのエンドポイントでsearchはあり（リスト取得との差別化を明示する）
- 2.5.5 クエリパラメータとパスの使い分け
  - 一意のリソースを表すのに必要かどうか
  - 省略可能かどうか

### 2.6 ログインとOAuth2.0 P.49

### 2.7 ホスト名とエンドポイントの共通部分 P.57

- `api.example.com/v1` の形がおすすめ

### 2.9 HATEOAS と　REST LEVEL3 API P.60

- 実際にlevel3はあまり使われてなさそうなので読み飛ばした

## 3. レスポンスデータの設計 P.65 - P.99

- gender: 社会的・文化的性別, sex: 生物学的性別
- 日付フォーマットは、RFC3339を使うのが無難( 2022-10-12T11:30:22+09:00)
  - タイムゾーンは日本だったら `+09:00`　でも良いが、インターネットは世界中と繋がっているので、UTC(世界標準時間) `+00:00` or `Z` を使うのも良い
- 32bitを超えるような巨大な数値は文字列化する 

### 3.7 まとめ

- 今はJSON形式でOK
- データを不要なエンベロープで包まない
- レスポンスはできるだけフラットな構造にする
- 各データの名前が簡潔で理解しやすく、適切な単数・複数形が使われている
  - 多くのAPIで利用されている一般的な英単語を用いる
  - なるべく少ない単語数
  - 単語を連結する場合は連結ルールを統一する (jsonはキャメルケースにしよう)
  - なるべく変な省略形は使わない
  - 単数・複数に気をつける
- エラー形式を統一し、クライアント側でエラー詳細を機械的に理解できる

## 4. HTTPの仕様を最大限利用する

 ステータスコードを正しく使おうという話

### 200系

- 200番台は200以外の時に後続処理が必要な可能性もあるので注意が必要
  - sendgridで202 Acceptedでメールが送信されないケースがあった

### 300系

- 301,302,303,307がリダイレクトで、レスポンスヘッダにLocationが含まれており、そこにリダイレクト先のURIの情報がある
- 300 Multiple Choice: ファイルストレージサービスで指定したキーに対して複数データが存在する場合などに返ってくる
- 304 NotModified 前回のデータ取得から更新されてないことを示す。レスポンスボディは空になる

### 400系

- 400は400系の他に適切なコードがない場合に使われる（bad request）
- 401 unauthorizedは認証エラー、403 Authorizationは認可エラー
- 405 Method not allowed エンドポイントはあるがメソッドが許可されてない
- 406 Not Acceptable クライアントが指定したデータ形式にAPIが対応してない
- 408 Request Timeout クライアントがサーバに送信するのに時間がかかりすぎ
- 409 Conflict リソース競合（例：IDが競合）
- 410 Gone かつては存在したけど今はもうない
- 413 Request Entity Too Large  リクエストボディが大きすぎ
- 414 Request URI Too Long リクエストヘッダが大きすぎ
- 415  Unsupported Method Type Content-Typeの指定データがサーバ側で対応してない
- 429 Too Many Request アクセス許容範囲の閾値を超えた

### 500系

- サーバ側のバグや問題なので、エラーログをきちんと監視して管理者に通知がいくようにしよう

### 4.3 キャッシュとHTTP仕様 P.110

#### キャッシュのメリット

- サーバ通信を減らせるのでユーザーの体感速度向上
- ネットワークが切れたユーザーもサービス継続可能
- ユーザーの通信コスト削減
- サーバの維持費削減

過去の天気予報データなどは更新がまれなので、キャッシュに向いてる。
中継するプロキシサーバについても考える。

P.112から