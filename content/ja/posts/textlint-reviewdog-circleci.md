---
title: "textlintとreviewdogを使ってCircleCIでテキスト校正する"
slug: "textlint-reviewdog-circleci"
date: 2021-10-09
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "CircleCI"
  - "npm"
  - "textlint"
draft: false
---

# 概要
長文の執筆をする際にテキスト校正を自動化しておきたかったのでやってみた。

# 構成
テキストはGithub上で管理するようにしており、ディレクトリ構成は以下のようになっている。

```sh
├── .circleci
│   └── config.yml
├── README.md
├── documents
│   ├── はじめに.md
│   └── おわりに.md
├── images
├── .textlintrc
├── package-lock.json
└── package.json
```

# npmパッケージインストール

初期設定。
```sh
npm init -y
```

textlintとtextlintで使用するルールをインストール。
```sh
npm install --save-dev  textlint  textlint-rule-preset-ja-spacing     textlint-rule-preset-ja-technical-writing     textlint-rule-spellcheck-tech-word textlint-rule-preset-jtf-style textlint-rule-preset-japanese
```

# textlintのルール設定
.textlintrc
```sh
{
  "filters": {},
  "rules": {
    "preset-ja-spacing": true,
    "preset-ja-technical-writing": true,
	"preset-japanese": true,
	"preset-jtf-style": true,
    "spellcheck-tech-word": true
  }
}
```

# CircleCIの設定
Githubで`repo`だけを許可したトークンを発行して、`REVIEWDOG_GITHUB_API_TOKEN`という名前で環境変数をセットしておく。

config.ymlの設定は以下の通り。
```yml
version: 2
jobs:
  build:
    docker:
      - image: vvakame/review:latest
        environment:
          REVIEWDOG_VERSION: latest
    steps:
      - checkout
      - restore_cache:
          keys:
            - npm-cache-{{ checksum "package-lock.json" }}
      - run:
          name: Setup
          command: npm install
      - save_cache:
          key: npm-cache-{{ checksum "package-lock.json" }}
          paths:
            - ./node_modules
      - run:
          name: install reviewdog
          command: "curl -sfL https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh| sh -s $REVIEWDOG_VERSION"
      - run:
          name: lint for ja
          command: "$(npm bin)/textlint -f checkstyle documents/*.md | tee check_result"
      - run:
          name: reviewdog
          command: >
              if [ -n "$REVIEWDOG_GITHUB_API_TOKEN" ]; then
                cat check_result | ./bin/reviewdog -f=checkstyle -name=textlint -reporter=github-pr-review
              fi
          when: on_fail
```

# CIを回してみる
textlintに引っかかるとreviewdogがコメントしてくれる。

![スクリーンショット 2021-10-09 22 38 19](/assets/images/posts/textlint-reviewdog-circleci/136660112-ef84725e-2c4b-4dda-8476-ae30bdb029a2.png)

# 参考
- [gist.github.com - taichi/config.yml](https://gist.github.com/taichi/fef9839243c2fbd017d272c7d838dbde)
- [github.com - reviewdog/reviewdog](https://github.com/reviewdog/reviewdog)
- [qiita.com - textlint と VS Code で始める文章校正](https://qiita.com/takasp/items/22f7f72b691fda30aea2)
- [budougumi0617.github.io - textlint/reviewdogで文書校正エラーをGitHubのプルリクエストにコメントする 2019年6月版
](https://budougumi0617.github.io/2019/06/22/textlint-with-reviewdog-for-review/)
