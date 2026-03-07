---
title: Conversion Between Binary, Decimal, and Hexadecimal
slug: binary-decimal-hexadecimal-conversion
date: 2018-11-27T00:00:00Z
author: bmf-san
categories:
  - Mathematics
tags:
  - Base Conversion
  - Discrete Mathematics
translation_key: binary-decimal-hexadecimal-conversion
---

# Overview
This article summarizes the calculation methods for converting between binary, decimal, and hexadecimal.

# Weight
Before performing calculations, it is essential to understand the concept of weight.

Weight refers to the value represented by each digit.

ex. Decimal 1234

10^0*4 = 4
10^1*3 = 30
10^2*2 = 200
10^3*1 = 1000
     sum 1234

10^0, 10^1, 10^2... are the weights.

ex. Binary 1101
2^0*1 = 1
2^1*0 = 0
2^2*1 = 4
2^3*1 = 8
    sum 13

2^0, 2^1, 2^2... are the weights.

# Base
The base for binary is 2, for decimal is 10, and for hexadecimal is 16.

# Binary to Decimal
Multiply the weights by the respective digit values and sum them all.

ex. 1010
2^0*0 = 0
2^1*1 = 2
2^2*0 = 0
2^3*1 = 8
    sum 10

# Decimal to Binary
Perform a unique division. To convert from decimal to binary, divide by 2, and if there is a remainder, it is 1; if not, it is 0. Finally, arrange the remainders from the last calculation.

```
ex. 100
100/2 = 50 remainder 0
50/2  = 25 remainder 0
25/2  = 12 remainder 1
12/2  = 6  remainder 0
6/2   = 3  remainder 1
3/2   = 1  remainder 1  // Stop when the last is 1

Arrange from the bottom 110100

Ans. 110100
```

# Binary to Octal
Three binary digits represent 2^3=8. To convert binary to octal, group the digits in sets of three and calculate. Finally, arrange the results from each group.

```
ex. 100100

100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
    sum 4

100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
    sum 4

Ans. 44
```

```
ex. 1100
100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
    sum 4

1
    sum 1

Ans. 14
```

# Octal to Binary
Represent each digit as a 3-digit binary number. Finally, omit leading zeros.

```
ex. 117
7 → 111
1 → 001
1 → 001

001001111 → 100111

Ans. 100111
```

# Binary to Hexadecimal
Four binary digits represent 2^4=16. To convert binary to decimal, group the digits in sets of four and calculate. Finally, arrange the results from each group.

Hexadecimal
0 1 2 3 4 5 6 7 8 9 A B C D E F 10
Decimal
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16

```
ex. 11001100
1100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
2^3*1 = 8
    sum 12 → C

1100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
2^3*1 = 8
    sum 12 → C

Ans. CC
```

```
ex. 1100
1100
2^0*0 = 0
2^1*0 = 0
2^2*1 = 4
2^3*1 = 8
    sum 12 → C

Ans. C
```

# Hexadecimal to Binary
Represent each digit as a 4-digit binary number. Finally, omit leading zeros.

```
ex. 8B6
6 → 0110
B → 1011
8 → 1000

Ans. 100010110110
```

# References
- [Base Conversion: Methods for Bidirectional Conversion Between Decimal, Binary, and Hexadecimal](http://share-answers.com/category1/entry4.html)
- [Surprisingly Easy Once You Remember the Tips! How to Convert Binary to Decimal, Octal, and Hexadecimal](http://challengdirector.hatenablog.com/entry/2017/06/25/154349)