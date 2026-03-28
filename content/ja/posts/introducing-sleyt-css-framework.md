---
title: "sleyt — グラスモーフィズムを活用したミニマルCSSフレームワークの紹介"
description: '純粋CSS実装のミニマルフレームワーク『sleyt』の紹介。グラスモーフィズムデザイン、データ可視化コンポーネント（棒グラフ・折れ線グラフ・ドーナツチャート）、豊富なコンポーネントライブラリ、完全なダークモード対応をJavaScriptランタイム不要で提供。'
slug: introducing-sleyt-css-framework
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - CSS
  - Frontend
translation_key: introducing-sleyt-css-framework
---

# sleyt — グラスモーフィズムを活用したミニマルCSSフレームワークの紹介

## なぜ作ったか

自分のプロジェクトで使うための軽量でシンプルなCSSフレームワークが欲しかったので自作した。

軽量で、JavaScriptなし、モダンなデザインを意識して実装している。

UIコンポーネントや柔軟性が足りていない部分があるが、最低限利用できる状態になっている。

![ドキュメントサイト](/assets/images/posts/introducing-sleyt-css-framework/01_docs_home.png)

## コンポーネントライブラリ

sleytは4つのディレクトリに20以上のコンポーネントを整備する。

<!-- textlint-disable ja-technical-writing/sentence-length -->
**Components** (`src/components/`): accordion、alerts、badges、buttons、cards、charts、code、forms、modals、navbar、navigation、progress、prose、sidebar、spinners、showcase、swatch、tables、tabs、tooltip。
<!-- textlint-enable ja-technical-writing/sentence-length -->

**Layout** (`src/layout/`): container、dashboard、flexbox、grid。

**Base** (`src/base/`): reset、themes、variables。

**Utilities** (`src/utilities/`): spacing、colors、typography、borders、effects、glass、display、position、transforms、transitions。

![コンポーネント一覧](/assets/images/posts/introducing-sleyt-css-framework/06_demo_components.png)

![ブログデモ](/assets/images/posts/introducing-sleyt-css-framework/03_demo_blog.png)

### データ可視化

`charts.css`は**CSSだけ**の棒グラフ、折れ線グラフ、ドーナツチャートを提供する。JavaScriptのグラフライブラリ不要。棒グラフはCSSの高さでデータを表現する純CSS実装だ。折れ線グラフはSVG `<path>`要素をCSSでスタイリングし、ドーナツチャートはSVG `<circle>`の`stroke-dasharray`と`stroke-dashoffset`でセグメントを描画する。データ対応はCSSカスタムプロパティ経由で行う。

![ダッシュボードデモ](/assets/images/posts/introducing-sleyt-css-framework/05_demo_dashboard.png)

![チャートデモ](/assets/images/posts/introducing-sleyt-css-framework/08_demo_charts.png)

### グラスモーフィズム

`.glass`、`.glass-light`、`.glass-heavy`、`.frosted`のユーティリティクラスで透明感とバックドロップブラーを付与できる。`backdrop-blur`ユーティリティも複数段階用意されており、細かいブラー強度の調整が可能だ。

![グラスモーフィズムユーティリティ](/assets/images/posts/introducing-sleyt-css-framework/07_docs_glass.png)

### ダークモード

`@media (prefers-color-scheme: dark)`によりダークモードが自動的に切り替わる。全コンポーネントの色はCSSカスタムプロパティで定義されており、1ファイル変更でパレットをカスタマイズできる。

### アクセシビリティ

コンポーネントのマークアップはセマンティックHTML5パターンに準拠する。ボタンはネイティブの`<button>`要素、ナビゲーションは`<nav>`を使用する。色彩コントラストはWCAG AAを目標としている。

## インストール

```bash
npm install sleyt
```

CSSでインポートするだけで使い始められる。

```css
@import "sleyt/dist/css/index.css";
```

CDN経由で使う場合は、HTMLの`<head>`に直接追加できる。

```html
<link rel="stylesheet" href="https://unpkg.com/sleyt@latest/dist/css/index.css">
```

## デモとドキュメント

[デモページ](https://bmf-san.github.io/sleyt/demo.html)で全コンポーネントをダークモードバリアントやグラフ種別も含めて確認できる。ブログレイアウト・記事詳細・管理ダッシュボードの3つの実用的なデモページも用意されており、実際のUIパターンを確認できる。

![記事詳細デモ](/assets/images/posts/introducing-sleyt-css-framework/04_demo_blog_detail.png)

[ドキュメントサイト](https://bmf-san.github.io/sleyt/)はインストール、CSS変数によるカスタマイズ、コンポーネント使用例を解説する。

## まとめ

sleytは読みやすさとモダンなビジュアルスタイルに特化した軽量CSS専用フレームワークである。JavaScriptランタイムコストなし。

- **npm**: [sleyt](https://www.npmjs.com/package/sleyt)
- **GitHub**: [bmf-san/sleyt](https://github.com/bmf-san/sleyt)
- **デモ**: [bmf-san.github.io/sleyt/demo.html](https://bmf-san.github.io/sleyt/demo.html)
