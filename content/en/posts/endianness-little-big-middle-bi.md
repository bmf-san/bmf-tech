---
title: Differences Between Little Endian, Big Endian, Middle Endian, and Bi-Endian
slug: endianness-little-big-middle-bi
date: 2020-08-25T00:00:00Z
author: bmf-san
categories:
  - Computer Architecture
tags:
  - Memory
description: An overview of the differences between little endian and big endian.
translation_key: endianness-little-big-middle-bi
---

# Overview
This post summarizes the differences between little endian and big endian.

# What is Endianness?
- The arrangement of multiple bytes is called endianness or byte order.
- It refers to how data is arranged when loaded into memory.
- Endianness is determined by the CPU, protocol, or OS.
- Issues related to endianness often arise when exchanging data between different systems or networks.
  - e.g., Endianness conversion is required when analyzing binary data.

# Big Endian
- A method where data is arranged in memory starting from the *lower address*, with the *most significant byte* placed first, followed by the less significant bytes.
- e.g., Hexadecimal 00 01 02 03 → 00 01 02 03

# Little Endian
- A method where data is arranged in memory starting from the *higher address*, with the *most significant byte* placed first, followed by the less significant bytes.
- e.g., Hexadecimal 00 01 02 03 → 03 02 01 00

# Middle Endian
- A more unconventional method compared to the two above.

# Bi-Endian
- A method that allows switching between big endian and little endian.

# References
- [wikipedia.oorg - エンディアン](https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3)
- [ponsuke-tarou.hatenablog.com - エンディアンは複数バイトのデータを並べる方法です。](https://ponsuke-tarou.hatenablog.com/entry/2017/10/09/224023)
- [uquest.co.jp - Endianってなに？](https://www.uquest.co.jp/embedded/learning/lecture05.html)
- [ertl.jp - バイトオーダ - ビッグエンディアン/リトルエディアン](http://www.ertl.jp/~takayuki/readings/info/no05.html)
- [wa3.i-3-i.info -「ビッグエンディアン」と「リトルエンディアン」の違い](https://wa3.i-3-i.info/diff112endiannes.html#:~:text=%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E4%B8%A6%E3%81%B9%E3%82%8B%E9%A0%86%E7%95%AA%E3%81%8C,%E3%81%8B%E3%82%89%E9%A0%86%E7%95%AA%E3%81%AB%E3%80%8D%E4%B8%A6%E3%81%B9%E3%81%BE%E3%81%99%E3%80%82&text=%E3%81%AE%E3%82%88%E3%81%86%E3%81%AB%E6%9C%80%E5%88%9D%E3%81%8B%E3%82%89%E4%B8%A6%E3%81%B9%E3%81%A6%E7%BD%AE%E3%81%8F%E3%82%84%E3%82%8A%E6%96%B9%E3%81%A7%E3%81%99%E3%80%82,%E3%81%8B%E3%82%89%E9%80%86%E9%A0%86%E3%81%AB%E3%80%8D%E4%B8%A6%E3%81%B9%E3%81%BE%E3%81%99%E3%80%82)
- [xlsoft.com - ビッグエンディアンとリトル・エンディアンの順序](https://jp.xlsoft.com/documents/intel/cvf/vf-html/pg/pg10_01_03_02_01.htm)
- [ap-siken.com - 応用情報技術者平成23年特別 午前問11](https://www.ap-siken.com/kakomon/23_toku/q11.html)
