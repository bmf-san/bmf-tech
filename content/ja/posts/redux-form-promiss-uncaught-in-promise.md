---
title: "redux-formのサーバーサイドバリデーションでpromissをいじってたら「Uncaught (in promise) error」"
slug: "redux-form-promiss-uncaught-in-promise"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "React"
  - "Redux"
  - "redux-form"
  - "Tips"
draft: false
---

# 概要
redux-formでサーバーサイドのバリデーションを実装している時に、promissをいじってredux-formの`SubmissionError`を投げていたら`Uncaught (in promise) error`と怒られた話です。

# 解決策
`return`がなかっただけでした。

修正前

```
class Categories extends Component {
  onSubmit(props) {
    const {createCategory, fetchCategories, reset} = this.props;

    createCategory(props).then((res) => {
      if (res.error) {
        console.log('error');
        throw new SubmissionError({name: 'User does not exist', _error: 'Login failed!'});
      } else {
        console.log('success');
        reset();
        fetchCategories();
      }
    });
  }

  // 以下色々省略
}
```

修正後

```
class Categories extends Component {
  onSubmit(props) {
    const {createCategory, fetchCategories, reset} = this.props;

    return createCategory(props).then((res) => {
      if (res.error) {
        console.log('error');
        throw new SubmissionError({name: 'User does not exist', _error: 'Login failed!'});
      } else {
        console.log('success');
        reset();
        fetchCategories();
      }
    });
  }

  // 以下色々省略
}
```

# 所感
jsむずい。promissまだ良くわかっていない。（コールバックを楽にするためのもの程度の認識。。。）

Laravel×ReactでSPAつくっているよーという方、ぜひ[Lara Cafe](https://laracafe.connpass.com/)
にご参加ください！（助けてぇ）

# 参考
[Redux Form -Submit Validation Example](http://redux-form.com/6.0.0-alpha.4/examples/submitValidation/)
[throw new SubmissionError() causing Uncaught (in promise) error](https://github.com/erikras/redux-form/issues/2269)

