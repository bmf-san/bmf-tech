---
title: Getting Started with JavaScript Testing Using Jest
slug: jest-javascript-testing-introduction
date: 2018-09-20T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - ES5
  - JavaScript
  - babel
  - babel-jest
  - ESModules
  - jest
description: An introduction to writing JavaScript tests with Jest.
translation_key: jest-javascript-testing-introduction
---

# Overview
Write JavaScript tests using Jest.

# Preparation
Since we want to use jest and ESModules, install `babel-preset-2015`.
(Note: `babel-jest` is included with jest.)

`npm install --save-dev jest babel-preset-2015`

The contents of `.babelrc` should look like this:

```
{
  "presets": ["es2015"]
}
```

The `package.json` file should look like this:

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

The directory structure should look like this:
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

There are two patterns for creating test files:
- Files located under a directory named `__tests__` are considered test files.
- Files with extensions `*.spec.js` or `*.test.js` are considered test files.

In this case, we will use the latter format and place the test files in the `test` directory.

# Writing Tests for Native JavaScript
Implement functions for addition and subtraction.

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

Test whether each function returns the correct calculation result.

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

`describe(name, fn)` creates a block to group multiple tests into a test suite.

Run the test:
`npm test ./test/native/calc.test.js`

Test results:
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

# Writing Tests for JavaScript Using ESModules
Create a class that implements methods for addition and subtraction.

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

Test whether each method returns the correct calculation result.

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

# Matchers
Refer to [Jest - Expect](https://jestjs.io/docs/expect).

# Thoughts
I think Jest's API is well-organized and easy to understand, even for first-time users. The documentation was also easy to read. It was easier than expected to get started with testing, which makes me want to actively write JavaScript tests.

# References
- [jest](https://jestjs.io/docs/getting-started)
- [babel-preset-es2015](https://babeljs.io/docs/en/babel-preset-es2015)
- [github - dooburt/jest-test](https://github.com/dooburt/jest-test)
- [github - LarsBergqvist/jest_playground](https://github.com/LarsBergqvist/jest_playground)
