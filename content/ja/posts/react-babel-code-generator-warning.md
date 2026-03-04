---
title: "Reactでrequireしてたら、[BABEL] Note: The code generator has deoptimised the styling of "
slug: "react-babel-code-generator-warning"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "React"
  - "Tips"
draft: false
---

ReactでComponentをrequireしてたら見慣れないエラーがでました。

# 環境
* Laravel
* Elixir
* Babel
* Browserify


# ファイルでかいぞ！

`[BABEL] Note:The code generator has deoptimised the styling of "D:/path/to/hoge.js" as it exceeds the max of "100KB"`

というエラーなわけですが、babelを使っていてファイルが大きすぎると発生するエラーのようです。

[What does “The code generator has deoptimised the styling of [some file] as it exceeds the max of ”100KB“” mean?](http://stackoverflow.com/questions/29576341/what-does-the-code-generator-has-deoptimised-the-styling-of-some-file-as-it-e)

特に気にすることはないみたいですが、警告を非表示にしたいのであれば、babelのcompactというoptionをfalseにすればいいみたいです。

[Laravel-Elixirのbrowserifyを使った時のbabelのpluginsの編集方法](http://qiita.com/fagai/items/c4dbe5d2adeb79e42e40)

これはpluginsの編集方法についてですが、optionに関しても参考になるかと思います。

しかし警告を消してしてもファイルは大きいままなので、ファイルを圧縮させることで対応しました。

`gulp --production`


optionの変更の仕方がよくわからなかったですヽ(´ー｀)ノ

