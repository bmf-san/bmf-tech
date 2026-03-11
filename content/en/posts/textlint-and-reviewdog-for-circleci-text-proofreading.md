---
title: Proofreading Text with textlint and reviewdog on CircleCI
slug: textlint-and-reviewdog-for-circleci-text-proofreading
date: 2021-10-09T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - CircleCI
  - npm
  - textlint
description: Automating text proofreading using textlint and reviewdog on CircleCI.
translation_key: textlint-and-reviewdog-for-circleci-text-proofreading
---

# Overview
I wanted to automate text proofreading when writing long documents, so I gave it a try.

# Structure
The text is managed on GitHub, and the directory structure is as follows:

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

# Installing npm Packages

Initial setup.
```sh
npm init -y
```

Install textlint and the rules used by textlint.
```sh
npm install --save-dev  textlint  textlint-rule-preset-ja-spacing     textlint-rule-preset-ja-technical-writing     textlint-rule-spellcheck-tech-word textlint-rule-preset-jtf-style textlint-rule-preset-japanese
```

# Setting textlint Rules
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

# CircleCI Configuration
Generate a token that only allows `repo` on GitHub and set it as an environment variable named `REVIEWDOG_GITHUB_API_TOKEN`.

The config.yml settings are as follows:
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

# Running the CI
If textlint catches any issues, reviewdog will comment on them.

![Screenshot 2021-10-09 22 38 19](/assets/images/posts/textlint-and-reviewdog-for-circleci-text-proofreading/136660112-ef84725e-2c4b-4dda-8476-ae30bdb029a2.png)

# References
- [gist.github.com - taichi/config.yml](https://gist.github.com/taichi/fef9839243c2fbd017d272c7d838dbde)
- [github.com - reviewdog/reviewdog](https://github.com/reviewdog/reviewdog)
- [qiita.com - Starting Text Proofreading with textlint and VS Code](https://qiita.com/takasp/items/22f7f72b691fda30aea2)
- [budougumi0617.github.io - Commenting Document Proofreading Errors on GitHub Pull Requests with textlint/reviewdog June 2019 Edition](https://budougumi0617.github.io/2019/06/22/textlint-with-reviewdog-for-review/)