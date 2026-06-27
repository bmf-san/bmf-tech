---
title: "リアルタイムWeb通信の比較：ポーリング・SSE・WebSocket・WebRTC・WebTransport"
slug: realtime-web-communication-comparison
date: 2026-06-27
author: bmf-san
categories:
  - ネットワーク
tags:
  - ポーリング
  - SSE
  - WebSocket
  - WebRTC
  - WebTransport
  - リアルタイム
description: "ポーリング・SSE・WebSocket・WebRTC・WebTransportという5つのリアルタイムWeb通信技術を、通信方向やトランスポートで分類し、概要・最小実装・pros/consと選定基準を横断的にまとめる。実時間通信の方式を選ぶときの地図となる記事。"
translation_key: realtime-web-communication-comparison
draft: false
---

# はじめに
Webでサーバーからクライアントへ「すぐに」データを届けたい場面は多い。

チャット、通知、株価、対戦ゲーム、ビデオ通話など、リアルタイム性が求められる用途はさまざまである。

本記事では、代表的な5つのリアルタイムWeb通信技術を横断的に比較する。

- ポーリング
- SSE（Server-Sent Events）
- WebSocket
- WebRTC
- WebTransport

それぞれの概要・最小実装・pros/consを整理し、最後に選定基準を示す。

なお、API全体のスタイル比較は[各種APIスタイルの比較](https://bmf-tech.com/ja/posts/api-styles-comparison)にまとめている。本記事はそのうちリアルタイム通信を掘り下げる位置づけである。

# 分類軸
比較の前に、整理のための軸を示す。

- 通信方向: サーバーからの単方向か、双方向か
- トランスポート: HTTP上に乗るか、独自プロトコル（UDPなど）か
- 主な用途: 通知・フィード配信か、双方向の対話か、P2Pのメディア交換か

# ポーリング
クライアントが一定間隔でサーバーに問い合わせ、更新を取得する最も素朴な方式である。

一定間隔で繰り返すショートポーリングと、サーバーが更新まで応答を保留するロングポーリングがある。

```js
// ショートポーリング
setInterval(async () => {
  const res = await fetch("/api/messages");
  const data = await res.json();
  render(data);
}, 3000);
```

## pros / cons
- 利点: 特別な仕組みが不要で、HTTPだけで実装できる。互換性が高い
- 欠点: 無駄なリクエストが多く効率も悪い。リアルタイム性と負荷のトレードオフになる

# SSE（Server-Sent Events）
サーバーからクライアントへ、1本のHTTP接続で継続的にイベントを送り続ける単方向の仕組みである。

```js
const es = new EventSource("/api/stream");
es.onmessage = (e) => render(JSON.parse(e.data));
```

## pros / cons
- 利点: HTTP上で動き実装が簡単。自動再接続が標準で備わる。通知やフィードに最適
- 欠点: サーバーからクライアントへの単方向のみ。バイナリは扱いにくい

# WebSocket
1本のコネクション上で、クライアントとサーバーが双方向にメッセージを送り合う仕組みである。

HTTPからアップグレードして常時接続を張る。

```js
const ws = new WebSocket("wss://example.com/socket");
ws.onmessage = (e) => render(e.data);
ws.onopen = () => ws.send("hello");
```

## pros / cons
- 利点: 双方向・低遅延。テキストとバイナリの両方を扱える。チャットや対戦ゲームに向く
- 欠点: 接続の維持で状態管理やスケールが難しい。再接続は自前で実装する

# WebRTC
ブラウザ間でP2Pに、音声・映像・データを低遅延でやり取りする仕組みである。

接続確立には、別経路でのシグナリング（SDPやICE候補の交換）が必要になる。

```js
const pc = new RTCPeerConnection();
const channel = pc.createDataChannel("chat");
channel.onmessage = (e) => render(e.data);
// この後、SDP と ICE 候補をシグナリングサーバー経由で交換する
```

## pros / cons
- 利点: P2Pで超低遅延。音声・映像に最適。サーバーを介さずに通信できる
- 欠点: シグナリングやNAT越え（STUN/TURN）の設計が複雑。実装の難度が高い

# WebTransport
HTTP/3（QUIC）上で動く、双方向かつ多重化された新しい通信APIである。

WebSocketの後継と目され、信頼性のあるストリームと、順序を保証しないデータグラムの両方を扱える。

```js
const wt = new WebTransport("https://example.com:4433/wt");
await wt.ready;
const stream = await wt.createBidirectionalStream();
```

## pros / cons
- 利点: QUICベースで低遅延。ヘッドオブラインブロッキングを回避。ストリームとデータグラムを使い分けられる
- 欠点: 新しく対応環境が限られる。サーバー側の実装・運用がまだ発展途上

# 横断比較表
| 観点 | ポーリング | SSE | WebSocket | WebRTC | WebTransport |
| --- | --- | --- | --- | --- | --- |
| 方向 | 単方向（pull） | 単方向（push） | 双方向 | 双方向（P2P） | 双方向 |
| トランスポート | HTTP | HTTP | TCP（HTTPから昇格） | UDP（P2P） | HTTP/3（QUIC） |
| ブラウザ対応 | 全環境 | 広い | 広い | 広い | 限定的 |
| 再接続 | 自前 | 標準で自動 | 自前 | 自前 | 自前 |
| 主な用途 | 簡易な更新確認 | 通知・フィード | チャット・ゲーム | 通話・映像 | 大容量・低遅延 |
| 実装の複雑さ | 低 | 低 | 中 | 高 | 中〜高 |

# 選定基準
要件から技術を選ぶときの目安を示す。

- とにかく手軽に、たまの更新を取れればよい: ポーリング
- サーバーからの通知やフィードを流したい（単方向）: SSE
- 双方向の対話やゲームが必要: WebSocket
- 音声・映像やP2Pの超低遅延が必要: WebRTC
- 最新環境で大容量・低遅延を最大化したい: WebTransport

迷ったら、まずSSEかWebSocketを基準に考えるとよい。単方向で足りるならSSE、双方向ならWebSocketが扱いやすい。

# まとめ
リアルタイムWeb通信は、通信方向とトランスポートで整理できる。

- HTTPベース: ポーリング・SSE・WebSocket
- 専用トランスポート: WebRTC（UDP/P2P）・WebTransport（QUIC）

シンプルさならポーリングやSSE、双方向ならWebSocket、メディアならWebRTC、最先端ならWebTransport、と要件で選び分けるとよい。

# 参考
- RFC 6455: The WebSocket Protocol. https://www.rfc-editor.org/rfc/rfc6455
- HTML Living Standard: Server-sent events. https://html.spec.whatwg.org/multipage/server-sent-events.html
- W3C: WebRTC 1.0: Real-Time Communication Between Browsers. https://www.w3.org/TR/webrtc/
- W3C: WebTransport. https://www.w3.org/TR/webtransport/
- RFC 9000: QUIC: A UDP-Based Multiplexed and Secure Transport. https://www.rfc-editor.org/rfc/rfc9000

