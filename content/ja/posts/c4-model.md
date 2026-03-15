---
title: C4モデルとは
description: C4モデルとはについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: c4-model
date: 2024-08-15T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - C4モデル
  - アーキテクチャ
translation_key: c4-model
---


# 概要
アーキテクチャ図を書くときにいつも粒度に迷っていたが、C4モデルという技法があることを知ったので、調べてみた。

# C4モデルとは
ソフトウェアアーキテクチャのモデル化技法の一つ。

C4はContext、Containers、Components、Codeの略で、システムをこれらの要素に分解することでソフトウェアアーキテクチャを記述する。

C4モデルは抽象度の高い順に以下の4つのビューを提供する。

- レベル1: システムコンテキスト図
- レベル2: コンテナ図
- レベル3: コンポーネント図
- レベル4: コード図（ex. UMLやERなど）

レベルが上がるほどシステムの内部構造が詳細になる。

# システムコンテキスト図
ソフトウェアシステム単位での図。ソフトウェアシステムとは１つ以上のコンテナを含むもので、ユーザーに価値を提供する構成要素のこと。（ex. アプリケーション、製品、サービスなど）

システムの外部との関係および境界を示し、外部のシステムやユーザーとのインターフェースを表現する。

# コンテナ図
対象のソフトウェアシステムを構成するコンテナの構成を示す図。コンテナはソフトウェアシステム全体が機能するために実行されている要素のこと。（ex. アプリケーションやデータストア）

ソフトウェアアーキテクチャの高レベルな構成と責務を表現する

# コンポーネント図
対象のコンテナを構成するコンポーネントの構成を示す図。コンポーネントはコンテナ内で実行される要素のこと。（ex. クラス、モジュール、サービス）

コンポーネント自身の責務と実装の詳細を表現する。

# コード図
コンポーネント内部のコードを示す図。コード図はUMLやER図などの詳細な図を指す。

コードレベルの詳細を表現する。

# 書き方
原点である[c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/)を熟読して書くのが良さそう。

かなり分かりやすく説明がされているので、これを参照しながら書くと良さそう。

アーキテクチャ図の粒度についての認識を合わせる難しさもこれで解消できるような気がする。

# 所感
もっと早く知っておきたかった技法だった。アーキテクチャ図を書くときに粒度に迷うことが多かったが、C4モデルを使うことで明確な基準に従って図を書くことができそう。

# 参考
- [c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/)
- [ja.wikipedia.org - C4モデル](https://ja.wikipedia.org/wiki/C4%E3%83%A2%E3%83%87%E3%83%AB)
- [www.infoq.com - ソフトウェアアーキテクチャのためのC4モデル](https://www.infoq.com/jp/articles/C4-architecture-model/)
- [namaraii.com - ソフトウェアアーキテクチャのためのC4モデル](https://namaraii.com/notes/c4-model)
	- mermaidでC4モデルの各図を書いている
