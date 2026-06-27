---
title: "Comparing Real-Time Web Communication: Polling, SSE, WebSocket, WebRTC, and WebTransport"
slug: realtime-web-communication-comparison
date: 2026-06-27
author: bmf-san
categories:
  - Network
tags:
  - Polling
  - SSE
  - WebSocket
  - WebRTC
  - WebTransport
  - Real-time
description: "Five common real-time web communication technologies, Polling, SSE, WebSocket, WebRTC, and WebTransport, compared side by side. This article sorts them by direction and transport, lays out each one's overview, minimal implementation, and pros and cons, and gives selection criteria. A map for choosing a real-time method."
translation_key: realtime-web-communication-comparison
draft: false
---

# Introduction
On the web, you often need to push data from server to client right away.

Chat, notifications, stock prices, multiplayer games, and video calls all demand real-time delivery.

This article compares five common real-time web communication technologies side by side.

- Polling
- SSE (Server-Sent Events)
- WebSocket
- WebRTC
- WebTransport

For each, it lays out an overview, a minimal implementation, and pros and cons, then closes with selection criteria.

For a broader comparison of API styles, read [Comparing API Styles](https://bmf-tech.com/posts/api-styles-comparison). This article drills into the real-time slice.

# Classification Axes
Before the comparison, here are the axes.

- Direction: server-to-client one-way, or bidirectional
- Transport: rides on HTTP, or a dedicated protocol (such as UDP)
- Main use: notifications and feeds, two-way interaction, or P2P media exchange

# Polling
The client asks the server at a fixed interval and picks up updates. It ranks as the most basic approach.

Short polling repeats at a fixed interval; long polling holds the response open until an update arrives.

```js
// short polling
setInterval(async () => {
  const res = await fetch("/api/messages");
  const data = await res.json();
  render(data);
}, 3000);
```

## Pros / Cons
- Pros: needs no special machinery and runs on plain HTTP. Wide compatibility
- Cons: wasted requests pile up, and efficiency drops. Real-time feel trades off against load

# SSE (Server-Sent Events)
The server streams events to the client continuously over a single HTTP connection, one-way.

```js
const es = new EventSource("/api/stream");
es.onmessage = (e) => render(JSON.parse(e.data));
```

## Pros / Cons
- Pros: runs over HTTP and stays simple. Built-in automatic reconnection. Ideal for notifications and feeds
- Cons: server-to-client only. Binary is awkward

# WebSocket
Over a single connection, client and server exchange messages both ways.

It upgrades from HTTP and holds a persistent connection.

```js
const ws = new WebSocket("wss://example.com/socket");
ws.onmessage = (e) => render(e.data);
ws.onopen = () => ws.send("hello");
```

## Pros / Cons
- Pros: bidirectional and low-latency. Handles text and binary. Fits chat and multiplayer games
- Cons: holding the connection makes state and scaling hard. You build reconnection yourself

# WebRTC
WebRTC exchanges audio, video, and data peer-to-peer between browsers at low latency.

Establishing a connection needs out-of-band signaling (exchanging SDP and ICE candidates).

```js
const pc = new RTCPeerConnection();
const channel = pc.createDataChannel("chat");
channel.onmessage = (e) => render(e.data);
// then exchange SDP and ICE candidates through a signaling server
```

## Pros / Cons
- Pros: P2P with ultra-low latency. Ideal for audio and video. Communicates without a server in the middle
- Cons: signaling and NAT traversal (STUN/TURN) get complex. High implementation difficulty

# WebTransport
WebTransport is a newer API that runs over HTTP/3 (QUIC), bidirectional and multiplexed.

People see it as a successor to WebSocket; it offers both reliable streams and unordered datagrams.

```js
const wt = new WebTransport("https://example.com:4433/wt");
await wt.ready;
const stream = await wt.createBidirectionalStream();
```

## Pros / Cons
- Pros: QUIC-based and low-latency. Avoids head-of-line blocking. Lets you pick streams or datagrams
- Cons: new, with limited support. Server-side implementation and operations remain early

# A Cross-Cutting Comparison
| Aspect | Polling | SSE | WebSocket | WebRTC | WebTransport |
| --- | --- | --- | --- | --- | --- |
| Direction | one-way (pull) | one-way (push) | bidirectional | bidirectional (P2P) | bidirectional |
| Transport | HTTP | HTTP | TCP (HTTP upgrade) | UDP (P2P) | HTTP/3 (QUIC) |
| Browser support | all | wide | wide | wide | limited |
| Reconnection | manual | automatic | manual | manual | manual |
| Main use | simple update checks | notifications, feeds | chat, games | calls, video | high-volume, low-latency |
| Complexity | low | low | medium | high | medium-high |

# Selection Criteria
Here is guidance that maps needs to a technology.

- Pick up occasional updates simply: Polling
- Push notifications or feeds one-way: SSE
- Two-way interaction or games: WebSocket
- Audio, video, or P2P ultra-low latency: WebRTC
- Maximize high-volume, low-latency on a modern stack: WebTransport

When in doubt, start from SSE or WebSocket. If one-way suffices, SSE is easy; if you need two-way, WebSocket is easy.

# Summary
Real-time web communication sorts by direction and transport.

- HTTP-based: Polling, SSE, WebSocket
- Dedicated transport: WebRTC (UDP/P2P), WebTransport (QUIC)

Pick by need: Polling or SSE for simplicity, WebSocket for two-way, WebRTC for media, WebTransport for the cutting edge.

# References
- RFC 6455: The WebSocket Protocol. https://www.rfc-editor.org/rfc/rfc6455
- HTML Living Standard: Server-sent events. https://html.spec.whatwg.org/multipage/server-sent-events.html
- W3C: WebRTC 1.0: Real-Time Communication Between Browsers. https://www.w3.org/TR/webrtc/
- W3C: WebTransport. https://www.w3.org/TR/webtransport/
- RFC 9000: QUIC: A UDP-Based Multiplexed and Secure Transport. https://www.rfc-editor.org/rfc/rfc9000

