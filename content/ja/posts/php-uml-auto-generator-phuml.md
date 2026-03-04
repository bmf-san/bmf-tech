---
title: "PHPでUMLを自動生成してくれるツールーphUML"
slug: "php-uml-auto-generator-phuml"
date: 2020-06-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "PHP"
draft: false
---

#  概要
クラス設計の外観を把握したい時にUMLを自動生成してくれるツールが欲しかった。
phpstormなら標準でいい感じにdiagramを生成してくれる機能があるらしいが、vscodeに入信してしまったのでいい感じのツールを探すしかない。

# phUML
ぐぐると色々ツールはあるのだが、簡単に使えそうなやつを探してみた。

[github.com - MontealegreLuisphuml](https://github.com/MontealegreLuis/phuml)
[ドキュメント](https://montealegreluis.com/phuml/)

本家？[github.com - jakobwsthoff/phuml](https://github.com/jakobwesthoff/phuml)はメンテ終了しているようなのだが、探してみると上記のfork版のようなやつが見つかった。

スターは少なくてあまり使う人いないのかな・・？という印象だが、ちゃんと使えそうだったので触ってみた。

phpのバージョン対応は`^7.1`。

自分は7.3環境で使ってみた。

## インストール

```sh
$ wget https://montealegreluis.com/phuml/phuml.phar
$ wget https://montealegreluis.com/phuml/phuml.phar.pubkey
$ chmod +x phuml.phar
$ mv phuml.phar /usr/local/bin/phuml
$ mv phuml.phar.pubkey /usr/local/bin/phuml.pubkey
```

composerでインストールすることもできる。

```php
composer require phuml/phuml
```

インストールできたら、

```php
vendor/bin/phuml phuml:diagram -r -a -i -o -p dot path/to/classes example.png
```

こんな感じの怪しげなオプションをいっぱいつけるとクラス図を生成してくれる。

オプションはドキュメントで確認。
[phUML - Generate a class diagram](https://montealegreluis.com/phuml/docs/class-diagram.html)

出力したくないアクセサをオプションで指定したりできるぽい。

# 所感
大きめの設計を把握したい時に概要把握のために重宝しそう。
vscodeのプラグインであったら嬉しいが今の所はなかった。
