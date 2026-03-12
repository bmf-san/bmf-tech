---
title: マークダウン形式のファイルをPDFファイルに変換する（mermaid・emoji・toc対応）
description: マークダウン形式のファイルをPDFファイルに変換する（mermaid・emoji・toc対応）について、基本的な概念から実践的な知見まで詳しく解説します。
slug: markdown-to-pdf-conversion
date: 2022-09-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - markdown
  - marked
  - emoji
  - mermaid
  - JavaScript
translation_key: markdown-to-pdf-conversion
---


# 概要
マークダウン形式のファイルをPDFファイルに変換したいという要望に応えるための簡易的なドキュメント管理ツールを作った。

[bmf-san/docs-md-to-pdf-example](https://github.com/bmf-san/docs-md-to-pdf-example)

特に深く考えることなくあり物のライブラリを活用して作ったので、あまりサステナビリティを感じない構成になっている。

# モチベーション
単にマークダウン形式のファイルをPDFファイルに変換するだけであれば、[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)というライブラリを使うだけで良い。

このライブラリは、レジュメで管理でもお世話になっている。
cf. [Githubでレジュメを管理するようにした](https://bmf-tech.com/posts/Github%e3%81%a7%e3%83%ac%e3%82%b8%e3%83%a5%e3%83%a1%e3%82%92%e7%ae%a1%e7%90%86%e3%81%99%e3%82%8b%e3%82%88%e3%81%86%e3%81%ab%e3%81%97%e3%81%9f)

mermaid記法の対応やunicodeに登録されている絵文字以外の絵文字を使いたかったりという希望があったので、それに対応する形のものを作りたかった。

vscodeの拡張である[vscode-markdown-pdf](https://github.com/yzane/vscode-markdown-pdf)を使えば簡単に解決することができるのだが、vscodeが必要になるので、人によってはvscodeのインストールが必要になってしまう。

変換のためだけにvscodeを使うというのはナンセンスだと思ったので実装してみた。

# 設計
[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)というライブラリは使いやすく素晴らしいライブラリなのだが、現状以下の機能が標準でサポートされていない。

- mermaid記法
- emoji（unicodeに登録されているもの以外）
- TOCの生成

[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)は[markedjs/marked](https://github.com/markedjs/marked)の設定拡張が可能であるため、どれも[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)をカスタマイズすることで実現可能そうではある。

TOCについてはサポートされる予定があるらしい。
[Generate TOC (table of contents) #74](https://github.com/simonhaenisch/md-to-pdf/issues/74)

[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)を使うでも良かったが、やや手間がかかりそうだったので、なるべくハッカソンのような感じで手短に実装したかったので、[md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng)というライブラリを使うことにした。

これは[md-to-pdf](https://github.com/simonhaenisch/md-to-pdf)を拡張してmermaid記法に対応させたライブラリで、あまりメンテナンスされていないようではあるが、一応問題なく使用できる。

[md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng)をベースに、emoji対応は[node-emojify](https://github.com/jesselpalmer/node-emojify)を、TOC生成は[doctoc](https://github.com/thlorenz/doctoc)というライブラリを使って実現する形とした。

# 実装
以下をnpmでインストール。

- [md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng)
- [node-emojify](https://github.com/jesselpalmer/node-emojify)
- [doctoc](https://github.com/thlorenz/doctoc)

※おまけでtextlintを入れているがそのへんは割愛。

emoji対応はmarkedを拡張するような形で対応するので、次のような設定ファイルを用意。

```js
const marked = require('marked');
const { emojify } = require('node-emoji');

const renderer = new marked.Renderer();

renderer.text = emojify;

module.exports = {
	marked_options: { renderer },
};
```

package.jsonのscriptsに次のようなコマンドを定義。

```json
doctoc --notitle md/ && md-to-pdf md/*.md --config-file config.js && mv md/*.pdf pdf/
```

まずdoctocでTOCを生成、次にマークダウンをPDFに変換、最後にディレクトリを移動といった感じ。

md-to-pdfの生成したpdfのアウトプット先をディレクトリ単位で指定できると良いのだがそういうオプションはなさそうだったので、`mv md/*.pdf pdf/`という安直な方法で対応している。

# 所感
この手のものを作ろうとするとやはり結構外部のライブラリに依存しがちになってしまう。
できれば自分で全て実装したいがかなり大変そうに思う。
そのうち機会があればPDFのデータ構造を学んだり、Goで似たようなCLIツールを作ったりしてみたい。
