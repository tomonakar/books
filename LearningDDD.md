# Learning Domain-Driven Design

<https://www.amazon.co.jp/dp/B09J2CMJZY/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1>

## 序章

## DDDの戦略的側面

- 「何」の質問に答えることを扱う
- 「なぜ」ソフトウェアを構築している理由も扱う

## DDDの戦術的側面

- 「方法」つまり各コンポーネントの実施方法に関するもの

## 1章 ビジネスドメイン分析

## ビジネスドメインとは？

- 会社の主な活動領域

## サブドメインとは

- ビジネスドメインの目標と目的を達成するための企業の複数の小さい事業・細かい活動領域
- 会社の全てのサブドメインは、ビジネスドメイン、つまり顧客に提供するサービスを形成する

### サブドメインの種類

- コア、ジェネリック、サポートの３種類のサブドメインを区別する

#### コアサブドメイン

- 企業が競合他社とは異なる方法で行う活動
  - 新しいサービスの発明、既存プロセスの最適化、コスト削減などが含まれがち
  - 例）Uberはコアドメインを最適化し進化させる方法として、同じ方向に向かうライダーを一致させてコスト削減した
  - 例）Googleのランキングアルゴリズム
- 実装が簡単なコアサブドメインは短期の競争優位性しか提供できない。よって、通常コアサブドメインは複雑になる

#### コアサブドメインとコアドメイン

- コアサブドメインは、エヴァンス本ではコアドメインとも読んでいる
- 本書ではビジネスドメインと混同するリスクを考慮し、コアサブドメインと明確に分けて表現する

### ジェネリックサブドメイン

- 全ての企業が同じように実行しているビジネス活動

### サブドメインのサポート

- ビジネスをサポートする
- 競争上の優位性にはならない
- 具体的にどんなもの？

## サブドメインの比較

- 競争上の優位性はコアサブドメインのみが提供する
- サポートサブドメインは複雑にはならないが、コアサブドメインは複雑になる

![スクリーンショット 2022-04-27 20.42.37.png](:/c716d2fe03904e71a92b7d200f487a9d)

## ボラティリティ

- コアサブドメインは頻繁に変更が入る(競争優位性に直結するため)
- サポートサブドメインは頻繁には変わらない（競争優位性には関わらないため）

## ソリューション戦略

- コアサブドメインは自分達で作る。購入したり他から持ってき

ても、他社がそれを真似できてしまうので優位性にはつながらない
\- 外部委託も賢明ではない

- ササブドメインのサポートは競争優位性が関係ないので外部委託しやすい

サブドメインのまとめ
![スクリーンショット 2022-04-28 13.37.33.png](:/8ab09361f2d445228469d364a26d1efe)

## 設計上の決定

- サブドメインの種類とその違いを知ることで、すでにいくつかの戦略的な設計上の決定を行うことができる

## ドメインエキスパートは誰か？

- 経験上、要件を考え出す人か、ソフトウェアのエンドユーザーがドメインエキスパート

## 結論

- コアサブドメイン : 競合他社とは異なる方法で実行している活動であり、そこから競争優位性を獲得している
- ジェネリックサブドメイン : 全ての企業が同じようにやっていること。既存ソリューションを利用してコスパを高くできる
- サポートサブドメイン：明らかな解決策の問題。会社が社内で対応しなければならない可能性があるが、競争優位性にはつながらないので外注しやすい

## 2章 ドメイン知識の発見

### 知識発見 P.22

- ソフトウェアを効果的にするには、ドメインの専門家の問題への考え方・メンタルモデルを模倣する必要がある
- ソフトウェアの成功は、ドメイン専門家とソフトウェアエンジニアの間で知識を共有することの有効性に掛かっている

### コミュニケーション

- 知識の共有に必要なコミュニケーションが効果的に取られることは滅多にない
- 大抵、ビジネスサイドとエンジニアは直接会話しない
- ドメイン知識は専門家からPOなどを通じてエンジニアに提供される（押し付けられる）
- 従来のソフトウェア開発ライフサイクルでは、ドメイン知識を、コーディングに必要な分析モデルに変換する過程で必要な情報が欠けてしまう
  - ドメイン知識の前提や、本来必要なことがソフトウェアに必要なモデルとして表現されることは稀

### ユビキタス言語とは？

- P.24 常識はそれほど一般的ではない
- P.25 ユビキタス言語はビジネス言語のみで構成されるべき

### ビジネスドメインのモデル化

- P27 モデルとは、特定の側面を意図的に強調し、他の側面を無視するもの。または現象の単純化された表現。特定の使用を念頭に置いた抽象化。 - レベッカ・ウーフス・ブロック
- P.28 効果的なモデルは、対象とする課題の解決に特化している（決して現実世界のコピーではない）
- 本質的にモデルは抽象化。抽象化により不要な情報を省略し、問題を解決するために必要なもののみを残し、複雑さを処理する。抽象化の目的は曖昧ではなく、絶対的に正確な新しいセマンティックレベルを作成すること
- ユビキタス言語は、ビジネス全てを切り取るのではなく、必要なシステムの実装を可能にするのに十分なビジネスの側面が含まれているもの

### 継続的なプロセス

- ユビキタス言語の育成は継続的なプロセス

### ツール

- Wiki
- ガーキン言語
- NDepend...ユビキタス言語の用語の使用法を検証できる静的コード分析ツール

### 課題

- ユビキタス言語の育成は難しいから、専門家に積極的に質問しよう

# 3章 ドメインの複雑さの管理（P.33 - P.47)

## P.41 - P.50

- P.43 セマンティックドメイン：意味の領域とそれについて話される単語
- P.45 モデルは何かを解決するためのもので、完璧さは求めない
  - 冷蔵庫のモデルは段ボールの切り抜きと巻尺で十分

## 結論 P.46

- ドメインエキスパートのメンタルモデルに内在する対立が発生したら、ユビキタス言語を複数の有界コンテキストで分解する
- サブドメインが検出されている間、制限されたコンテキストが設計される。ドメインを制限されたコンテキストに分解することは戦略的な設計上の決定
- 制限されたコンテキストとそのユビキタス言語は、１チームで実装・保守可能
  - 2つのチームが同じ制限されたコンテキストで作業を共有することは不可能
  - 1チームが複数の制限されたコンテキストで作業することは可能
- コンテキスト境界は、システムを物理コンポーネントに分解する

# 4章 境界のあるコンテキストの統合 (P.49 - P.60)

- 境界づけられたコンテキスト間の関係と統合を定義するためのDDDパターンについて
- [このサイトがこの章を読むにあたって参考になった](https://www.ogis-ri.co.jp/otc/hiroba/technical/DDDEssence/chap3.html)

## 協力パターン

### パートナーシップモデル

- 境界コンテキスト間の統合がアドホックな方法で調整される
  - 同期と通信の問題を引き起こす可能性があるので地理的に分散したチームには適してない

### 共有カーネル

> 互いに関連しあう複数のコンテキストを別々のチームで開発しているときに、わざわざ苦労してモデルの変換層を用意しコンテキストの自立性を保つよりも、モデルの一部を直接共有してしまう方が現実的なこともある。共有される核（カーネル）の部分には、ドメインモデル、ソースコード、DBスキーマなどが含まれる。コンテキスト内ほど頻繁でなくてもいいが、カーネルを共有する複数コンテキストをまとめて、定期的に継続的な統合を実施すること。共有カーネルは特別な部分になるため、他チームへの相談なしに勝手に変更を加えてはならない。

- サブドメインの同じモデルまたはその一部が複数の境界で実装されている
- 共有スコープ：境界コンテキスト同士のライフサイクルを結合するので、小さくした方が良さそう。理想論として、`コンテキスト境界を超えて渡されることを目的とした統合コンストラクタとデータ構造のみで構成されているのが良い`とテキストは言っている
- 実装： 共有カーネルへの変更がそれをしようする全ての青玄されたコンテキストに即座に反映されるように実装する
  - 単一リポジトリであれば、複数の制限されたコンテキストによって参照される同じソースファイルである可能性が高い
  - 共有リポジトリを使用できない場合は、共有カーネルを専用PJに切り出して、制限されたコンテキストでリンクライブラリとして参照する
  - CIは必ず入れた方が良い

#### 適用における検討

- 導入するとコンテキスト間に強い依存を生じるので、複製コストが調整コストよりも高い場合に適用すべき
  - 統合コストと複製コストの違いはモデルの揮発性によって異なる
    - 頻繁に変更が生じるモデルほど統合コストは高くなる
    - よって共有カーネルは最も変化が激しいサブドメインに適用されがち
- レガシーシステムを最新化する場合にも共有カーネルパターンは使われる
  - 共有コードベースは、システムを制限されたコンテキストに徐々に分解するための中間ソリューションとして利用する

## Customer/Supplier Development Teams（顧客／供給者の開発チーム）パターン

> コンテキスト境界の分かれた2システム間の関係が、一方が他方を利用する形になっていることがある。その場合、各チームの役割は必然的に異なったものとなるので、両者の運営を円滑にするために、2チーム間に顧客（使う側）と供給者（使われる側）の関係を導入する。供給チームは顧客チームに対して、実際に交渉や見積もりを行なう。供給チーム側のシステム修正を容易にするため、受け入れテストを供給チーム側の継続的な統合に組み入れる。

## Conformist（順応者）パターン

> 使われる側のチームが力を持っている場合など、主に政治的な理由で顧客／供給者の関係がうまく構築できないこともある。そのときはコンテキスト境界を固持せずに、あえて使われる側のドメインモデル（ユビキタス言語）に隷属的に従うという選択肢もある。当然、使う側では理想的なドメインモデルを設計できなくなるが、その代償としてアプリケーション統合が劇的に簡単になる。

## Anticorruption Layer（腐敗防止層）パターン

> 新規に構築するアプリケーションも、たいていはレガシーなど既存の外部システムと連携しなければならない。既存システムが持つドメインモデルは、これから構築しようとする新しいドメインモデルにはそぐわないことが多い。両者のモデルの対応付けが容易でない場合は、新システムと既存システムとの間に隔離層（腐敗防止層）を設けて、そこで両方向に対するモデルの変換を実装して両者のモデルを完全に独立させてしまう。ただし、腐敗防止層の構築には大きなコストがかかるので注意。

## Separate Ways（別々の道）パターン

> もし2つのアプリケーションがほとんど機能を共有しておらず、統合することにメリットがないのであれば、無理して統合せずに互いに無関係な別々のコンテキストとして開発していった方がいい。

## Open Host Service（公開ホストサービス）パターン

> あるアプリケーションが他の多くのアプリケーションから利用されるような場合は、1つ1つについて変換層を用意するような面倒をせず、プロトコルが公開された共有サービスの形にする。サービスの機能を拡張したいときは、プロトコルを拡張していく。

## Published Language（公表された言語）パターン

> コンテキスト間連携の通信媒体として、共通の情報モデルを定めたいことがある。しかし、既存の自家製ドメインモデルでは不適切なことが多い。モデルが未成熟だったり、無駄に複雑だったり、あるいは文書化されていなかったりするためだ。そのようなときは、きちんと文書化されていて広く周知された共有言語を採用するのがよい。

## Context Map（コンテキストマップ）パターン

> コンテキスト境界を設定し、その中で一貫したドメインモデルが維持できたら、今度はもっと大きな全体像を把握するために（非OOなレガシーシステムを含む）コンテキスト間の対応関係を明確にする。各コンテキストには、チームのユビキタス言語となるような名前を付ける。それから、あるコンテキストのモデルが別のコンテキストのどのモデルに対応するか、というモデル間の変換マップを考える。システム間連携においては、アプリケーション毎にドメインモデルが全く異なっているのがほとんどなので、コンテキストマップをしっかり作っておくことが重要になる。

- [context mapper](https://contextmapper.org/) を使って表現できるよ

# 5章 シンプルなビジネスロジックの実装 (P.63 - P.73)

## トランザクションスクリプト

- [マーチンファウラー](https://bliki-ja.github.io/pofeaa/TransactionScript/)

> ビジネスロジックをプロシージャ群によって形成する。各プロシージャはプレゼンテーションからの単一のリクエストを処理する。
> 多くのビジネスアプリケーションは、一連のトランザクションであると考えられる。 トランザクションは情報の集まりを、ある約束事に基づいてまとめられたものとみなし、変更を加えることもある。 クライアントシステムとサーバーシステムとのやりとりには、 一定量のロジックが含まれる。 データベースの情報を表示するだけの簡単なロジックもあれば、 検証や計算のために多くのステップを含むロジックもある。
> トランザクションスクリプトは、これらすべてのロジックを単一のプロシージャとして構成する。 プロシージャは、データベースに直接（またはデータベースラッパー経由）でアクセスする。 サブタスクはサブプロシージャに分解されてしまうが、 各トランザクションはそれぞれ自分のトランザクションスクリプトを持っている。

```java
DB.StartTransaction();

var job =DB.LoadNextJob();
var json = LoadFile(job.Source);
var xml = ConvertJsonToXml(json);
WriteFile(job.Destination, xml.ToString());
DB.MarkJobAsCompleted(job);

DB.Commit();
```

- トランザクションスクリプトはより高度なビジネスロジック実装パターンの基盤。明らかに単純だが間違いをとても起こしやすい。

### 1. トランザクション動作の欠如

包括的なトランザクションなしで複数の更新を発行する間違いがよくある。
以下の例では、Usersテーブルのレコードを更新後、問題が発生すると、次のログレコードの追加が成功する前にシステムは不整合な状態に陥る

```java
public class LogVisit {
    ...
    public void Execute(Guid userId, DataTime visitedOn) {
        _db.Excecute("UPDATE Users SET last_visit=@p1 WHERE user_id=@p2",
                    visitedOn, userId);
        _db.Excecute(@"INSERT INTO VisitsLog(user_id, visit_date) 
                           VALUES(@p1, @p2)", userId, visitedOn);
    }
}
```

解消するには、両方のデータ変更を含む適切なトランザクションを導入することで修正できる

```java
public classs LogVisit {
    ...
    public void Execute(Guid userId, DataTime visitedOn) {
        try {
            _db.StartTransaction();
            
            _db.Excecute("UPDATE Users SET last_visit=@p1 WHERE user_id=@p2",
                                visitedOn, userId);
            _db.Excecute(@"INSERT INTO VisitsLog(user_id, visit_date) 
                           VALUES(@p1, @p2)", userId, visitedOn);
            _db.Commit();
        } catch {
            _db.Rollback();
            throw;
        }
    }
}
```

### 分散トランザクション

分散トランザクションで統合できない複数のストレージメカニズムを使用している場合は、さらに事態が複雑になる. 最近の分散システムでは、データベース内のデータに変更が加えられてから、メッセージをメッセージバスに公開することにより、システムの他のコンポーネントに変更を通知するのが一般的な方法なので、その方法で実装例を示す

```java
public class LogVisit {
    ...
    public void Execute(Guid userId, DataTime visitedOn) {
            _db.Excecute("UPDATE Users SET last_visit=@p1 WHERE user_id=@p2",
                                visitedOn, userId);
            _messagBus.Publish("VISIT_TOPIC",
                               new { UserId = userId, VisitDate = visitedOn});
    }
}
```

8章でCQRSパターンの例を示し、複数ストレージメカニズムにデータを取り込む方法を学ぶ。9章では、別のデータベースに変更をコミットした後にメッセージを確実に公開できるようにする送信トレイパターンを学ぶ。

### 暗黙のトランザクション

以下の例では、現在のカウンターの状態に1を追加しているが、この処理が失敗して再度実行した場合にカウンターは1増加する。

```java
public class LogVisit {
    public void Execute(Guid userId, DataTime visitedOn) {
            _db.Excecute("UPDATE Usewrs SET visits=visits+1 WHERE user_id=@p1",
                         userId)
    }
}

```

コンシューマーにカウンターの値を渡すように依頼することで、解消できる。カウンターの値を指定するには、呼び出し元は最初に現在の値を読み取り、ローカルで値を増やしてから更新された値をパラメーターとして渡す

```java
public class LogVisit {
    ...
    public void Execute(Guid userId, long visits) {
        _db.Execute("UPDATE Users SET visits = @p1 WHERE user_id=@p2",
                    visits, userId)
    }
}
```

楽観的同時実行制御をしようすることでも解消できる。LogVisit操作を呼び出す前に、呼び出し元は、カウンターの現在値を読み取り、それをパラメータとしてLogVisitに渡す。LogVisitは呼び出し元が最初に読み取った値と等しい場合にのみ、カウンターの値を更新する。

```java
public class LogVisit {
    ...
    public void Execute(Guid userId, long expectedVisitis) {
        _db.Execute(@"UPDATE Users SET visits=visits+1
                    WHERE user_id=@p1 and visits = @p2",
                   userId, visits)
    }
}
```

### トランザクションスクリプトを使用するのはどんな時か？

ビジネスロジックが単純な手続型操作に似ている場合、使いやすい。例えば、ETL操作では、データの抽出、変換、宛先にロードという単純な操作のため、利用しやすい。

トランザクションスクリプトパターンは、ビジネスロジックが単純なサポートサブドメインに適合する。また外部システムと統合するためのアダプターとして、または腐敗防止レイヤーの一部として使用することもできる.

複雑なビジネスロジックでは、トランザクション間のビジネスロジックが重複する傾向にあるため、利用すべきではないが、適切に使えばアンチパターンではない。

## アクティブレコード

> データベーステーブルまたはビューの行をラップし、データベースアクセスをカプセル化し、そのデータにドメインロジックを追加するオブジェクト

トランザクションスクリプトと同様にビジネスロジックが単純な場合をサポートするが、データ構造が複雑な場合でもアクティブレコードは扱いやすくなっている。

```java
public class CreateUser {
   ...
   public void Execute(userDetails) {
    try {
     _db.StartTransaction();
     
     var user = new User();
     user.Name = userDetails.Name;
     user.Email = userDetails.Email;
     user.Save();
     _db.Commit();
    } catch {
     _db.Rollback();
     throw;
    }
   }
}
```

このパターンでは、メモリないオブジェクトをデータベースのスキーマにマッピングする複雑さをカプセル化すること。アクティブレコードは永続性を担当することに加えて、ビジネスロジックを含めることができる。

### アクティブレコードが適するのはいつか？

データベースへのアクセスを最適化することがアクティブレコードの本質的な役割なので、中身としてはトランザクションスクリプトと同様である。なので、せいぜいユーザーの入力を検証するCRUD操作などの比較的単純なビジネスロジックのみをサポートするのが良い。

アクティブドメインレコードは、貧血ドメインモデルのアンチパターンとも呼ばれている。貧血ドメインモデルを否定的に捉えるのではなく、単なるツールとして捉えよう。ビジネスロジックが単純な場合には適している。

参考：[ドメイン駆動設計で保守性をあげたリニューアル事例 〜 ショッピングクーポンの設計紹介](https://techblog.yahoo.co.jp/entry/2021011230061115/)

# 6 複雑なビジネスロジックへの取り組み(P.75 - P. 97)

## 実装

- 動作とデータの両方を組み込んだドメインのオブジェクトモデル
- 構成要素は、集約、値オブジェクト、ドメインイベント、ドメインサービス
  - 仕様は含まれない？

## 値オブジェクト

```java
class Color
{
    int _red;
    int _green;
    int _blue;
}
```

プリミティブ型だけでドメインオブジェクトを表現すると、原始的執着のコードスメルと言われたりする

```java
class Person
{
    private int _id;
    private string _firstName;
    private string _lastName;
    private string _landlinePhone;
    private string _mobilePhone;
    private string _email;
    private int _heightMetric;
    private string _countryCode;
    
    public Person(...) {...}
}
    
static void Main(String[] args)
{
    var dave = new Person(
        id: 30217,
        firstName: "Dave",
        lastName: "Ancelovici",
        ...
        contoryCode: "BG");
}
```

値オブジェクトを使って書き直した例

```java
class Person {
    private PersonId _id;
    private Name _name;
    private PhoneNumber _landline;
    private PhoneNumber _ mobile;
    private EmailAddress _email;
    private Height _height;
    private CountoryCode _county;
    
    public Person(...) {...}
}

static void main(String[] args) {
    var dave = new Person(
        id: new personId(40319),
        name: new Name("Dave", "Ancelovic"),
        landline: PhoneNumber.Parse("0237465001"),
        mobile: PhoneNumber.Parse("08029103921"),
        email: Email.Parse("dave@gmail.com"),
        height: Height.FromMetric(180),
        countory: CountryCode.Parse("BG"));
}
```

Height値オブジェクトは意図を明確にし、特定の測定単位から測定を切り離す。例えば、メートルとインペリアル単位の両方を使用して初期化できるため、ある単位から別の単位への変換、文字列表現の生成、単なり単位の値の比較が容易になる

```java
var heigtMetric = Height.Metric(180);
var heigtImperial = Height.Imperial(5,3);

var string1 = heightMetric.ToString();    // 180cm
var string2 = heightImperial.ToString();   // 5feet 3inches
var string3 = heightMetric.ToImperial().ToString(); // 5feet 11 inches

var fistIsHigher = heightMetric > heightImperial;   // true
```

PhoneNumber値オブジェクトは文字列を解析し、それを検証し、電話番号のsまざまな属性を抽出するロジックをカプセル化できる。例えば、それが属する国と電話番号タイプ

```java
var phone = PhoneNumber.Parse("+35987723503");
var country = phone.Country;     // BG
var phoneType =  phone.PhoneType; // MOBILE
var isValid = phoneNumber.IsValid("+9743230120"); // false
```

データを操作して値オブジェクトの新しいインスタンスを生成するすべてのロジックをカプセル化する場合の例

```java
var red = Color.FromRGB(255,0,0);
var green = Color.Green;
var yellow = red.MixWith(green);
var yellowString = yellow.ToString();  // #FFFF00
```

## 実装

- 値オブジェクトはいずれかのフィールドを変更すると異なる値になるため、不変オブジェクトとして実装する
- 値オブジェクトのフィールドの１つを変更すると概念的には別の値、つまり値オブジェクトの別のインスタンスが作成される

```java
public class Color {
    public readonly byte Red;
    public readonly byte Green;
    public readonly byte Blue;
    
    public Color(byte r, byte g, byte b) {
        this.Red = r;
        this.Green = g;
        this.Blue = b;
    }
    
    public Color MixWith(Color other) {
        return new Color(
        r: (byte) Math.Min(this.Red + other.Red, 255),
        g: (byte) Math.Min(this.Green + other.Green, 255),
        b: (byte) Math.Min(this.Blue + other.Blue, 255),
        );
    }
}
```

値オブジェクトの同値性は、idフィールドや参照ではなく、それらの値に基づいているため、同値性チェックをオーバーライドして適切に実装する必要がある

```java
public class Color {
    ...
    public override bool Equals(object obj) {
    var other = obj as Color;
    return other != null &&
            this.Red = other.Red &&
            this.Green == other.Green &&
            this.Blue == other.Blue;
    }
    
    public static bool operator == (Color lhs, Color rhs) {
        if (Object.ReferenceEquals(lhs, null) {
            return Object.ReferenceEquals(rhs, null);
        }
      return lhs.Equals(rhs);
    }
            
    public static bool operator != (Color lhs, Color rhs) {
        return !(lhs == rhs);
    }
    
    public override int GetHashCode() {
        return ToString().GetHashCode();
    }
    ...
}
```

プリミティブなString型を値オブジェクトとして利用することは、値オブジェクトの概念と矛盾するけど、.NETやJavaなどでは、String型は値オブジェクトとして正確に実装される.Immutableだし、豊富な動作カプセル化が既におこなわれているので。

## 値オブジェクトを利用するのはいつ？

- できる限り使う
- ロジックをカプセル化できるし、Immutableなのでスレッドセーフに動作する

## Entities

エンティティはインスタンスごとに区別するために明示的な識別フィールドが必要

```java
class Person {
    public Name Name { get; set; }
    public Person(Name name) {
        this.Name = name;
    }
}
```

上記だと名前被りが起こると人を区別できない

```java
class Person {
    public readonly PersonId id;
    public Name Name { get; set; }
    
    public Person(PersonId id, Name name) {
        this.Id = id;
        this.Name = name;
    }
}
```

これだとIdが識別子になる
PersonIdはビジネスドメインのニーズにあったデータ型を使用できる。
例えば、社会保障番号やGUID, 文字列, 数値など。
IDは基本的に、エンティティのライフサイクルに置いて不変になる.

## アグリゲート（集約）

- 集約はエンティティ
- 集約の目的はデータの一貫性の保護

### 一貫性の強制

- データが破損する可能性のある複数の方法に対するアダプター・境界になる
- 集約のロジックはすべての変更要求時にビジネスルールに矛盾してないかバリデーションを実行する

### 実装

- 集約のビジネスロジックのみがその状態を変更できるようにして整合性を強化する
- 集約の外部にあるすべてのプロセス・オブジェクトは集約の状態の読み取りのみが許可されている
- 状態は、集約の publicなメソッドを実行することによってのみ変更可能
- 集約のpublic インターフェースとして公開されている状態変更メソッドは「何かを行うためのコマンド」として扱われ、`コマンド` と呼ばれる。

コマンドの実装例1

```java
public class Ticket {
    ...
    public void AddMessage(UserId from, string body) {
        var message = new Message(from, body);
        _message.Append(message);
    }
    ...
}
```

上記のコマンドをパラメータオブジェクトとして表すこともできる。
コマンドの実行に必要なすべての入力をカプセル化する

```java
public class Ticket {
    ...
    public void Execute(AddMessage cmd) {
        var message = new Message(cmd.from, cmd.body);
        _messge.Append(messge);
    }
    ...
}

```

コマンドが集約コード中でどのように表現されるかは、好みの問題。
著者の好みは、コマンド構造を定義し、それらを関連するExecuteメソッドに多相的に渡す方法。

アグリゲートのパブリックインターフェースは、入力を検証し、関連するすべてのビジネスルールと不変条件を実行する責任を負います。この厳密な境界は、アグリゲートに関連するすべてのビジネスロジックが、アグリゲート自体という1つの場所に実装されることを保証します。

このため、アグリゲートの操作を編成するアプリケーション層 は、アグリゲートの現在の状態をロードし、必要なアクションを実行し、変更された状態を永続化し、操作の結果を呼び出し元に返すだけでよく、非常にシンプルです。

以下のように、本質的にアプリケーション層の操作はトランザクションスクリプトパターンを実装する。操作をアトミックトランザクションとして調整する必要があるが、集約全体への変更は、成功・失敗のみとなり、部分的な更新をコミットすることはない

```java
public ExecutionResult Escalate(TicketId id, EscalationReason reason) {
    try {
        // 集約の現在の状態をロード
        var ticket = _ticketRepository.Load(id);
        var cmd = new Escalate(reason);
        
        // 必要なアクションの実行
        ticket.Execute(cmd);
        _ticketRepository.Save(ticket);
        
        // 操作の結果を返す
        return ExecutionResult.Success();
        
  // 集約はデータの一貫性を保つことが目的なので、並行性のチェックに注意する
    } catch (ConcurrencyException ex) {
        return ExecutionResult.Error(ex)
    }
}
```

集約は一貫性の保護が目的なので、複数のプロセスが同じ集約を同時に実行指定りる場合、後者のトランザクションが最初のトランザクションによってコミットされた変更を盲目的に上書きしてはならない。

アプリケーション層は、トランザクションスクリプトのコレクションであり、競合する更新によってシステムのデータが破損することを防ぐには、同時実行管理が不可欠となる。

集約の保存に使用されるDBは同時実行管理をサポートする必要がある。
もっとも単純な形式では集約のたびにインクリメントされるバージョンフィールドを保持する

```java
class Ticket {
   TicketId _id;
   int _version;
   ...
}

```

DBコミット時は、このバージョンが一致しているかSQLでチェックすることになる。

```java
UPDATE tickets
SET ticket_status = @new_status,
agg_verson = agg_version + 1 WHERE ticket_id=@id and agg_version=@expected_version;
```

## トランザクション境界

- 集約はトランザクション境界としても機能する
- 集約の状態に対するすべての変更は1つのアトミック操作としてコミットされる
- マルチアグリゲートトランザクションは許容せず、集約の状態変更は、データベーストランザクションごとに1つの集約として個別にのみコミットできる
- これはモデリングの制約で、トランザクション境界の設計を慎重に行う必要が生じる
- 複数のアグリゲートで変更をコミットする必要がある場合は、設計が間違っている

でもさ、同じトランザクションで複数のオブジェクトを変更する必要がある場合はどうするのさ？

## エンティティの階層化

エンティティを集約の一部として作成し、階層化することで解消する
以下の例は複数のエンティティにまたがるビジネスルールを示している

```java
public class Ticket {
    ...
    List<Messag> _messages;
    ...
    public void Execute(EvaluateAutomaticActions cmd) {
        // エージェントがエスカレーションされたチケットを開かなかった場合、
        // 応答時間制限の50%以内に、別の時間に自動的に再割り当てする
        if (this.IsEscalated && this.RemainingTimePercentage < 0.5 &&
           GetUnreadMessagesCount(for: AssignedAgent) > 0) {
            _agent = AssignNewAgent();
        }
    }
    public int GetUnreadMessagesCount(UserId id) {
        return _messages.Where(x => x.To == id && !x.WasRead).Count();
    }
    ...
}
```

### 集約の設計について

- データの一貫性は、集約を設計するための便利な指針になる
- 集約のビジネスロジックが強い一貫性を保つために必要な情報のみを集約の一部にすべき
- 結果生合成のあるすべての情報は集約の協会の外側に存在すべき
- エンティティがアグリゲートに属しているかどうかは、集約が最終的に一貫性のあるデータで動作する場合、無効なシステム状態につながる可能性のあるビジネスロジックを含んでいるかどうかを調べるとわかる

```java
public class Ticket {
    private UserId      _customer;
    private List<ProductId>  _products;
    private UserId     _assigneAgent;
    private List<Message>  _messages;
}
```

上記の例ではIDでエンティティを参照しているが、これにより、集約の外部にあることを明示できる

### 集約ルート

以下の例は集約は特定のメッセージを既読としてマークできるコマンドを公開している。この操作はメッセージエンティティのインスタンスを変更しているが、その集約ルートであるチケットを会してのみアクセスできる。

```java
public void Execute (AcknowledgeMessge cmd) {
    var message = _messages.Where(x =>. x.Id == cmd.id).First();
    message.WasRead = true;
}
```

### ドメインイベント

- 集約ルートのパブリックインターフェースに加えて、外部から集約と通信できる別のメカニズム
- ドメインイベントはビジネスドメインで発生した重要なイベントを説明するメッセージ
- ドメインイベントは既に起こったことを説明しているので、過去形で定式化する
- ドメインイベントは集約のパブリックインターフェースの一部
  - 外部システムもドメインイベントに応答して独自のロジックをサブスクライブして実行できる

ドメインイベントの例）

- チケットのアサイン
- チケットのエスカレーション
- メッセージの受信

イベントを過去形で表現

```java
{
    "ticket-id": "cdafdff-asdfa11-234dfq-asdfadf1",
    "event-id": 134,
    "event-type": "ticket-escalated",
    "escalation-reason": "missed-sla",
    "escalaton-time": 1231231401
}
```

```java
public class Ticket {
    ...
    private List<DomainEvent> _domainEvents;
    ...
    publid Exute(RequestEscalation cmd) {
        if (!this.IsEscalated && this.RemainingTimePercentage <= 0) {
            this.IsEscalated = true;
            var escalatedEvent = new TicketEscalated(_id, cmd.Reason);
            _domainEvents.Append(escalatedEvent);
        }
        ... 
    }
}
```

### ユビキタス言語

集約はユビキタス言語を反映する必要がある。集約の名前、データメンバー、アクション、ドメインイベントに使用される用語はすべて、協会コンテキストのユビキタス言語で定式化されるべき。Eric Evansが言うように、コードは開発者同士やドメインエキスパートと話すときに使うのと同じ言語に基づいていなければならない。これは、複雑なビジネスロジックを実装する際に特に重要なこと。

### ドメインサービス

どの集約や値オブジェクトにも属さない、あるいは複数の集約に関連すると思われるビジネスロジックは、ドメインサービスとして実装する

```java
public class ResponseTimeFrameCalculationService {
    ...
    public ResponseTimeframe CalculateAgentResponseDeadline(UserId agentId,
        Priority priority, bool escalated, DateTime startTime) {
        var policy = departmentRepository.GetDepartmentPolicy(agentId);
        var maxProcTime = poicy.GetMaxResponseTimeFor(priority);
        
        if (escalated) {
            maxProcTime = maxProcTime * policy.EscalationFactor;
        }
        var shifts = _departmentRepository.GetUpcomingShifts(agentId, startTime,
                                                             startTime.Add(policy.MaxAgentResponseTime));
        return CalculateTargetTime(maxProcTime, shifts);
    }
}
```

ドメインサービスを利用すると、複数の集約の作業を調整できる。
ただし、1つのデータベーストランザクションでアグリゲートの1つのインスタンスのみを変更するという制約を常に念頭に置く必要がある。ドメインサービスもこの制約に従う。
ただし、ドメインサービスは複数の集約のデータを読み取る必要があるビジネスロジックの実装に役立つ.

### 複雑さの管理

- システムの複雑さとは、システムの動作を制御・予測することの難しさの指標

# 7章 時間の次元のモデリング(P.99 - P.116)

この章ではイベントソーシングの概念説明と、イベントソーシングをドメインモデルパターンと組み合わせてイベントソースドメインモデルにする方法を学ぶ。

- イベントソーシングパターンは時間の次元をデータモデルに導入する

- 集約の現在の状態を反映するスキーマの代わりに、イベントソーシングベースのシステムでは集約のライフサイクル全ての変更を文書化するイベントを永続化する

  - state sourcing vs event sourcing
  - 要はアプリケーションの状態への全ての変更を一連のイベントとしてキャプチャする
- [参考：martinFowler.com](https://martinfowler.com/eaaDev/EventSourcing.html)

顧客の状態はドメインイベントから予測できる。
先に見たように、これらのドメインイベントから顧客の状態を簡単に投影することができる。あとは、各イベントに簡単な変換ロジックを順次適用していくだけである。

イベントによる単一の状態表現のみを投影するわけではない。
例えば、検索を実装する必要があるとする。
リードの連絡先情報は更新可能で、セールスエージェントが他のエージェントが適用した変更に気づかずに履歴値を含むコンタクト情報を使用してリードを見つけたいと思うかもしれない。その場合でも検索できるように、履歴情報を以下のよう投影できる。

```java
public class LeadSearchModelProjection {
    public long LeadId { get; private set; }
    public HashSet<string> FirstNames { get; private set; }
    public HashSet<string> LastNames { get; private set; }
    public HashSet<PhoneNumber> PhoneNumbers { get; private set; } 
    public int Version { get; private set; }
    
    // LeadInitializedイベントでリードの個人情報を入力する
    public void Apply(LeadInitialized @event) {
        LeadId = @event.LeadId;
        FirstNames = new HashSet<string>();
        LastNames = new HashSet<string>();
        PhoneNumbers = new HashSet<PhoneNumber>();
        
        FirstNames.Add(@event.FirstName);
        LastNames.Add(@event.LastName);
        PhoneNumbers.Add(@event.PhoneNumber);
        
        Version = 0;
    }
    
    // ContactDetailsChangedイベントでもリードの個人情報を変更する
    public void Apply(ContactDetailsChanged @event) {
       FirstNames.Add(@event.FirstName);
        LastNames.Add(@event.LastName);
        PhoneNumbers.Add(@event.PhoneNumber);
        Version += 1;
    }
    
    // 上記以外のイベントは、特定のモデルの状態に影響を与えないので、
    // イベント情報のみ更新する
    public void Apply(Contacted @event) {
        Version += 1;
    }
    
    .... 
}
```

上記のイベントを適用した結果、ドメインオブジェクトは以下の状態になる

```java
    LeadId: 12
    FirstNames: ['Casey']
    LastNames: ['David', 'Davis']
    PhoneNumbers: ['555-2951', '555-8101']
    Version: 6
```

イベントソーシングパターンが機能するには、オブジェクトの状態に対する全ての変更がイベントとして記録され、永続化される必要がある。

システムのイベントを格納するデータベースは、唯一の強く一貫したストレージある必要がある。

## イベントストア

- イベントの永続化で利用するDBはイベントストアと呼ぶ
- イベントストアは追記型ストレージ
- 変更や削除を許可してはならない

```java
interface IEventStore {
    IEnumerable<Event> Fetch(Guid instanceId);
    void Append(Guid instanceId, Event[] newEvents, int expectedVersion);
}
```

上記のAppendメソッドでは、同時実行管理を行うために、expectedVersionが必要となる。イベント追加時にバージョンをチェックして、問題があれば同時実行例外を発生させる。

イベントソーシングのパターンは新しくない。
金融業界では台帳変更を表現するためにイベントを使用し、追記型のログで状態を記録する。現在の状態は大腸の記録を投影することで推測可能となる。

## イベントソーシングのドメインモデル

- オリジナルのドメインモデルは、集約の状態表現を維持し、選択されたドメインイベントを発信する
- イベントソースのドメインモデルは、集約のライフサイクルをモデル化するため、ドメインイベントを排他的に使用する

### イベントソーシングにおける集約の操作

大体以下のスクリプトになる

- 集約のドメインイベントをロード
- 状態表現を再構築し、イベントをビジネス上の意思決定に使用できる状態表現に投影する
- 集約のコマンドを実行し、ビジネスロジックを実行し、その結果、新たなドメインイベントを生成する
- 新たなドメインイベントをイベントストアにコミットする

実装例

```java
public class TicketAPI {
    private ITicketsRepository _ticketsRepository;
    ...
        
        public void RequestEscalation(TicketId id, EscalationReason reson) {
        // 関連するイベントをロード
        var events = _ticketRepository.LoadEvents(id);
        // チケットを投影
        var ticket = new Ticket(events);
        var originalVersion = ticket.Version;
        var cmd = new RequestEscalation(reson);
        ticket.Execute(cmd);
        // 永続化
        _ticketRepository.CommitChanges(ticket, originalVersion);
    }
    ...
}

public class Ticket {
    ...
    private List<DomainEvent> _domaiEvents = new List<DomainEvent>();
    private TicketState _state;
    ...
    public Ticket(IEnumerable<IDomainEvents> events) {
        _state = new TicketState();
        // state projector クラスから、状態を取り出し、順次イベントを適用する
        foreach (var e in events) {
            AppendEvent(e);
        }
    }
    
    // 受信したイベントをTicketStateに私、チケットの現在の状態をメモリに保存する
    private void AppendEvent(IDomainEvent @event) {
        _domainEvents.Append(@event);
        ((dynamic)state).Apply((dynamic)@event);
    }
    
    // IsEscalatedフラグを明示的にTrueに設定しない
    // 代わりに適切なイベントをインスタンス化し、それをAppendEventに渡す
    public void Excecute(RequestEscalation cmd) {
        if (!_state.IsEscalated && _state.RemainingTimePercentatge <= 0) {
            var escalatedEvent = new TicketEscalated(_id, cmd.Reason);
            AppendEvent(escalatedEvent);
        }
    }
    ...
}


// 集約のイベントコレクションに追加された全てのイベントは、TicketStateクラスの状態投影ロジックに渡され、そこでイベントデータに従って関連するフィールドの値が変更される

public class TicketState {
    public TicketId Id { get; private set; }
    public int Versin { get; private set; }
    public bool IsEscalated { get; private set; }
    ...
    public void Apply(TicketInitialized @event) {
        Id = @event.Id;
        Version = 0;
        IsEscalated = false;
        ...
    }
    public void Apply(TicketEscalated @event) {
        IsEscalated = true;
        varsion += 1;
    }
    ...
}

```

hydrate: 水和。静的なものから動的なものにするイメージ。

### なぜ 「イベント・ソーシング・ドメインモデル」なのか？

ドメインモデルの集合体のライフサイクルの変化を表現するためにイベントソーシングを利用している、ということを明示的に示すため

### このパターンのメリット

- タイムトラベル
- システムの状態と動作に対する深い洞察を提供する（コアサブドメインの最適化に繋がる）
- 監査ログ
- 高度な楽観的同時実行管理

### デメリット

- 学習曲線
- モデルの進化 (イベントが不変な前提でシステムが構築されるので、変更が入るときつい)
- アーキテクチャの複雑さ

# 8章アーキテクチャパターン

## CQRS

- ビジネスロジックとインフラへの配慮はポート&アダプター方式を同じ
- データ管理方法が異なる
  - コマンド実行モデル・読み込みモデルの2つに分ける
- precached projection: プリキャッシュされたRDBの射影のこと
  - 必要な列を取り出して組成させる
  - プリキャッシュなので自由に削除・再生成できる
- 同期プロジェクション：キャッチアップ・サブスクリプションモデルでOLTPデータの変更を取り込む
- 非同期プロジェクション：コマンド実行モデルはコミットされたすべての変更をメッセージバスに後悔する
  - システムのプロジェクションエンジンは、発行されたメッセージを購読し、それを使って読み取りモデルを更新する
  - 分散コンピューティングの課題を多く抱えているため、システムが複雑になる
  - 例えば、重複・読み取り順などの課題

## CQRSの誤解

- コマンドはその結果のデータを返して良い
- ただし、返却されるデータは、強い一貫性のあるモデル（コマンド実行モデル）から発信されるべきという制約がある

## CQRSを利用する場合

- 異なる種類のDBにデータが格納され散る可能性がある複数のモデルで、同じデータを扱う必要がる場合に有効
- このパターンはDDDのコアバリュー
- 目の前の課題に対して、最も効果的なモデルで作業し、ビジネスドメインのモデルを継続的に改善することをサポートする
- インフラ観点では、異なるDBの強みを利用しやすい
  - 例えば、コマンド実行モデルの格納にはRDBを、全文検索のためのインデックでは、高速なデータ検索のためのプリレンダリングフラットファイルを、各DB同士で同期しながら利用することができる
- イベントソースモデルと相性が良い
  - イベントソースモデルでは、集約の状態に基づくレコード参照はできないが、CQRSはこれを可能にする

## 結論

- レイヤードアーキテクチャは、技術的な問題tんに基づいてコードベースw分解する。このパターンは、ビジネスロジックとデータアクセスの実装を結合するため、アクティブなレコードベースのシステムに適している
- ポート＆アダプターアーキテクチャは、ビジネスロジックを中心におき、それをすべてのインフラ依存関係から切り離す。ドメインモデルパターンで実装されたビジネスロジックと相性が良い
- CQRSパターンは、同じデータを複数のモデルで表現する。このパターンはイベントソースドメインモデルに基づくシステムには必須だが、複数の永続的なモデルを扱う方法が必要なシステムでも使用できる。

## イベント駆動型アーキテクチャの参考

- [AWS DevAx connect シーズン1](https://aws.amazon.com/jp/devax-connect-on-demand/)

# 9章 コミュニケーションパターン

- コンテキスト境界間で通信を行う際のデザインパターンを整理する
- コンテキスト間のモデルを翻訳する必要があるが、以下の２パターンが主な方法
  - 1.  チーム間で協力的にコミュニケーションをとってプロトコルをアドホックに調整する
  - 2.  共有カーネルを使う

## モデルの翻訳パターン

- ステートレスモデル翻訳：変換する側の境界コンテキストが、プロキシパターンを実装する

### 同期通信の場合

- コンテキストに変換ロジックを埋め込む
  - 場合によっては変換ロジックをAPIゲートウェイパターンなど外部にオフロードすることでコストを節約できる

### 非同期通信の場合

- 通信で使用されるモデルを変換するために、メッセージプロキシを実装する
  - 通信元の境界コンテキストからメッセージをサブスクライブする中間コンポーネントになる
  - プロキシは必要なモデル変換を適用し結果のメッセージを対象のサブスクライバー転送する
  - 加えてこのインターセプトコンポーネントは、無関係なメッセージを削除するフィルタの役割も果たす

![スクリーンショット 2022-05-16 8.54.52.png](:/37e54ed989e64e469f702e9cb6515a71)

#### よくある間違い

- モデルのオブジェクト用に公開された言語を設計し、ドメインイベントをそのまま公開することで、境界コンテキストの実装モデルが公開されてしまうこと。境界コンテキストの実装モデルを公開する必要はない。非同期翻訳はドメインイベントをインターセプトして公開言語に変換するために使われる。
- さらに、メッセージを公開言語に変換することで、バインドされたコンテキストの内部のニーズを満たすプライベートイベントと、他のバインドされたコンテキストとの統合を目的としたパブリックイベントの区別が可能になります。

![スクリーンショット 2022-05-16 9.06.49.png](:/8b15bfe8cc544dc087f457b8590abc7c)

## ステートフルモデルの翻訳

### 受信データの集約

境界コンテキストが着信要求を集約し、パフォーマンス最適化のためにバッチ処理をする場合、同期処理と非同期処理の両方で集約が必要になる場合がある

![スクリーンショット 2022-05-16 9.10.13.png](:/e79814a01f08403ea61685c2e7d9ad64)

複数のメッセージを結合して単一メッセージにすることもできる

![スクリーンショット 2022-05-16 9.11.09.png](:/33fe566fd15d4349881323d74439742d)

受信データを集約するモデルの変換はAPIゲートウェイを使用して実装できない。
そのため、複雑なステートふる処理が必要になる。
それに応じて処理するには変換ロジックに独自の永続ストレージが必要となる。

![スクリーンショット 2022-05-16 9.12.45.png](:/6dd2ad258f644e12ab00fa34dda31e18)

この複雑なステートフルトランスレーションは、既存製品を利用することで実装を回避可能

- ストレームプロセスプラットフォーム
  - kafka, aws kinesis
- バッチ
  - apach NiFi, aws glue, spark

### 複数のソースを一元管理

- 要するに複数のコンテキストの情報を結合する場合は複雑になるから気をつけろという話。フロントエンドは複数のバックエンドからデータを集約して表示する
- 複数の境界コンテキストからデータを集めてビジネスロジックを走らせる場合は、データを統合する処理は腐敗防止層に切り出して、ビジネスロジックと分けることで複雑性をカプセル化できる

## 集約の統合

ドメインイベントはどうやってメッセージバスに発行されて、共有されるのだろう。
解決策見る前に、まずはイベントパブリッシングプロセスでよくある失敗例をみる。

```java
public class Camplain
{
    ...
    List<DomainEvent> _events;
    ImessageBus _messageBus;
    ...
    public void Deactivate(string reason)
    {
        for (l in _locations.Values())
        {
            l.Deactivate();
        }
        IsActive = false;
        
   // 新しいイベントがインスタンス化
        var newEvent = CampaignDeactivated(_id, reason);
    // 集約のドメインイベントの内部リストに追加
        _events.Append(newEevent);
    // メッセージバスにパブリッシュ
        _messageBus.Publis(newEvent);
    }
}
```

集約からドメインイベントを直接パブリッシュする上の実装は以下の理由でだめ

1. 新たな集約のステートがDBに保存される前にディスパッチされる
 - 例えばサブスクライバーがキャンペーンの無効化を受け取ったとしても、その状態変更はまだ永続化されてないので不整合が生じる
2. レース条件・操作の無効化・DBトランザクションがコミットに失敗するなどが発生した場合に、トランザクションがロールバックしても、イベントは既にパブリッシュされているため、撤回できない

以下は新しいドメインイベントの発行をアプリケーションレイヤーで行う例

```java
public class ManagementAPI
{
 ...
 private readonly IMessageBus _mesageBus;
 private readonly ICampainRepository _repository;
 ...
 publid ExecutionResult DeactivateCampaign(CampaignId id, string reason)
 {
  try
  {
   // キャンペーン集約のインスタンスをロード
   var campain = repository.Load(id);
   campaign.Deactivate(reason);
   _repository.CommitChanges(campaign);
   
   // DBに更新された状態がコミットされた後にドメインイベントが
   // メッセージバスにパブリッシュされる
   var events = campaign.GetUnpublishedEvents();
   for (IDomainEvent e in events)
   {
    _messageBus.publish(e);
   }
   campaign.ClearUnpublishedEvents();
  }
  catch(Exception ex)
  {
   ...
  }
 }
}
```

実はこのコードもまだ安全ではない。
例えば、メッセージバスがダウンしている場合など, ロジックを実行ししているプロセスが何らかの理由でドメインイベントの発行に失敗したり、コードを実行しているサーバがコミット直後にダウンしたりするなどで、システムが矛盾した状態で終了するエッジケースが存在する

## 送信トレイ（outbox）パターン

以下のアルゴリズムで、ドメインイベントの確実な公開を保証する

- 更新された集約の状態と新しいドメインイベントの両方は、同じアトミックトランザクションでコミットされる
- メッセージのリレーは、新たにコミットされたドメインイベントをデータベースから取得する
- リレーはドミエンイベントをメッセージバスにパブリッシュする
- 公開に成功すると、メッセージリレーはDB上のイベントを公開済みマークをつけるか、イベントを削除する

![スクリーンショット 2022-05-16 15.24.15.png](:/08496f14e5d546538daa8613113960fc)

RDBを使う場合は、専用テーブルを使用してDBの機能を利用すると実装しやすい
![スクリーンショット 2022-05-16 15.28.34.png](:/61aafd1bea914d0ab0ca9183be96dd64)

マルチドキュメントトランザクションをサポートしないNoSQLの場合は、送信ドメインイベントは集約のレコードに埋め込む必要がある

```java
{
 "campaign-id": "askdjfalskdjf-asdfaklsjd-sadfkajsdk",
 "state": {
  "name": "Autumn 2017",
  "publishing-state": "DEACTIVATED",
  "ad-locations": [
   ...
  ]
  ...
 },
 "outbox": [
  {
   "campaign-id": "askdjfalskdjf-asdfaklsjd-sadfkajsdk",
   "type": "campaign-deactivated",
   "reason": "Goals met",
   "published": false
  }
 ]
}
```

上記の例では、outboxプロパティに公開が必要なドメインイベントのリストが含まれている

### 未公開イベントの取得

パブリッシングリレーは、プル型、プッシュ型の2つの方法でドメインイベントを取得できる

#### Pull

リレーは未発表のイベントに対して継続的DBに問い合わせる。ポーリングすることで生じるDBへの負荷を最小限にするために、インデックスを最適化する必要有り

#### Push

新しいイベントが追加された時に、パブリッシングリレーを呼び出す。例えば、NoSQLにコミットされた変更をイベントのストリームとして公開する。

### Saga

集約の原則は、各トランザクションを集約の単一のインスタンスに制限し、一貫性ある一連のビジネス機能をカプセル化することだが、複数の集約にまたがるビジネスプロセスを実装する必要性も場合によっては生じる。

- [Sagaパターン参考](https://qiita.com/yasuabe2613/items/b0c92ab8c45d80318420)

#### Sagaを利用する目的

>各サービスがそれぞれの Bounded Context (整合性の境界)で自前のデータストア（Database per Service）を持っているマイクロサービスアーキテクチャで、複数サービスにまたがるワークフローのデータ整合性を維持したい場合にSagaを利用する

#### どのような制約のもとでSagaを利用するか

>以下のような事情で、分散トランザクションは使いたくない場合にSagaを利用する
・ モダンでメジャーな NoSQL やメッセージブローカではサポートされていないものが多い。
・ CAP 定理の認知度が高まって、Consistency を絶対視する風潮が見直され、Availability をより重視するシステムも増えている。
・分散トランザクションは大きな同期プロセス間通信システムとも見ることもでき、可用性を損ねる要因になりうる。
・トランザクションが提供するような「幻想2」に頼らなくても、実際、現実世界は結果整合性で回っている。

#### どんなパターンか

>複数のマイクロサービスのローカルトランザクション（後述）を非同期メッセージングでつなげて、サーガと呼ばれるワークフローを構成するパターン。上の「制約」に書いたとおり分散トランザクションは使わず、Compensation4（補償）と countermeasures で、結果整合性（Eventual consistency）を実現する。

```java
public class CampaignPublishingSaga
{
 private readonly ICampaignRepository _repository;
 private readonly IPublishingServiceClient _publishingService;
 ...
 public void Process(CampaignActivated @event)
 {
  var campaign = _repository.Load(@event.CampaignId);
  var advertisingMaterials = campaign.GenerateAdvertisignMaterials();
  _publishingService.SubmitAdvertisement(@event.CampaignId, advertisignMaterials);
 }
 
 public void Process(PublishingConfirmed @event)
 {
  var campaign = _repository.Load(@event.CampaignId);
  campaign.TrackPublishingConfirmation(@eevnt.ConfirmationId);
  _repository.CommitChanges(campaign);
 }
 
 public void Process(PublishingRejected @event)
 {
  var campaign = _repository.Load(@event.CampaignId);
  campaign.TrackPublishingRejection(@event.RejectionReason);
  _repository.CommitChanges(campaign);
 }
}
```

上記は状態管理を必要としないSagaの例
状態管理が必要となる場合は、Sagaはイベントソースの集合体として実装し、受信したイベントと発行されたコマンドの完全な履歴を保持する。ただし、コマンド実行ロジックは、送信トレイパターンでディスパッチされる方法と同様、saga自体から移動し、非同期で実行する必要がある。

```java

public class CampaignPublishingSaga {
 private readonly ICampaignRepository _repository;
 private readonly IList<IDomainEvent> _events;
 ...
 public void Process(CampaignActivated activated) {
  var campaign = _repository.Load(activated.CampaignId);
  var advertisingMaterials = campaign.GenerateAdvertisingMaterials(); 
  var commandIssuedEvent = new CommandIssuedEvent(
   target: Target.PublishingService,
   command: new SubmitAdvertisementCommand(activated.CampaignId, advertisingMaterials));
  
   _events.Append(activated);
   _events.Append(commandIssuedEvent);
  }
 
 public void Process(PublishingConfirmed confirmed) {
  var commandIssuedEvent = new CommandIssuedEvent(
   target: Target.CampaignAggregate,
   command: new TrackConfirmation(confirmed.CampaignId,
                                   confirmed.ConfirmationId));
  _events.Append(commandIssuedEvent);
    _events.Append(confirmed);
 }
 
 public void Process(PublishingRejected rejected) {
  var commandIssuedEvent = new CommandIssuedEvent(
   target: Target.CampaignAggregate,
   command: new TrackRejection(rejected.CampaignId,
                                    rejected.RejectionReason));
  _events.Append(rejected);
  _events.Append(commandIssuedEvent);
 }
}
```

### 一貫性について

- Sagaはマルチコンポーネントトランザクションを調整するが、2つのトランザクションはアトミックではない。しかし結果生合成は担保する
- 集約の境界内のデータのみが強整合性があるとみなすことができ、他の全ては結果生合成がある

## プロセスマネージャ

Sagaと混同されるパターンとしてプロセスマネージャがあるが、別物。
Sagaは線形フローを管理する。Sagaは、イベントに対応するコマンドに一致させる。

プロセスマネージャはビジネスロジックベースのプロセスを実装することが目的。
シーケンスの状態を維持し、次の処理ステップを決定する中央処理装置として定義される。

![スクリーンショット 2022-05-17 8.34.28.png](:/51fcb7c803d64016bfcbdcfcdd61972b)

TODO: 写経

### 結論

- この章では、システムのコンポーネントを統合するためのパターンを学んだ
- 破損防止レイヤー、オープンホストサービスを実装するために使用できるモデル変換パターンを確認した
- 送信トレイパターンは集約のドメインイベントを公開するための信頼できる方法であることを確認した
- Sagaパターンは、単純なクロスコンポーネントのビジネスプロセスを実装するための方法であることを確認した
- プロセスマネージャパターンは、複雑なビジネスプロセスを実装できることを確認した

# 10章 デザインヒューリスティック

ヒューリスティック：解決法

>「状況によって異なります」はソフトウェアエンジニアリングのほぼ全てに当てはまるが、  実用的ではない。「それ」が何に依存するかを探る必要がある。

## 境界のあるコンテキスト

- 最初はあまりドメイン知識が無い状態でシステム開発をしているから、大きなコンテキスト境界で作業する戦略が良い
  - 物理的な境界（コンテキストを分ける）よりも論理的な境界（同一コンテキスト内のサブコアドメイン）方が変更のコストが低い
  - 1つの変更が複数の境界コンテキストのドメインに影響を与えるような構成では、変更のコストが非常に高くなるので、それは避けたい
  - よって、知識が高まってきたら徐々にコンテキストを細分化するのが良い
  - サイズに意味はない

## ビジネスロジックの実装パターン

トランザクションスクリプト・アクティブレコード・ドメインモデル・イベントソースドメインモデルの4つのビジネスロジック実装パターンについて、もう一度整理する。

### トランザクションスクリプト・アクティブレコード

- サブドメインのサポート、汎用サブドメインのサードパーティソリューションの統合など単純なビジネスロジックを持つサブドメインに適している
- トランザクションスクリプトは単純なデータ構造に向いている
- アクティブレコードは複雑なデータ構造の基盤となるデータベースへのマッピングをカプセル化するのに役立つ

### ドメインモデル・イベントソースドメインモデル

-複雑なビジネスロジックを持つサブドメインに向いている（コアサブドメイン）
 - 金銭的取引や監査ログの出力が法律で定められている場合が向いている
 - システム動作の詳細な分析を必要とするコアサブドメインの場合は、イベントソースドメインモデルが向いている
 
#### ビジネスロジック実装パターンを質問する際の効果的な質問

- サブドメインが金銭やその他金銭取引を追跡・または一貫した監査ログを追跡する
  - イベントソースモデルを使おう
- サブドメインのビジネスロジックが複雑な場合はドメインモデルを実装する
- サブドメインに複雑なデータ構造が含まれている場合はアクティブレコードパターンを使う

![スクリーンショット 2022-05-18 8.37.47.png](:/e40ed2e7ee344f768abcd9de46240659)

#### 複雑なビジネスロジックと単純なビジネスロジック

- 複雑なビジネスロジック
  - 複雑なビジネスルールが含まれる
  - 普遍量・アルゴリズム
  - ユビキタス言語自体が複雑になる
- 単純なビジネスロジック
  - 入力の検証が中心となる

#### 別の観点

- コアサブドメインと見なしているものを、アクティブレコード・トランザクションスクリプトで実装しているなら、それは汎用・サポートサブドメインである可能性が高い

## アーキテクチャパターン

階層化アーキテクチャ・ポート＆アダプタ・CQRSを選択する指針について見る

- イベントソースドメインモデルの場合、CQRSが必要。でなければ、データクエリオプションが極端に制限され、IDのみで単一のインスタンスがfetchされる
- ドメインモデルにはポート＆アダプタが必要。でなければ、集約と値オブジェクトが永続性を認識するのが困難になる。
- アクティブレコードパターンに、アプリケーションサービスレイヤーをもつ階層化アーキテクチャが最適。アクティブなレコードを制御するロジックを管理できる。
- トランザクションスクリプトの場合は、MVCでもいける

## テスト戦略

- ピラミッド
  - 単体テストを重視。集約や値オブジェクトのテストはUTで十分できる
- ダイアモンド
  - 結合テストを重視。アクティブレコードパターンは、ビジネスロジックがサービス層とビジネスロジック層に分散されるので、ダイアモンドテストが有効になる傾向が高い
- 逆ピラミッド
  - E2Eを重視。アプリケーションのワークフローを最初から最後まで検証するので、トランザクションスクリプトを利用する場合は、効果的に検証ができる

![スクリーンショット 2022-05-18 13.33.29.png](:/3731130e85a548038454eb2d7d7589d7)

![スクリーンショット 2022-05-18 14.34.17.png](:/e74898f8b16e4dae8b3e701c95e57cad)

# 11章 進化する設計上の決定

## 概要

- 競合他社の状況や時代の流れによってサブドメインのタイプが変わるよねという話
  - 例えば、ナビタイムは当初スマホで乗り換え案内が検索できることがコアサブドメインだったが、今ではスマホで検索は当たり前なので、それ自体は汎用サブドメインになっている
  - コアサブドメインと他のサブドメインとでは、そもそも実装自体も変わる
    - 自前で用意するものと、既存のソリューションを利用できるもの・開発を他社に依頼できるものという位置付けにおいても、実装がサブドメインのタイプによって異なるのは自明
- 設計変更のパターン
  - トランザクションスクリプトからアクティブレコード
    - データの操作が困難になった場合にアクティブレコードパターンにリファクタする
      - ビジネスロジックは手続型のスクリプトで管理
      - データと永続処理のマッピングをカプセル化する変更を加える
  - アクティブレコードからドメインモデル
    - ビジネスロジックが複雑になり不整合や重複が発生したら、ドメインモデルパターンにリファクタする
      - 値オブジェクトを導入することから始める
      - 次にデータ構造を分析し、トランザクション境界を見つける
        - 全ての状態変更ロジックが明示的であることを確認するには、全てのアクティブレコードのセッターをプライベートに変更する

修正前：

```java
public class Player
{
 public Guid Id { get; set; }
 public int Points { get; set; }
}

public class ApplyBonus
{
 ...
 public void Execute(Guid playerId, byte percentage)
 {
  var player = _repository.Load(playerId);
  player.Points += 1 + percentage/100.0;
  _repository.Save(player);
 }
}
```
   
セッターをprivateに変更：

```java
public class Player
{
 public Guid Id { get; private set; }
 public int Points { get; private set;}
}

public class AppyBounus
{
 ...
   public void Execute(Guid playerId, byte percentage)
   {
    var player = _repository.Load(playerId);
    player.Points *= 1 + percetage/100.0;
    _repository.Save(player);
   }
}

```

次は、ロジックをアクティブレコードの境界内に移動する：

```java
public class Player
{
 public Guid Id { get; private set; }
 public int Points { get; private set; }
 
 public void ApplyBonus(int percentage)
 {
  this.Points += 1 + percentage/100.0;
 }
}

```

### ドメインモデルからイベントソースドメインモデルへ

- 前提：適切にドメインモデルが作られていること・適切な集約の境界が存在すること
- 集約のデータを直接変更する代わりに集約のライフサイクルを表すために飛鳥なドメインイベントをモデル化する
- 移行のハードルは、既存の集約の履歴をどのようにイベントベースのモデルに移行するか
  - 過去のすべての状態変化を表すきめ細かいデータが無いため、ベストエフォートベースで過去のイベントを生成するか、移行イベントをモデル化する必要がある

#### 過去のイベントに対するアプローチ

- 過去の状態遷移の生成
  - 各集約のおおよそのイベントストリームを生成し、それを元の実装と同じ状態表現に投影できるようにする

```java
{
 "lead-id": 12,
 "event-id": 0,
 "event-type": "lead-initialized",
 "first-name": "Shauna",
 "last-name": "Mercia",
 "phone-number": "555-4753"
}, {
 "lead-id": 12,
 "event-id": 1,
 "event-type": "contacted",
 "timestamp": "2020-05-27T12:02:12.51Z"
}, {
 "lead-id": 12, 
 "event-id": 2,
 "event-type": "order-submitted",
 "payment-deadline": "2020-05-30T12:02:12.51Z",
 "timestamp": "2020-05-27T12:02:12.51Z"
 },
{
 "lead-id": 12,
 "event-id": 3,
 "event-type": "payment-confirmed", 
 "status": "converted",
 "timestamp": "2020-05-27T12:38:44.12Z"
}
```

- 上記イベントを1つずつ適用すると、元のシステムと同じように正確な状態表現に投影できる。回復されたイベントは状態を投影して元のデータと比較することでテストも容易になる
- 欠点は、状態遷移の完全な履歴を回復できないということ
- イベントソーシングの目的は、集約のドメインイベントの信頼性が高く、一貫性のある履歴を取得することなので、この目的から外れてしまう

- 過去のイベントに関する知識の欠如を認識するパターン
  - 過去の知識の欠如をイベントとして明示的にモデル化する
  - 現在の状態につながった可能性のあるイベントを回復する代わりに、移行イベントを定義し、それを使用して既存の集約インスタンスのイベントストリームを初期化する

```java

{
 "lead-id": 12,
 "event-id": 0,
 "event-type": "migrated-from-legacy", 
 "first-name": "Shauna",
 "last-name": "Mercia",
 "phone-number": "555-4753",
 "status": "converted",
 "last-contacted-on": "2020-05-27T12:02:12.51Z", 
 "order-placed-on": "2020-05-27T12:02:12.51Z", 
 "converted-on": "2020-05-27T12:38:44.12Z", 
 "followup-on": null
}
```

- メリットは、過去データの欠如を明確にできること
- デメリットは、レガシーシステムの痕跡がイベントスコアに永久に残ること

## 組織変更

- 組織変更に伴うシステム構成変更はあり得る
- 地理的に離れた場合、顧客とサプライヤーの関係に移行するとうまくいく場合がよくある

## ドメイン知識

- 最初はドメイン知識が少ないのでコンテキストは広くなる
- 徐々に知識が溜まってきたタイミングで細分化するのは良いヒューリスティック

## 成長

- サービスの成長によるコードベースのスパゲティ化は設計上の決定を再評価しないことから生じる
  - サブドメイン：ビジネスビルディングブロック
  - 境界コンテキスト：モデル
  - 値オブジェクト：不変性
  - 集約：一貫性
  - これらが正常に機能しているか、影響を調べる必要がある
- 集約
  - 経験上、集約を可能な限り小さくしビジネスドメインによって強力に一貫性のある状態を作る
  - ビジネスの状態に伴い機能追加を便利に行いたくなる誘惑によって、集約オブジェクトに一貫性を保つのに不要なデータが含まれる場合があるが、偶発的な複雑さを避ける意味で、弾くべき

# 12. イベントストーミング

ブレスト形式でドメインイベントを洗い出して、時系列で流れを整理してモデリングをする行為

step1. ブレスト形式で付箋にドメインイベントを書いて張り出していく（重複可）
step2. タイムラインに整理

- ハッピーパスシナリオ（成功するビジネスシナリオ）で始める
step3. 問題点の整理
- ボトルネック、自動化を必要とする手順、ドキュメントの欠落、ドメイン知識の欠落などの明示する
step4. 重要イベントをマーク
- ピボットイベント（コンテキストまたはヘーズの変化を示す重要なビジネスイベント）を探す
- ピボットイベントの前後のイベントを分割する垂直バーで線を引く
step5. コマンドの追加
- ドメインイベントは既に発生したもの（過去）
- コマンド：イベント自体orイベントのフローのトリガー
- コマンドを実行するアクターをここで追加（全てのコマンドにアクターが追加されるわけではない）
step6. ポリシー追加
- 特定の条件下でのみ発行されるコマンドのポリシー（発行条件や対象）を追加
step7. 読み込むモデルの追加
- アクターがコマンドを実行するために使用するドメイン内のデータのViewとなるモデルを追加する
step8. 外部システムの追加
- コマンド実行(input) or イベント通知(output) を行う外部システムを追加する
step9. 集約の追加
- 全てのイベントとコマンドが出揃ったところで、それらを集約で整理する
step10. 境界コンテキストに整理

### リモート

- miro使うと良い
- コミュニケーションは忍耐強く
- 五人程度に参加者を絞らないときつい

# 13章 実世界におけるドメイン駆動

ドメイン駆動設計は少しずつそのツールや要素を取り入れても恩恵を受けられる。
この章では実世界でどのようにドメイン駆動設計を取り入れていくかを確認する。

## 戦略と戦術

ちょっとおさらい

- 戦略：特定の目的を達成するため、対極的な視点で組織行動を計画・遂行する方策、通作。組織が前に進むにはどうしたら良いかを示すもの
- 戦術：戦略を達成するための具体的な手段

## 戦略的分析

まずはビジネスドメインを理解することから始める。

### 最初の問い

- 組織のビジネスドメインは何か
- 顧客は誰か？
- 組織は顧客にどのようなサービス・価値を提供するか？
- 組織はどの企業、製品と競合しているか？

### コアサブドメインの特定

競合他社にはない秘密のソースを見つける必要がある。

- 社内で設計された特許
- アルゴリズム
- 競争上の優位性
などなど、必ずしも技術的なものではなく、他社と差別化できる要因を探す。
また、ジェネリックサブドメインの特定とサポートサブドメインも特定する。

ここでは、全てのこ麻生ドメインを特定する必要はない。
代わりに全体的な構造を特定し、作業しているソフトウェアシスt目うにもっとも関連するサブドメインに注意を払う。

### 現状のシステム設計を調査する

- 高レベルのコンポーネントから理解を進める
- 境界コンテキストレベルではなく、ビジネスドメインをサブシステムに分解するような大きな単位で見ることから進めて良い
- 探すべき特徴的な特性は、コンポーネントの分離されたライフサイクルで、他のコンポーネントから分離・独立してテスト・進化できるかどうかを確認する

### 戦術設計を評価する

- 各高レベルのコンポーネントについてビジネスサブドメインとどのようような技術的設計が行われたか確認する
- その決定は、問題の複雑さに適しているか
- より良いデザインパターンが適用できないか

### 戦略的設計の評価

これまでの調査で見えたコンポーネントの知識を利用して、高レベルのコンポーネントがコンテキスト境界と見なして、コンテキストマップを作る。

そのコンテキストマップを分析し、ドメイン駆動の観点からアーキテクチャを評価する

- 同じ高レベルのコンポーネントに取り組んでいる複数のチームはないか
- コアサブドメインの重複した実装はないか
- アウトソーシング企業にコアサブドメインの実装を任してないか
- 統合が頻繁に失敗するような摩擦はないか
- 外部サービスやレガシーシステムから広がる厄介なモデルはないか

## 近代化戦略

- 大規模な書き直しの取り組みは失敗する可能性が高い
- 大きく考えながら、小さく始めるのが良い
- どこに投資するかを戦略的に決定する必要がある

戦略的にシステムの近代化を進めるために、システムのサブドメインを分割する境界を見る必要がある。これは、技術的な実装パターンの境界ではなく、ビジネスドメインの境界で考える。

そして、この論理的な境界を物理的な境界変えることで、価値が最大化される場所を考える。ただし、システムを可能な限り最小の境界コンテキストに分割することを時期尚早に行うことはリスクを伴うので注意深く行う。

この分割の行為自体は、ロジックや機能の再配置になることが多い。

### 境界コンテキストを抽出するための問いかけ

- 複数チームが同じコードベースで作業してないか？
  - その場合は、各チームの制限されたコンテキストを定義することにより開発ライフサイクルを切り離す
- 競合するモデルがさまざまなコンポーネントで使用されてないか？
  - その場合は、競合数るモデルを個別の境界コンテキストに再配置する
 
#### コンテキスト統合パターンを適切に利用する

- 組織の成長により以前のようにコミュニケーション・コラボレーションできなくなっているならば、コンシューマ・サプライヤパターンに移行したり、腐敗防止層を導入してリファクタする
- レガシーシステムがダウンストリームに非効率なモデルを展開させている場合は、腐敗防止層で吸収する
- 1つのコンポーネントの実装がシステム全体に影響を与えるようであれば、オープンホストサービスを導入し、実装モデルを公開するパブリックAPIから切り離す
- 共有機能が摩擦を起こしており、それがコアサブドメインでないなら、各チームがそれぞれその機能を実装するようにする（セパレートウェイ）

## 戦略的近代化

- ユビキタス言語の育成
- イベントストーミングの実行
- ストラングラーパターンを使った移行

### 戦術設計決定のリファクタリング

レガシーコードを最新化するときの注意点が2つある

#### 1つ目

設計変更のステップを小さく踏むことでリスクを小さくする。
トランザクションスクリプト・アクティブレコードパターンを直接イベントソースのドメインモデルにリファクタすることは避ける（というか、機能的にかなり異なるものになるのでリファクタではなく作り直しになりそう）

#### 2つ目

ドメインモデルへのリファクタリングは、徐々に行う。
一度に全て行おうとすると量が膨大となり現実的ではない。

### 戦術的なデザイン決定のための議論

権威に訴えて決定するのではなく、ロジカルに検討して決定する

- 明示的なトランザクション境界が必要なのはなぜか？
  - データの一貫性を担保するため
- DBトランザクションが複数の集約インスタンスを変更できないのはなぜか？
  - データの整合性の境界を正しく保つため
- 集約オブジェクトの状態を外部コンポーネントから直接変更できないようにするのはなぜか？
  - 関連するビジネスロジックが同じ場所に配置され、重複を避けるため
- 集約の機能の一部をストアドプロシージャにオフロードできないのはなぜか？
  - 物理的に離れたコンポーネントは同期が取れなくなり、データ破損に繋がる傾向があるため
- なぜ小さな集約の境界を目指す必要があるか？
  - トランザクションの範囲が広いと集約の複雑さが増し、パフォーマンスに悪影響を及ぼすため
- イベントソーシングの代わりにイベントログをファイルに書き込むのがNGなのはなぜか？
  - 長期的なデータの一貫性を担保できないため（エッジケースなどで不整合が生じがち）

# 14章 マイクロサービス

The OASIS Service Oriented Architecture Reference Model TC

# 15章 イベント駆動アーキテクチャ

イベント駆動アーキテクチャ（EDA）とDDDの相互作用について説明する。

## イベント駆動型アーキテクチャ

システムのコンポーネントがイベントメッセージを交換することによって非同期的に相互通信するアーキテクチャ

### イベント工藤とイベントソーシングの違い

- イベント駆動は、サービス間の通信を指す
- イベントソーシングは、サービス内で行われる状態変化を一連のイベントとしてキャプチャするための方法
  - イベントソーシングように設計されたイベントは、サービスに実装

# 16章 データメッシュ

分析データ管理システムとドメイン駆動設計の接続の話

## 導入

- ファクトテーブル
- ディメンションモデル

## 分析データ管理プラットフォーム

- DWHの一般的な流れ（エンタープライズ）
  - ETLスクリプトでデータソースからのデータを変換
  - 必要に応じてステージングデータベースに一時保存
  - DWHにロード
  - Userが分析・レポート・マイニング等で利用
    - データマートが部門間などで分かれる場合は、専用のETLプロセスを実装し、運用システムから直接データをロードすることもある
      - 結果、部門間データをクエリする際にパフォーマンスに問題を抱えるなどの辛みが発生する
  - ETLプロセスが分析システムと運用システムとの強力な結合を生みがちでこれも大きな辛み。

### データレイクによる対策

- データレイクに一旦生データを放り込んで、その後ETLで変換してDWHに貯める
  - データレイクはスキーマが無いので、データ量が多くなればなるほど、データの品質が低下し、且つモデリングが複雑になり、データの利用が難しくなりがち

### DDDによるアプローチ

- データメッシュアーキテクチャは分析データのドメイン駆動といえる
- コンテキスト境界に従って、分析モデルの生成基盤の境界を作る
- 境界間の統合はDDDと同様のパターンを利用して考えることができる

# 参考文献

## Advanced DDD

- エヴァンス本
- 実践DDD
- <https://leanpub.com/esversioning/read>
- [Living Documentation: Continuus Knowledge Sharing by Design](https://www.oreilly.com/library/view/living-documentation-continuous/9780134689418/)
  - 知識の共有、文書化、テストに対するDDDのベースアプローチが提案されている
  - 戦略と戦術の詳細な例を提供

## アーキテクチャと統合パターン

- <https://www.oreilly.com/library/view/data-mesh/9781492092384/>
- エンタープライズ アプリケーションアーキテクチャパターン
- <https://www.oreilly.com/library/view/enterprise-integration-patterns/0321200683/>
- マイクロサービスパターン

## レガシーシステムのモダン化

- <https://leanpub.com/arch-modernization-ddd>
- モノリスからマイクロサービスへ

## イベントストーミング

- <https://leanpub.com/introducing_eventstorming>
- <https://leanpub.com/eventstorming_handbook>

# 付録A. DDD適用: ケーススタディ
