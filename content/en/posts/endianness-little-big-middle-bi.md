---
title: Differences Between Little Endian, Big Endian, Middle Endian, and Bi-endian
description: 'Understand byte ordering: Little Endian, Big Endian, Middle Endian, Bi-endian, and memory arrangement across systems.'
slug: endianness-little-big-middle-bi
date: 2020-08-25T00:00:00Z
author: bmf-san
categories:
  - Computer Architecture
tags:
  - Memory
translation_key: endianness-little-big-middle-bi
---

# Overview
This article summarizes the differences between Little Endian and Big Endian.

# What is Endian?
- The method of arranging multiple bytes is called Endian or byte order.
- It refers to how data is arranged when loaded into memory.
- Endian is determined by the CPU, protocol, and OS.
- Problems related to Endian can easily arise when exchanging data between different systems or networks.
  - e.g. Endian conversion is necessary when analyzing binaries.

# Big Endian
- A method where data is arranged from the *lower address* in the order of *higher bytes* to lower bytes.
- e.g. Hexadecimal 00 01 02 03 → 00 01 02 03

# Little Endian
- A method where data is arranged from the *higher address* in the order of *higher bytes* to lower bytes.
- e.g. Hexadecimal 00 01 02 03 → 03 02 01 00

# Middle Endian
- A more irregular method than the above two.

# Bi-endian
- A method that switches between Big Endian and Little Endian.

# References
- [wikipedia.oorg - Endian](https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3)
- [ponsuke-tarou.hatenablog.com - Endian is a method of arranging multiple byte data.](https://ponsuke-tarou.hatenablog.com/entry/2017/10/09/224023)
- [uquest.co.jp - What is Endian?](https://www.uquest.co.jp/embedded/learning/lecture05.html)
- [ertl.jp - Byte Order - Big Endian/Little Endian](http://www.ertl.jp/~takayuki/readings/info/no05.html)
- [wa3.i-3-i.info - Differences Between Big Endian and Little Endian](https://wa3.i-3-i.info/diff112endiannes.html#:~:text=%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E4%B8%A6%E3%81%B9%E3%82%8B%E9%A0%86%E7%95%AA%E3%81%8C,%E3%81%8B%E3%82%89%E9%A0%86%E7%95%AA%E3%81%AB%E3%80%8D%E4%B8%A6%E3%81%B9%E3%81%BE%E3%81%99%E3%80%82&text=%E3%81%AE%E3%82%88%E3%81%86%E3%81%AB%E6%9C%80%E5%88%9D%E3%81%8B%E3%82%89%E4%B8%A6%E3%81%B9%E3%81%A6%E7%BD%AE%E3%81%8B%E3%82%8A%E6%96%B9%E3%81%A7%E3%81%99%E3%80%82,%E3%81%8B%E3%82%89%E9%80%86%E9%A0%86%E3%81%AB%E3%80%8D%E4%B8%A6%E3%81%B9%E3%81%BE%E3%81%99%E3%80%82)
- [xlsoft.com - Order of Big Endian and Little Endian](https://jp.xlsoft.com/documents/intel/cvf/vf-html/pg/pg10_01_03_02_01.htm)
- [ap-siken.com - Applied Information Technology Engineer Exam Heisei 23 Special Morning Question 11](https://www.ap-siken.com/kakomon/23_toku/q11.html)