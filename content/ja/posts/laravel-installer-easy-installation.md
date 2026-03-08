---
title: Laravelをinstallerでカンタンインストール
slug: laravel-installer-easy-installation
date: 2016-05-15T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: laravel-installer-easy-installation
---


Laravelをインストールする時、composerを使って毎回面倒なコマンドをうっていたのですが、インストーラを使ったほうが楽でした（今更感）。


# 準備
- composer
- MacOS


# インストーラをグローバルインストール

`composer global require "laravel/installer"`


# パスを通す

 MacOSならこれでいけると思います。（Winは知りません・・・）

`export PATH="~/.composer/vendor/bin:$PATH"`

#  新規プロジェクト作成

`laravel new PROJECTNAME`


最新版のLaravelちゃんがカレントディレクトリにインストールされます。

ドキュメントにもかいてありますが、**composerを使うよりも早く動く**ようです。


# 感想

早くて楽。

# 補足
MacOSのホームディレクトリの.bash_profileに`export PATH="~/.composer/vendor/bin:$PATH"`を書き足しておく。
.bash_profileがない場合はつくる。　.bashrcとの違いが知りたい場合はググる。

