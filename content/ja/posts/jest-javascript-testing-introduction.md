---
title: Jestで始めるJavaScriptのテスト
description: Jestで始めるJavaScriptのテスト
slug: jest-javascript-testing-introduction
date: 2018-09-20T00:00:00Z
author: bmf-san
categories:
  - テスト
tags:
  - ES5
  - JavaScript
  - babel
  - babel-jest
  - ESModules
  - jest
translation_key: jest-javascript-testing-introduction
---


# 概要
Jestを使ってJavaScriptのテストをかいてみる。

# 準備
jestとESModulesを使いたいのでbabel-preset-2015をインストールしておく。
（babel-jestはjestに用意されている。）

`npm install --save-dev jest babel-preset-2015`

`.babelrc`の中身はこんな感じ。

```
{
  "presets": ["es2015"]
}
```

`package.json`はこんな感じ。

```
{
  "scripts": {
    "test": "jest"
  },
  "devDependencies": {
    "babel-preset-es2015": "^6.24.1",
    "jest": "^23.6.0"
  }
}
```

ディレクトリ構成はこんな感じ。
`tree -a -I "node_modules"`

```
.
├── .babelrc
├── package-lock.json
├── package.json
├── src
│   ├── esmodules
│   │   └── calc.js
│   └── native
│       └── calc.js
└── test
    ├── esmodules
    │   └── calc.test.js
    └── native
        └── calc.test.js

6 directories, 7 files
```

テストファイルの作成方法については2パターンある。
- `__tests__`という名前のディレクトリ以下に存在するファイルをテストファイルとするパターン
- `*.spec.js`または`*.test.js`を拡張子とするファイルをテストファイルとするパターン

今回は後者の形式をとり、`test`ディレクトリにテストファイルを設置する。

# ネイティブのJavaScriptのテストをかいてみる
足し算引き算をする関数を実装する。

`./src/native/calc.js`

```javascript
const counter = 1

const add = function add(num) {
  return counter + num
}

const subtract = function subtract(num) {
  return counter - num
}

module.exports = {
  add, subtract
}
```

それぞれの関数が正しい計算結果を返すかテストする。　

`./test/native/calc.test.js`

```javascript
const calc = require("../../src/native/calc")

describe('Calc - native', () => {
  test('add', () => {
    expect(calc.add(1)).toBe(2)
  })

  test('subtract', () => {
    expect(calc.subtract(1)).toBe(0)
  })
})
```

`describe(name, fn)`は複数のテストをテストスイートしてグループ化したブロックを作成する。

テストを実行。
`npm test ./test/native/calc.test.js`

テスト結果。
```
> @ test /Users/k.takeuchi/localdev/project/til/javascript/test/jest
> jest "./test/native/calc.test.js"

 PASS  test/native/calc.test.js
  Calc - native
    ✓ add (2ms)
    ✓ subtract (1ms)

Test Suites: 1 passed, 1 total
Tests:       2 passed, 2 total
Snapshots:   0 total
Time:        0.887s, estimated 1s
Ran all test suites matching /.\/test\/native\/calc.test.js/i.
```

# ESModulesを使ったJavaScriptのテストをかいてみる
足し算引き算をするメソッドを実装したクラスを作成する。

`./src/esmodules/calc.js`

```javascript
export default class Calc {
  constructor(counter) {
    this.counter = counter
  }

  add(num) {
    this.counter += num

    return this.counter
  }

  subtract(num) {
    this.counter -= num

    return this.counter
  }
}
```

それぞれのメソッドが正しい計算結果を返すかテストする。　

`./test/esmodules/calc.test.js`

```javascript
import Calc from "../../src/esmodules/calc"

describe('Calc - esmodules', () => {
  test('add', () => {
    const calc = new Calc(1)
    expect(calc.add(1)).toBe(2)
  })

  test('subtract', () => {
    const calc = new Calc(1)
    expect(calc.subtract(1)).toBe(0)
  })
})
```

# マッチャー
[Jest - Expect](https://jestjs.io/docs/ja/expect)を参照。

# 所感
初見でもわかりやすいようにjestのAPIは整理されていると思う。
ドキュメントも読みやすかった。
思ったよりも簡単にテストが始められたので、JavaScriptのテストを積極的に書いていきたい気持ち。

# 参考
- [jest](https://jestjs.io/docs/ja/getting-started)
- [babel-preset-es2015](https://babeljs.io/docs/en/babel-preset-es2015)
- [github - dooburt/jest-test](https://github.com/dooburt/jest-test)
- [github - LarsBergqvist/jest_playground](https://github.com/LarsBergqvist/jest_playground)

