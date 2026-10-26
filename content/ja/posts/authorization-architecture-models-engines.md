---
title: "認可のアーキテクチャとモデルとは？PEP/PDP・RBAC/ABAC/ReBAC・ポリシーエンジンの選び方"
slug: authorization-architecture-models-engines
date: 2026-10-26
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - 認可
  - RBAC
  - ABAC
  - Open Policy Agent
  - システム設計
description: "認可を「どこで・どう判定するか」を整理する。PEP/PDP/PIP/PAPの参照アーキテクチャと業務ドメインとの位置関係、RBAC/ABAC/ReBACの選び方と組み合わせ方、RBAC標準・NIST SP 800-162・XACML・NGACの教訓、OPA・Cedar・Cerbos・OpenFGA・SpiceDBの比較をまとめる。"
translation_key: authorization-architecture-models-engines
draft: false
---

# 概要

認可基盤の心臓部は「権限をどう表現し、どこで・どう判定するか」である。この記事では、その答えを次の4つの問いに分けて整理する。

1. 判定に関わる部品は何で、業務ドメインに対してどこに置くのか（PEP/PDP/PIP/PAP）
2. 権限を何にもとづいて決めるのか（RBAC/ABAC/ReBACと、PBACの位置づけ）
3. 古典的な標準（RBAC標準、NIST SP 800-162、XACML、NGAC）から今も何を学べるのか
4. ポリシーエンジン（OPA、Cedar、Cerbos、OpenFGA、SpiceDB）をどう選ぶのか

認可とは何か、IAMのどのサブドメインが関わるかは[認可とは何か](/ja/posts/authorization-domain-iam-subdomains/)で、認可基盤をどこまで作るかは[認可基盤はどこまで作るか](/ja/posts/authorization-platform-scope-stages/)で扱った。本記事はその続きで、アクセス判定の中身に踏み込む。ACL・RBAC・ABACの基本は[権限管理の設計](/ja/posts/permission-management-design/)にまとめたので、ここでは繰り返さない。

# 全体像：5つの層に分けて考える

認可の議論では「モデル」「言語」「エンジン」「API」が混ざりやすい。先に層を分けておくと、どの選択がどの選択を縛るかが見える。

```mermaid
flowchart TB
    M["モデル<br/>何にもとづいて許可するか<br/>RBAC / ABAC / ReBAC"]
    L["表現<br/>ポリシーをどう書くか<br/>XACML・ALFA / Rego / Cedar / YAML / 関係スキーマ"]
    E["エンジン（PDP の実装）<br/>OPA / Cedar・AVP / Cerbos / OpenFGA / SpiceDB"]
    A["API（PEP と PDP の線）<br/>各エンジン独自 / AuthZEN"]
    D["データ（PIP）<br/>属性・関係・環境条件の取得元"]
    M --> L --> E
    E --- A
    D --> E
```

| 層 | 決めること | 代表 |
|---|---|---|
| モデル | 許可の根拠を何に置くか | RBAC、ABAC、ReBAC |
| 表現 | ポリシーを書く言語と書き手 | XACML、ALFA、Rego、Cedar、YAML、関係スキーマ |
| エンジン | 評価の実装、配置、データの持ち方 | OPA、Cedar、Cerbos、OpenFGA、SpiceDB |
| API | PEPとPDPの間の要求と応答の形 | 各エンジンのAPI、AuthZEN |
| データ | 判定に使う属性や関係の権威と鮮度 | ディレクトリ、人事、契約、業務DB、関係ストア |

モデルを決めても表現やエンジンは1つに決まらない。逆に、エンジンを選ぶと表現とデータの持ち方はほぼ決まる。だから順序は「モデル → データ → エンジン」が安全である。

# 参照アーキテクチャ：PEP・PDP・PIP・PAP

## 4つの役割

PEPとPDPという呼び名は、IETFのRFC 2753（2000年、ポリシーベースのアドミッション制御の枠組み）が先に定めたものである。用語はPolicy Framework WGのRFC 3198にまとめられた。OASISのXACMLはこの2つを取り込み、PAPとPIPを加えて4つの役割にまとめた。NIST SP 800-162（ABACのガイド）もこの語彙で参照アーキテクチャを示し、NIST SP 800-207（ゼロトラスト）も同じPDPとPEPの語を使う。AuthZENもこの語彙を前提にしている。

```mermaid
flowchart LR
    AUTHOR["ポリシーの書き手"]
    PAP["PAP<br/>管理点"]
    PIP["PIP<br/>情報提供点"]
    PDP["PDP<br/>判定点"]
    CALLER["呼び出し元"]
    PEP["PEP<br/>強制点"]
    RES["資源"]

    AUTHOR -- "作成・テスト・レビュー" --> PAP
    PAP -- "ポリシーの配布" --> PDP
    PIP -- "属性・関係" --> PDP
    PDP <-- "問い合わせ / 許可・拒否（＋義務）" --> PEP
    CALLER -- "要求" --> PEP
    PEP -- "許可された要求だけ" --> RES
```

| 役割 | 何をするか | 置き場所の例 |
|---|---|---|
| PDP（Policy Decision Point） | ポリシーと入力を評価し、許可・拒否を返す。義務（obligation）や助言（advice）を添えることもある | 認可サービス、サイドカー、ライブラリとして埋め込んだエンジン |
| PEP（Policy Enforcement Point） | 要求の経路でPDPに問い合わせ、結果を強制する | APIゲートウェイ、サービス内のミドルウェア、サイドカー、データ層 |
| PIP（Policy Information Point） | 判定に必要な属性や関係をPDPに渡す | ディレクトリ、人事DB、契約DB、業務DB、関係ストア |
| PAP（Policy Administration Point） | ポリシーの作成、テスト、レビュー、版管理、配布を担う | Gitリポジトリ、CI、管理画面 |

## 業務ドメインに対してどこに置くか

PEP/PDP/PIP/PAPは、組織の中で持ち主が違う。[認可とは何か](/ja/posts/authorization-domain-iam-subdomains/)で示したIAMのサブドメインに当てはめると、次のようになる。

```mermaid
flowchart LR
    GW["APIゲートウェイのPEP<br/>粗い判定"]
    subgraph BIZ["業務ドメイン（各サービス）"]
        PEP1["サービス内のPEP<br/>資源単位の判定"]
        DATA["業務データ<br/>資源の属性・所有者"]
    end
    subgraph AD["アクセス判定（横断）"]
        PDP1["PDP"]
        PAP1["PAP"]
    end
    subgraph SRC["属性の権威"]
        IDN["アイデンティティ<br/>主体の属性"]
        GOV["アクセスガバナンス<br/>ロールの割り当て"]
        ADJ["隣接ドメイン<br/>契約・リスクシグナル"]
    end
    GW -- "要求" --> PEP1
    GW -. "判定の問い合わせ" .-> PDP1
    PEP1 -- "判定の問い合わせ" --> PDP1
    PAP1 --> PDP1
    IDN -- "PIP" --> PDP1
    GOV -- "PIP" --> PDP1
    ADJ -- "PIP" --> PDP1
    DATA -- "PIP" --> PDP1
```

ゲートウェイからPDPへの問い合わせを点線にしたのは、粗い判定をゲートウェイの中で済ませる構成もあるためである。

- 資源単位の判定を強制するPEPは、業務ドメインの側にある。資源を守るのは資源の持ち主であり、このPEPを外に出すと迂回路ができる。入口のAPIゲートウェイにもPEPを置けるが、そこで行うのは粗い判定（トークンの有効性やスコープなど）にとどめる。資源単位の判定はサービス内のPEPに任せる。
- PDPとPAPはアクセス判定のサブドメインに属する。ただし、ポリシーの中身をすべて中央で書く必要はない。テナント分離や全社の禁止事項は中央、業務固有のルールは各ドメインのチーム、と書き手を分け、PAPのレビューと配布で束ねる。SP 800-162も、下位組織が上位のポリシーと矛盾しない範囲で独自のポリシーを持つ階層を想定している。
- PIPは「属性の権威」がいる場所につながる。主体の属性はアイデンティティ、ロールの割り当てはアクセスガバナンス、資源の属性は業務ドメインが持つ。PDPが全データを抱えるのではなく、権威から引くか、権威から複製するかを属性ごとに決める。

どこまでを横断の基盤にするかを[認可基盤はどこまで作るか](/ja/posts/authorization-platform-scope-stages/)で、PEPとPDPの配置の詳細を[認可基盤のアーキテクチャと設計判断](/ja/posts/authorization-platform-architecture-decisions/)で扱う。

# 認可モデルの分類

## 3つの独立した軸

本記事では、認可モデルを次の3つの軸に分けて整理する。軸どうしは独立しているので、組み合わせて語る。たとえば「RBACでDACを実現し、ACLとして保存する」という構成もありうる。

| 軸 | 選択肢 | 区別の要点 |
|---|---|---|
| 何にもとづいて許可するか | RBAC、ABAC、ReBAC | ロール、属性と条件、主体と資源の関係のどれを中心にモデリングするか |
| 誰が権限を変えられるか | DAC、MAC | 資源の所有者が裁量で決めるか、中央のラベルで強制するか |
| 権限をどこに持つか | ACL、Capability | アクセス行列を資源ごと（列）に持つか、主体ごと（行）に持つか |

3つ目の軸はLampsonの1971年の論文が示したアクセス行列にさかのぼる。RBACでMACとDACを模倣できることは、Osborn、Sandhu、Munawerの2000年の論文が示している。本記事が主に扱うのは1つ目の軸である。

## RBAC・ABAC・ReBACの比較

| モデル | 判定の基準 | 強い場面 | 弱い場面 |
|---|---|---|---|
| RBAC | ロール（職務に結び付いた権限の束） | 管理と見直しが明快。「誰が何をできるか」を答えやすい | 時刻や場所などの文脈を扱えない。属性の組み合わせをロールにするとロールが爆発する |
| ABAC | 主体・資源・環境の属性と条件 | 細かい条件、組織をまたぐ判定、未知の利用者への対応 | 評価のコスト、ポリシーの複雑さ、事前の見直し（誰が何をできるか）の難しさ |
| ReBAC | 主体と資源の関係のグラフ | 所有者、共有、メンバー、フォルダ継承、組織の階層 | 関係データの一貫性の維持と運用 |

3つは排他ではない。SP 800-162は、ACLとRBACをある意味でABACの特殊な場合とみなす。ACLはidentityという属性、RBACはroleという属性で判定している、という見方である。関係も属性の一種として表せる。違いは「何を中心にモデリングするか」であり、1つのシステムで組み合わせて使える。

## PBACは「書き方」として捉える

PBAC（Policy-Based Access Control）には、定まった1つの定義がない。NISTの用語集も2つの定義を載せている。そのうちCNSSI 4009から引いた定義は、評価する要素（ID、ロール、クリアランス、リスクなど）を柔軟に選べる認可ポリシーを使うアクセス制御の一形態としている。筆者は、PBACを上の3軸とは別の、ポリシー言語で宣言的に書くという表現方法の名前として捉えると、議論を整理しやすいと考える。CNSSI 4009の定義はABACに近いが、RBACとReBACもポリシー言語で書ける。この捉え方に立つと、「RBACとPBACのどちらを選ぶ」という問いは軸が混ざっている。「RBACのモデルを、Policy-as-Codeで書く」のように分けて語ると議論がかみ合う。

# モデルの選び方と組み合わせ方

## どのモデルが要るかを確かめる

3つの問いを順に確かめ、当てはまるものを組み合わせる。

```mermaid
flowchart TD
    Q1{"権限の大半は<br/>組織上の役割で決まるか"}
    Q2{"所有者・共有・メンバー・<br/>親子の継承が本質か"}
    Q3{"時刻・金額・リスク・<br/>端末などの条件が重要か"}
    R["RBAC で表す"]
    B["ReBAC で表す"]
    A["ABAC の条件で表す"]
    E["当てはまったものを組み合わせる"]
    Q1 -- "はい" --> R
    Q1 -- "いいえ" --> Q2
    R --> Q2
    Q2 -- "はい" --> B
    Q2 -- "いいえ" --> Q3
    B --> Q3
    Q3 -- "はい" --> A
    Q3 -- "いいえ" --> E
    A --> E
```

## 「判定」と「一覧」は別の問い

モデルを選ぶとき見落としやすいのが、答えるべき問いの種類である。

- 判定：「このユーザーはこの資源を読めるか」
- 一覧：「このユーザーが読める資源はどれか」「この資源を読めるのは誰か」

一覧の問いは、画面の検索結果の絞り込みや、権限の棚卸しでよく出てくる。NIST SP 800-178は、XACMLのような論理式でポリシーを書く方式では、この種の見直しを効率よく行えないと指摘する。見直しが、NP完全である命題論理の充足可能性問題に相当するためである。属性の条件だけで書いたポリシーに「誰が読めるか」を問うのは本質的に難しい。

同じ文書は、NGACのように関係を列挙して持つ方式なら、この見直しは比較的単純だとする。ReBACも関係をデータとして持つので、この問いに答えやすい。属性中心のエンジンでは、部分評価やクエリプランでポリシーをDBの検索条件に変換して対処する。一覧の要件がある資源は、モデル選びの段階で印を付けておく。

## ロールで上限を決め、条件で絞る

組み合わせ方の指針として、NISTのKuhnらが示した「ロール中心」の方式が今も使える。Kuhn（NIST）、Coyne、Weilの2010年の論文は、RBACに属性を足す方式を3つに分けた。

| 方式 | 権限の決まり方 | 評価 |
|---|---|---|
| 動的ロール | 属性からロールを決め、ロールから権限を決める | 既存のRBACの上に層として足せる |
| 属性中心 | ロールを属性の1つとして扱う | 属性が増えるほどRBACの管理の簡潔さが急速に失われる |
| ロール中心 | ロールで決めた権限を属性の条件で制約する | 属性は権限を減らすだけで増やさない。権限の上限を求められる |

ロール中心の方式では、使える権限は「有効なロールの権限」と「ルールが認める権限」の積集合になる。ロールが上限を決めるので、最小権限を保ち、見直しもしやすい。代わりに、権限の集合がロールに縛られるので、ABACの柔軟性の一部は失われると論文は述べている。論文は、静的な属性（職位、技能、勤務地）をロールにし、動的な属性（時刻など）をルールにする設計も示している。属性が10個（静的4、動的6）なら、純粋なRBACかABACでは最大1,024のロールかルールが要る。分ければ最大16のロールと64のルールで済む。

認可基盤の設計に置き換えると、次の形になる。論文が扱うのはロールと属性だけで、関係の位置づけは筆者の整理である。

- ロールで「この人が最大で何をできるか」を決める。文書の共有のように、それだけで権限を与える関係も上限の側に入る。
- 関係（担当、所有、所属）と属性（時間帯、端末、リスク）の条件で絞る。たとえば、承認者のロールを、自分が担当する部署の申請だけに限る。
- 属性に権限を「足す」ルールは書かない。足したいならロールか、権限を与える関係として表す。

# 原典から今も学べること

古い標準を今そのまま採用する理由は薄い。それでも、その設計判断は新しいエンジンを評価する物差しになる。

## RBAC標準（FK92、NISTモデル、INCITS 359）

RBACは1992年にFerraioloとKuhnが形式化し、NISTのモデルが2004年にANSI/INCITS 359-2004として米国の国家標準になった。2012年に改訂され、359-2012は2022年に再確認されている。NISTのRBACプロジェクトのページは現在アーカイブ扱いで、更新されない。

今も効く要点は次のとおりである。

- すべてのアクセスはロールを通る。NISTのFAQによれば、グループ方式で利用者に直接権限を付けられる点は「抜け穴（loophole）」になりうる。例外の付与も、例外用のロールとして表すほうが見直しやすい。
- 1992年のモデルは「only if」でアクセスを認める。有効なロールに認められたトランザクションしか実行できないが、認められていれば必ず実行できるわけではない。追加の制約をかける余地を最初から残しており、前節のロール中心の方式はこの設計の延長にある。
- セッションでは、割り当てられたロールの一部だけを有効にできる。これは、OAuthのアクセストークンに載せるスコープを要求が必要とする分だけに絞る設計と同じ考え方である。
- 職務分離は2種類ある。静的な職務分離は割り当てのときに効き、動的な職務分離はロールを有効にする時点で効く。「自分の起票を自分で承認しない」のような取引単位の分離は、PDPのルールで扱う。
- 2004年版の承認前のドラフト（2003年、NISTが公開）は、ロール階層の意味（上位を有効にすると下位も自動で有効か）と、ロールを外したときの既存セッションの扱いを実装に委ねている。2012年版の本文は有償で、この扱いが変わったかは公開資料では確かめられない。自分たちの基盤では明文化しておく。

## NIST SP 800-162（ABAC）

SP 800-162は2014年に出たABACのガイドで、2019年の更新で信頼の連鎖の図に環境条件が加わった。ABACの判定の入力を「主体の属性、対象の属性、環境条件、ポリシー」の4つと定義している。

この文書の後半の多くは、ポリシー言語ではなく属性の話に割かれている。扱うのは、属性の権威、共有の方法、鮮度、改ざんの防止、組織間での名前の対応付けである。ここから得られる教訓は、エンジンを選ぶ前に「判定に使う属性の表」を作ることである。属性ごとに、権威（人事、契約、テナント設定など）、PIPへの経路、許容できる古さを書く。

ほかにも次の指摘がある。

- 自然言語のポリシー（NLP）、機械で実行できるポリシー（DP）、ポリシーについてのポリシー（MP）を分ける。NLPからDPへの追跡は、ポリシーのコードに元の規程へのリンクとテストを持たせれば実現できる。
- 属性で判定していても、要求と判定を主体の識別子に結び付けて記録しないと、誰が何にアクセスしたかを追えない。
- 人以外の主体（NPE）の資格情報だけで判定すると、誰のための操作かを追えない。この問題はAIエージェントでそのまま再燃している（[AIエージェントで認可の何が変わるか](/ja/posts/authorization-for-ai-agents/)）。

## XACML 3.0

XACML 3.0は2013年にOASIS Standardとして承認された。2017年には正誤を取り込んだ版が出ている。XMLで書く負担は大きいが、判定の4値、結合アルゴリズム、義務の扱いまで評価の意味論を細かく定めている点は、今も他の方式を比べる物差しになる。

判定は許可・拒否だけでなく、NotApplicable（該当ポリシーなし）とIndeterminate（評価エラー）を加えた4値で返る。複数のルールやポリシーの結果は結合アルゴリズムでまとめる。主なものを次の表に示す。ほかに、評価の順序を固定する変種（ordered-deny-overrides、ordered-permit-overrides）と、ポリシーの結合だけに使うonly-one-applicableがある。

| 結合アルゴリズム | 規則の要点 |
|---|---|
| deny-overrides | 拒否が1つでもあれば拒否 |
| permit-overrides | 許可が1つでもあれば許可 |
| first-applicable | 記載順に評価し、NotApplicable以外になった最初の結果を採る |
| deny-unless-permit | 許可が1つでもあれば許可、それ以外はすべて拒否 |
| permit-unless-deny | 拒否が1つでもあれば拒否、それ以外はすべて許可 |

ここから持ち帰れる設計は次のとおりである。

- fail-closedの作り方。最上位をdeny-unless-permitにし、PEPを「許可で、かつ義務をすべて履行できる場合だけ許可」する拒否寄りの動作にすれば、該当なしとエラーがどちらも拒否になる。
- 義務と助言の区別。PEPは義務を必ず実行し、助言は無視してよい。PEPが理解できない義務や履行できない義務を受けたら、PEPは自分のバイアスに従う。拒否寄りのPEPなら許可しない。AuthZENは、応答の任意（MAY）のcontextに載せる情報の例として義務と助言を挙げるだけで、contextの意味は実装に委ねている。contextを理解できないPEPについても、許可（true）の判定を拒否してよい（MAY）とするにとどまり、XACMLの拒否寄りのPEPのように拒否を求めはしない。AuthZENで義務を返すなら、この取り決めを自分たちで決めておく。順序を決めない結合アルゴリズムでは、どの義務が返るかが決まらない点にも注意する。
- 否定のルールはできるだけ避ける。XACMLの仕様も、慎重に使い、可能なら避けることを勧めている。「〜でなければ許可」は、属性が届かないと許可に倒れることがある。
- 主体の種類を分ける。XACMLは要求を始めた主体、経由した主体、結果を受け取る主体を別のカテゴリで表せる。「ユーザー → エージェント → API」の多段の呼び出しを判定に入れるときの語彙になる。

後継として、XACML TCはXMLに依存しない中核モデルのACAL 1.0を策定している。表現はXML（XACML 4.0）、JSON（JACAL）、YAML（YACAL）に分ける構成である。2026年2月にCommittee Specification Draft 01（中核、XML版、JSON版）が出た段階で、採用を判断するには早い。XACMLを人が書きやすくした記法のALFAもある。2015年にXACML TCで作業草案（Working Draft 01）が出たが、OASISの標準にはなっていない。

## NGAC

NGAC（Next Generation Access Control）は、NISTの研究（Policy Machine）から生まれたABACの標準である。ANSI/INCITSでINCITS 499-2018（機能アーキテクチャ）とINCITS 526-2016（操作とデータ構造）として規格になっている。NIST SP 800-178（2016年）がXACMLと比べている。

違いはポリシーの持ち方にある。XACMLは属性の値を使った論理式でポリシーを書く。NGACは利用者、利用者の属性、対象、対象の属性を節にした割り当てと関連の構成として、ポリシーを列挙する。SP 800-178によれば、NGACは利用者ごと・対象ごとの見直しを効率よくこなせる。一方、論理式の方式ではこの見直しが難しい。その代わり、属性の値ごとに関係を作る必要が生じうるため、関係の数は膨らみやすい。

関係を列挙してグラフとして持ち、そこから「誰が何にアクセスできるか」を引くという考え方は、後のZanzibar系のReBACと通じる。判定と一覧の両方を重視するならグラフで持つ、という選択の根拠として読める。

## SP 800-207（ゼロトラスト）

SP 800-207は、Policy EngineとPolicy AdministratorをまとめてPDPとし、PEPとともにすべてのアクセスを動的に判定するゼロトラストの構成を示した。認可基盤のPDPは、ゼロトラストの文脈では同じ役割を担う。

# ポリシーエンジンの比較

## 一覧

次の表は、2026年10月時点で各プロジェクトの公式サイト、GitHub、CNCFのランドスケープで確かめた情報である。

| エンジン | モデルの中心 | ポリシーの表現 | データの持ち方 | 一覧・絞り込み | 状態・提供形態 |
|---|---|---|---|---|---|
| OPA | ABAC（汎用のポリシー） | Rego | ステートレス。入力とバンドルで配るデータを評価 | 部分評価で検索条件に変換 | CNCF Graduated、Apache-2.0 |
| Cedar | RBAC・ABAC（関係も表せる） | Cedar（permit/forbid） | ステートレス。要求ごとにエンティティを渡す | 部分評価は実験的な機能。主用途は個々の判定 | CNCF Sandbox、Apache-2.0。マネージドはAmazon Verified Permissions |
| Cerbos | RBAC・ABAC | YAML | ステートレスなPDP | クエリプランAPIで検索条件を返す | Apache-2.0。商用の管理面としてCerbos Hub |
| OpenFGA | ReBAC（Zanzibar系） | 認可モデルのDSL＋関係タプル | 関係をストアに持つ | ListObjects、ListUsers | CNCF Incubating、Apache-2.0。商用版にAuth0 FGA |
| SpiceDB | ReBAC（Zanzibar系） | スキーマ言語＋関係 | 関係をストアに持つ | LookupResources、LookupSubjects | Apache-2.0。開発元はAuthZed |

Zanzibarは製品ではなく、Googleが2019年にUSENIX ATCで発表した論文の社内システムである。OpenFGAとSpiceDBはこの論文から着想を得たOSSである。

## 汎用のポリシーエンジン

OPA（Open Policy Agent）は、Kubernetesのアドミッション制御からマイクロサービスの認可、IaCの検査まで使える汎用のポリシーエンジンである。Regoで書いたポリシーをJSONの入力に対して評価する。バンドルによるポリシーの配布、部分評価、判定ログを備える。2021年にCNCFのGraduatedになった。2025年8月には、OPAの作者たちが、自分たちとStyraの多くのメンバーがAppleに移ったとプロジェクトのブログで公表した。CNCFのプロジェクトとしての統治とライセンスは変わらない、と説明されている。OPA自体の解説は[Open Policy Agent（OPA）とは](/ja/posts/open-policy-agent/)、使い方のパターンは[OPAを利用したアクセス制御のパターン](/ja/posts/access-control-patterns-using-opa/)と[OPAを使用したマイクロサービスのアクセス制御システムの実装](/ja/posts/microservices-access-control-opa/)にまとめた。

Cedarは、AWSが公開した認可専用のポリシー言語である。主体・操作・資源・文脈（PARC）に対してpermitとforbidを書く。既定は拒否で、forbidがpermitに優先する。スキーマでポリシーを検証でき、言語の仕様はLeanで形式化され、性質の証明が公開されている。ポリシーを記号的に解析するツールもある。汎用のRegoに対し、Cedarは認可に絞った分だけ書き方が制約され、解析しやすい。CNCFのSandboxに採択されている。Amazon Verified Permissionsは、Cedarを使うマネージドのPDPである。

Cerbosは、YAMLでポリシーを書くステートレスなPDPである。サービスやサイドカーとして動かし、判定のAPIと、検索条件を返すクエリプランのAPIを持つ。ポリシーの配布やテストは商用のCerbos Hubでも管理できる。

## 関係ベースのエンジン（Zanzibar系）

Zanzibarは、Googleの多くの製品が共有する認可システムである。権限を「オブジェクトのある関係に、あるユーザー（またはユーザーの集合）がいる」というタプルで表し、「editorはすべてviewerでもある」のような書き換え規則で関係を継承する。ACLの更新の順序を無視したり、古いACLを新しい内容に当てはめたりすると、外したはずの利用者に新しい内容が見えてしまう。論文はこれを「new enemy」問題と呼ぶ。更新の因果順を保つタイムスタンプ（外部一貫性）と、zookieという一貫性のトークンで鮮度の下限を決めた読み取りで防ぐ。深いグループの入れ子はLeopardという索引で高速化する。

OpenFGAは、Okta（Auth0）の開発者が始めたZanzibar系のOSSで、2022年にCNCFのSandboxに入り、2025年にIncubatingへ移った。DSLで認可モデルを定義し、関係タプルを書き込み、Check、ListObjects、ListUsersで問い合わせる。条件付きの関係タプルで、属性の条件も一部扱える。

SpiceDBは、AuthZedが開発するZanzibar系のOSSである。スキーマ言語で関係と権限を定義し、ZedTokenと要求ごとの一貫性の指定でnew enemy problemを避ける。caveatという仕組みで、条件付きの関係も表せる。

ReBACのエンジンを入れた後も業務DBに所有者や所属などの関係が残るなら、同じ関係データを2か所で持つことになる。その場合、業務の更新と関係の更新をどう同期するか、どちらを正とするかが、導入で特に重い判断になる。

# エンジンの選び方

## 判断の軸

| 軸 | 問い | 傾き |
|---|---|---|
| モデルの中心 | ロールと条件か、関係のグラフか | 前者はOPA・Cedar・Cerbos、後者はOpenFGA・SpiceDB |
| データの持ち方 | 判定に使うデータを要求ごとに渡せるか、エンジンが持つべきか | 渡せるならステートレス、関係が大きく深いならストア型 |
| 一覧の要件 | 「読める資源の一覧」を高速に返す必要があるか | ストア型のReBACが得意。属性型なら部分評価やクエリプランで補う |
| 書き手と検証 | 誰がポリシーを書き、どこまで機械で検査したいか | 汎用性ならRego、解析のしやすさならCedar、設定に近い形ならYAML |
| 用途の広さ | アプリの認可だけか、インフラの制御も含めるか | インフラまで1つの言語でそろえるならOPA |
| 運用 | 自前で動かすか、マネージドを使うか | AWS中心ならAmazon Verified Permissions、ReBACならAuth0 FGAなどの商用版 |
| 差し替え | 将来エンジンを替える余地を残すか | PEPとPDPの間をAuthZENに寄せておく。関係の書き込みやポリシーの配布はAuthZENの対象外である |

## 組み合わせる構成

1つのエンジンにすべてを載せる必要はない。組み合わせの一例は、ReBACのエンジンを関係の権威として持ち、属性の条件はポリシーエンジンで評価する構成である。この場合、ReBACのエンジンはPDPであると同時に、ポリシーエンジンから見たPIPにもなる。ロール中心の考え方に沿うなら、ロールと、権限を与える関係で上限を決め、ポリシーエンジンの条件で絞る。

## 避けたい選び方

- 属性の権威と鮮度を整理する前にエンジンを決める。どのエンジンでも、属性が古いか誤っていれば判定も誤る。
- 業務データをすべてPDPやReBACのストアに複製する。同期の失敗が、そのまま権限の誤りになる。
- エンジン独自のAPIを業務コードのあちこちから直接呼ぶ。PEPの層を1か所にまとめ、エンジンを替えられる形にしておく。

# まとめ

- 認可は「モデル」「表現」「エンジン」「API」「データ」の層に分けて考える。選ぶ順序は、モデル、データ、エンジンが安全である。
- 資源単位の判定を強制するPEPは業務ドメインの側、PDPとPAPはアクセス判定の側、PIPは属性の権威の側に置く。
- RBAC・ABAC・ReBACは排他ではない。ロール（と権限を与える関係）で上限を決め、属性や関係の条件で絞る構成が見直しやすい。PBACは、モデルではなく書き方の名前として捉えると整理しやすい。
- 原典の教訓は今も使える。例はRBACの「only if」、SP 800-162の属性の権威、XACMLの結合アルゴリズムと義務である。NGACからは、一覧の問いに関係の列挙で答えるという考え方を学べる。
- エンジンは、モデルの中心、データの持ち方、一覧の要件で絞り込む。PEPとPDPの間を標準に寄せておけば、選び直しの余地が残る。

# 関連記事

- [認可とは？認証との違いとIAMのサブドメインで見る全体像](/ja/posts/authorization-domain-iam-subdomains/)
- [認可基盤はどこまで作るべきか？横断と各ドメインの分担、判断の軸](/ja/posts/authorization-platform-scope-stages/)
- [AuthZEN Authorization APIとは？PEPとPDPをつなぐ認可判定の標準](/ja/posts/authzen-authorization-api/)
- [認可基盤に必要なOAuth関連標準とは？委譲・トークン・継続評価の地図](/ja/posts/oauth-standards-map-authorization-platform/)
- [ワークロードアイデンティティとは？WIMSE・Transaction Tokens・SPIFFEクライアント認証で見るサービス間の身元と文脈](/ja/posts/workload-identity-wimse-transaction-tokens/)
- [AIエージェントの認可とは？委譲の連鎖・人の承認・監査をMCP、AuthZEN、AIMS、A2A/AP2で整理する](/ja/posts/authorization-for-ai-agents/)
- [MCPの認可はどう設計する？起きうる事故と後から直せない判断](/ja/posts/mcp-authorization-design/)
- [認可基盤のアーキテクチャとは？コンテナ構成・代表シーケンス・PEPとPDPの配置と非機能の設計判断](/ja/posts/authorization-platform-architecture-decisions/)
- [権限管理の設計｜アクセス制御の考え方を整理する](/ja/posts/permission-management-design/)
- [Open Policy Agent (OPA)とは？ポリシー・アズ・コードの実践ガイド](/ja/posts/open-policy-agent/)
- [OPAを利用したアクセス制御のパターン](/ja/posts/access-control-patterns-using-opa/)
- [OPAを使用したマイクロサービスのアクセス制御システムの実装](/ja/posts/microservices-access-control-opa/)
- [PostgreSQLのRow Level Security（RLS）｜行単位のアクセス制御](/ja/posts/postgresql-row-level-security/)

# 参考リンク

- [NIST: Role Based Access Control（プロジェクトページ）](https://csrc.nist.gov/projects/role-based-access-control) / [RBAC FAQ](https://csrc.nist.gov/projects/role-based-access-control/faqs)
- [NIST CSRC Glossary: policy-based access control](https://csrc.nist.gov/glossary/term/policy_based_access_control)
- [D. Ferraiolo, R. Kuhn: Role-Based Access Controls (1992)](https://csrc.nist.gov/pubs/conference/1992/10/13/rolebased-access-controls/final)
- [R. Kuhn, E. Coyne, T. Weil: Adding Attributes to Role-Based Access Control (2010)](https://csrc.nist.gov/pubs/journal/2010/06/adding-attributes-to-rolebased-access-control/final)
- [INCITS 359-2012 (R2022)](https://webstore.ansi.org/standards/incits/incits3592012r2022) / [BSR INCITS 359 ドラフト（2003-04-04、NIST）](https://csrc.nist.gov/CSRC/media/Projects/Role-Based-Access-Control/documents/rbac-std-ncits.pdf)
- [RFC 2753: A Framework for Policy-based Admission Control](https://www.rfc-editor.org/rfc/rfc2753) / [RFC 3198: Terminology for Policy-Based Management](https://www.rfc-editor.org/rfc/rfc3198)
- [NIST SP 800-162: Guide to Attribute Based Access Control (ABAC) Definition and Considerations](https://csrc.nist.gov/pubs/sp/800/162/upd2/final)
- [NIST SP 800-178: A Comparison of Attribute Based Access Control (ABAC) Standards for Data Service Applications](https://csrc.nist.gov/pubs/sp/800/178/final)
- [INCITS 499-2018: Next Generation Access Control - Functional Architecture (NGAC-FA)](https://webstore.ansi.org/standards/incits/incits4992018)
- [INCITS 526-2016: Next Generation Access Control - Generic Operations and Data Structures (NGAC-GOADS)](https://webstore.ansi.org/standards/incits/incits5262016)
- [NIST SP 800-207: Zero Trust Architecture](https://csrc.nist.gov/pubs/sp/800/207/final)
- [OASIS: XACML Version 3.0 Plus Errata 01](https://docs.oasis-open.org/xacml/3.0/errata01/os/xacml-3.0-core-spec-errata01-os-complete.html)
- [OASIS: ACALの成果物](https://docs.oasis-open.org/xacml/acal/)
- [OASIS XACML TC: alfa-for-xacml-v1.0-wd01.doc uploaded（2015-03-10）](https://groups.oasis-open.org/communities/community-home/digestviewer/viewthread?GroupId=19&MessageKey=4d628175-5a72-4e71-9f75-1bd70a4309d1&CommunityKey=67afe552-0921-49b7-9a85-018dc7d3ef1d)
- [B. Lampson: Protection (1971)](https://bwlampson.site/08-Protection/WebPage.html)
- [S. Osborn, R. Sandhu, Q. Munawer: Configuring role-based access control to enforce mandatory and discretionary access control policies (2000)](https://doi.org/10.1145/354876.354878)
- [Zanzibar: Google's Consistent, Global Authorization System (USENIX ATC 2019)](https://research.google/pubs/pub48190/)
- [Open Policy Agent ドキュメント](https://www.openpolicyagent.org/docs/)
- [Note from Teemu, Tim, and Torin to the Open Policy Agent community](https://www.openpolicyagent.org/blog/note-from-teemu-tim-and-torin-to-the-open-policy-agent-community-2dbbfe494371)
- [Cedar](https://www.cedarpolicy.com/) / [cedar-policy/cedar-spec（Leanによる形式化）](https://github.com/cedar-policy/cedar-spec)
- [Amazon Verified Permissions](https://docs.aws.amazon.com/verifiedpermissions/latest/userguide/what-is-avp.html)
- [Cerbos ドキュメント](https://docs.cerbos.dev/)
- [OpenFGA ドキュメント](https://openfga.dev/docs)
- [CNCF: OpenFGA Becomes a CNCF Incubating Project](https://www.cncf.io/blog/2025/11/11/openfga-becomes-a-cncf-incubating-project/)
- [SpiceDB ドキュメント](https://authzed.com/docs)
- [CNCF Landscape](https://landscape.cncf.io/)
- [OpenID AuthZEN Authorization API 1.0](https://openid.net/specs/authorization-api-1_0.html)
