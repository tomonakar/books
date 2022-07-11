# 並行処理

## 並行処理モデル

- ErlangやScalaはアクターモデル
- GoはCSP(Communicating Sequential Processes)モデル

## Concurrency(並行) vs parallelism(並列)

- 並行処理とはある時間範囲において、複数のタスクを扱うこと
- 並列処理とはある時間の点において、複数のタスクを扱うこと

## [Concurrency is not parallelism](https://go.dev/blog/waza-talk)

- 並行処理は、複数の処理を独立に実行できる構成のこと
- 並列処理は、複数の処理を同時に実行すること

つまり、並行処理とは、構造に関するもので、並列処理とは実行に関すること。

参考：[並行処理と並列処理](https://zenn.dev/hsaki/books/golang-concurrency/viewer/term)

### 環境変数

`GOMAXPROCS` に必要なコア数を設定することで、GOのアプリケーションで必要なコア数を指定することができる

## CSP vs アクターベース concurrency

- CSPのプロセスは匿名ですが、アクターにはIDがあります。
- CSPはメッセージパッシングにチャネルを使用しますが、アクターはメールボックスを使用します。
- アクターはメッセージ配信を介してのみ通信する必要があるため、ステートレスになります。
- CSPメッセージは、送信された順序で配信されます。
- アクターモデルは分散プログラム用に設計されているため、複数のマシンにまたがって拡張できます。
- アクターモデルは、CSPよりも分離されています。

参考：[CSP vs Actor model for concurrency](https://dev.to/karanpratapsingh/csp-vs-actor-model-for-concurrency-1cpg)

## Goroutines


