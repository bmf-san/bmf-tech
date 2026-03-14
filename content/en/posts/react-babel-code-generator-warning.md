---
title: 'Encountering ''[BABEL] Note: The code generator has deoptimised the styling of'' in React'
description: "Fix Babel's code generator deoptimization warnings in React through strategic file size optimization and compression techniques."
slug: react-babel-code-generator-warning
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
  - Tips
translation_key: react-babel-code-generator-warning
---


ReactでComponentをrequireしてたら見慣れないエラーがでました。

# Environment
* Laravel
* Elixir
* Babel
* Browserify


# The File is Too Large!

`[BABEL] Note:The code generator has deoptimised the styling of "D:/path/to/hoge.js" as it exceeds the max of "100KB"`

This error occurs when using Babel and the file size is too large.

[What does “The code generator has deoptimised the styling of [some file] as it exceeds the max of ”100KB“” mean?](http://stackoverflow.com/questions/29576341/what-does-the-code-generator-has-deoptimised-the-styling-of-some-file-as-it-e)

It seems there's nothing much to worry about, but if you want to hide the warning, you can set the `compact` option in Babel to `false`.

[How to edit Babel plugins when using Laravel-Elixir's browserify](http://qiita.com/fagai/items/c4dbe5d2adeb79e42e40)

This is about editing plugins, but it might be useful for options as well.

However, even if you hide the warning, the file remains large, so I addressed it by compressing the file.

`gulp --production`

I wasn't quite sure how to change the options ヽ(´ー｀)ノ
