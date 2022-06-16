# ソフトウェアアーキテクチャの基礎

<https://www.oreilly.co.jp/books/9784873119823/>

# 1章 イントロダクション

- ソフトウェアアーキテクトが人気の仕事なのにキャリアパスが不明確
  - 1. 定義が明確ではない
  - 2. 役割が広範囲
  - 3. エコシステムの急激な進化
  - 4. 資料が歴史書化
- この本ではソフトウェアアーキテクトの仕事を現代の状況に合わせて再構築したい

## 1.1 ソフトウェアアーキテクチャの定義

- システム構造
- アーキテクチャ特性（ility)
- アーキテクチャ決定
- 設計指針

### 議事

- ソフトウェアアーキテクトは人気だけどキャリアパスがよくわからん
- Wakuconeではソフトウェアアーキテクトっていました？（小林
  - いなかったと思う（みんな好き勝手作ってた）（茶
- マインドマップよき
- ソフトウェアアーキテクトは今までやってなかったのでよくわからん
  - 今までは無意識になんとなくやってた
- **期待値が高まった！**

## P.8 - P.11

## 1.2 アーキテクトへの期待

- アーキテクチャ決定を下す
- アーキテクチャを継続的に分析する
- 最新のトレンドを把握し続ける
- 決定の順守を徹底する
- 多様なものに触れ、経験している
- 事業ドメインの知識を持っている
- 対人スキルを持っている
- 政治を理解し、かじ取りする

### 1.2.1 アーキテクチャ決定を下す

- チームや部門、あるいは企業全体の技術的な決定を導くために使用される、 アーキテクチャ決定や設計指針を定義することが期待されている。
  - 指定ではなく、「ガイドする」

### 1.2.2 アーキテクチャを継続的に分析する

- アーキテクチャや現在の技術環境を継続的に分析し、その上で改善策を提案することが期待されている
  - アーキテクチャの存続力に目を向ける

### 1.2.3 最新のトレンドを把握し続ける

### 議事

- アーキテクチャと継続的に分析する
  - クリーンアーキテクチャでよく言われる。大事
  - ほとんどのプロジェクトで崩壊する。経験上マジで崩壊する
  - 崩壊してから設計を見直す
    - 継続的に設計し続ける。終わりはない
- 茶谷さんはソフトウェアアーキテクトになりたかった
  - ちょっとあきらめかけてきた
  - スクラムマスターとは別物かな？
  - 政治か。。。大変そうな仕事だ
  - 技術系のトップ、ここまで行ければゴール。
    - あきらめずにがんばろう

-

## 1.3 アーキテクチャと交わる

### 1.3.1 エンジニアリングプラクティス

### 議事

- 政治を理解し、舵取りをするのインパクト！
  - 意思決定を含むからそうだよな。。
- コミュニケーションについてうまく立ち回れよと書かれているが
  - 立ち回り方も書いて欲しい
  - このあたりのストレス耐性も持ち合わせていないといけないのか
- 人と会話するとき
  - バックグランドとかリモートになって読み取りづらくなった
  - リモートの対人スキルも必要そう
  - リモートだと会話ではなく、チャットベースになるのでテキストベースになるので、誤解を与えないように考えるコストがかかる
    - ただ、ここは本質ではない
- 見積もり
  - 難しい！
  - 「未知の未知」は突然現れる！
    - マネージャーは「未知の未知」があることを言いたがらない？
    - 途中で要件が変わって、最初作ったレイヤーに収まらなくて、commonが肥大化しがち
    - よくあるケースだけど、これがどんなに開発を難しくするかを分かってもらうには説得が難しい？
    - 未知の未知を発生させやすいのはドメインの部分で、そこを閉じて〜〜だけど、結局銀の弾丸ではないので世にドメインの本が溢れている
      - 被害を最小限に抑えるしかない
  - ウォーターフォールではバッファをかなり積んでいたけど、それでも残業して燃えていた
    - イテレーティブな方がソフトウェアアーキテクチャの性質にあっている
    - 決めうちはできない！
- 具体的なやり方が疑問
  - <https://bliki-ja.github.io/StranglerApplication/>
  - <https://martinfowler.com/articles/feature-toggles.html>

### 1.3.2 運用とDevOps

- アーキテクチャ関連分野の最も明らかな最近の交わりは、DevOpsの登場
- アーキテクチャと運用を連携させることで
  - アーキテクトは設計をシンプルに
  - 運用は最適な対応を任せられる
  - リソースの割り当てが適切ではないことが偶発的な複雑さを生んでいた

### 1.3.3 プロセス

- ソフトウェアアーキテクチャはソフトウェア開発プロセスとは直接関係がない

### 1.3.4 データ

- コードとデータはどちらか一方がかけても意味をなさない共存関係にある
- アーキテクチャの運用面やアーキテクチャ量子については、DB(外部)の重要な関心ごとを含めている

## 1.4 ソフトウェアアーキテクチャの法則

- ソフトウェアアーキテクチャの範囲は広いが、単一的な要素も存在する
- 「どうやって」よりも「なぜ」の方がずっと重要
  - 数ある選択肢の中で「なぜ」その選択がなされたのか

# 第1部 基礎

- 重要なトレードオフを理解するには以下の基本的な概念と用語を理解しなければいけない
  - コンポーネント / モジュール性 / 結合 / コナーセンス

## 2章 アーキテクチャ思考

- アーキテクトは開発者とは異なる視点で物事を見る
- アーキテクトらしく考えるとは4つの側面がある

### 2.1 アーキテクチャと設計

- アーキテクトと開発者の責務の違い
  - アーキテクト
    - ビジネス要件を分析
    - アーキテクチャ特性(イリティ(-ility))を抽出・定義
    - 問題領域に適したアーキテクチャパターンやスタイルを選択
  - 開発者
    - コンポーネントごとのクラス図
    - ユーザーインターフェイス
    - 開発・テスト
- アーキテクトと開発の密なコラボレーションが不可欠
- アーキテクトと設計に境界はない

# P.17 - P.23

- 昔、lib管にいた。ビルドとデプロイだけするチーム（5名くらいの体制）

### 2.2 技術的な幅

- アーキテクトらしさのためには、これまで培ってきた専門性のいくつかを犠牲にしても技術的な幅を広げるためにコストを払う必要がある
- 開発者からソフトウェアアーキテクトに移行する上での躓きポイント
  - 1. 専門性を維持しようとしてボロボロになる
  - 2. 専門性の陳腐化
    - これに気付けないと氷漬け原始人になる

### 2.3 トレードオフを分析する

- アーキテクチャは全てトレードオフ（正解はない）
- プログラマは利点は分かるがトレードオフはわかっていないが、アーキテクトは両方理解が必要
  - トレードオフスライダーを作って分析する

### 2.4 ビジネスドライバーを理解する

- アーキテクトらしく考える
  - システムの成功に必要な維持ネスドライバーを理解し、アーキテクチャに変換する

### 2.5 アーキテクティングとコーディングのバランスをとる

- ソフトウェアアーキテクトも一定レベルのコードが書けるべき
- ソフトウェアアーキテクトとして、アーキテクティングとコーディングを取ることも求められる
  - クリティカルパス上のコードの所有権を握らない
  - コードを書くことで、プロセス・手順・開発環境に精通し、開発チームの問題をよく理解できる
- 開発チームと一緒にコーディングできない場合のバランスの取り方（現場感を持ちるづけるために）
  - POCを頻繁に行う
  - 技術的負債やアーキテクチャに関わるストーリーを引き受ける
    - 自動化の支援を行う
  - まめにコードレビューをする

## 3章 モジュール性

まずは言葉を揃えるところから始める

### 3.1 定義

- モジュールとは
  - より複雑な構造を構築するために使用できる標準化された商品または独立した単位の各集合
  - 本書においては、クラス・関数をはじめとすつ関連するコードのグループ化を表す一般用語として利用する
- [メタオブジェクトプロトコル](https://en.wikipedia.org/wiki/Metaobject)

-

### 3.2 モジュール性の計測

- 凝集度、結合度、コナーセンス

#### 3.2.1 凝集度

- モジュールの要素がどの程度そのモジュールに収まっているかの度合い
- モジュール内の要素同士の関連度
  - レベル7 機能的凝集：モジュール内容の要素が化不足ない状態
  - レベル6 逐次的凝集：一方の出力をもう一方の入力として2つのモジュールが相互作用する
  - レベル5 通信的凝集：2つのモジュールが通信の連鎖を形成する
  - レベル4 手続き的凝集：2つのモジュールが特定の順序でコードを実行する必要がある
  - レベル3 時間的凝集：モジュール間に時間的な依存関係がある
  - レベル2 論理的凝集：モジュール内のデータが論理的に関連し、機能的には関連しない
  - レベル1 偶発的凝集：モジュール内の要素に関連性がない
- 凝集度の欠如を判断するためのメトリクス
  - [CKメトリクス](https://oreil.ly/-1lMh)
  - LCOM(Lack of Cohesion in Methods）
    - フィールドを介して結合されてないメソッド群の総称
![](https://i.imgur.com/icXPyWT.png)

#### 3.2.2 結合度

- 求心性結合（Afferent coupling）
  - コードアーティファクト（コーンポーネント、クラス、関数etc）に外部から入力される接続数を計測
- 遠心性結合（Efferent coupling）
  - 他のコードアーティファクトに出力する接続数を計測

#### 3.2.3 抽象度、不安定度、主系列からの距離

- 抽象度:抽象と実装の比率
- 不安定度：遠心性結合と求心性結合の合計に対する遠心性結合の比率
  - コードベースの不安定さを判断

#### 3.2.4 主系列からの距離

- D = | A + I - 1|
- A：抽象度、I：不安定度
  - 無駄ゾーン: 抽象的すぎて使いづらい
  - 苦痛ゾーン： 具体的すぎてメンテしづらい

#### 3.2.5 コナーセンス

- システムの全体的な正しさを維持するために、あるコンポーネントの変更が別のコンポーネントの変更を必要とする場合、2つのコンポーネントはコナーセント（接続）されている
  - 静的：コードレベルの結合度合いを分析
  - 動的：実行時の呼び出しを分析

![](https://i.imgur.com/T5fhlRr.png)

#### 3.2.6 結合度とコナーセンスのメトリクスを結合する

- コナーセンスは物事がどのように結合されているかにも関心がある

##### 3.2.6.1 1990年代のコナーセンスが持っていた問題

1. コードの低レベルの詳細を見ている
2. 分散アーキテクチャで動機通信をするか非同期通信をするかに対処していない

### 3.3 モジュールからコンポーネントへ

- ほとんどのプラットフォームは、コンポーネントも何らかの形でサポートしている

## 4章 アーキテクチャ特性

- アーキテクチャの特性とは、以下の3つの基準を満たすもの

1. ドメインに依らない、設計に関する考慮事項を明らかにするもの
2. 設計の校是うてきな側面に影響を与えるもの
3. アプリケーションの成功に不可欠か重要なもの

### 4.1 アーキテクチャ特性の（部分的な）リスト

#### 4.1.1 アーキテクチャの運用特性

- 可用性
- 継続性
- パフォーマンス
- 回復性
- 信頼性 / 安全性
- 堅牢性
- スケーラビリティ

#### 4.1.2 アーキテクチャの構造特性

- 構成容易性
- 拡張性
- インストール容易性
- 活用性 / 再利用性
- ローカライゼーション
- メンテナンス容易性
- 可搬性
- アップグレード容易性

#### 4.1.3 アーキテクチャの横断的特性

- アクセシビリティ
- 長期保存性
- 認証
- 認可
- 合法性
- プライバシー
- セキュリティ
- サポート容易性
- ユーザビリティ / 達成容易性

### 4.2 トレードオフと少なくとも最悪でないアーキテクチャ

- アーキテクトはアーキテクチャをできるだけイテレーティブに設計するよう尽力すべき

## 5章 アーキテクチャ特性を明らかにする

- 明らかにするには、ドメインの問題を理解するだけでなく、ステークホルダーと協力し、ドミ炎の観点から本当に重要なことを見極める
- 3方向から特性を明らかにできる
    1. ドメインの関心毎
    2. 要件
    3. 暗黙的なドメイン知識

### 5.1 アーキテクチャ特性をドメインの関心ごとから捉える

- 全てのアーキテクチャ特性をサポートする汎用アーキテクチャ設計は、アンチパターン

### 5.2 要件からアーキテクチャ特性を抽出する

- アーキテクトにとってドメイン知識が常に有益である
[アーキテクチャ・カタのブログ](https://nealford.com/katas/)

### 5.3 事例：シリコンサンドイッチ

#### 5.3.1 明示的な特性

- 明示的なアーキテクチャの特性は、必要な設計の一部として要求仕様書に記載されている

#### 5.3.2 暗黙的な特性

- 要件文書で指定されることはない

## 6章 アーキテクチャ特性の計測と統制