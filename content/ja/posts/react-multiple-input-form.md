---
title: "Reactで複数のinputを扱うフォームを実装する"
slug: "react-multiple-input-form"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "React"
draft: false
---

Reactでフォームを実装するのはちょっと癖があります。

まずReactのフォームにはcontrolled formとuncontrolled formという２つのパターンがあることを先に理解しておくと良いかもしれません。参考：[React ドキュメント](https://facebook.github.io/react/docs/forms-ja-JP.html)

私はまだ理解が及んでいないところも多々ありますが、Reactの実装例が少ないので少しでも刺激になればという感じで記事を公開したいと思います。
（もっと楽なやり方とかこうした方がいいといった指摘があると幸いです。）

今回の実装にあたり色々ググったのですが、inputが一つしかない実装例ばかりで結構しんどかったです。。。


# やること
* 複数のinputを扱うフォームの実装

# やらないこと
* フォームの送信・・・きちんとデータが取得できているかの確認のみ行います。
* bindを用いた実装・・・bindがよくわかっていないので、、、

# 実装

複数のinputを扱う実装例はbindを使用したものがほとんどだと思います。
ドキュメントに載っている実装例や検索するとでてくる大抵の実装例にはbindを使用したケースが多かったです。

しかし、このbindが今ひとつ理解できない・・・。

ネストしていないシンプルなオブジェクトのstateを用意した例は多々あったのですが、ネストしたオブジェクトを利用した場合、どうすればいいのかよくわかりませんでした（泣）

そこで調べまくって見つけたのがこの記事でした。
[Stack Overflow - Best practice for ReactJS form components](http://stackoverflow.com/questions/26626454/best-practice-for-reactjs-form-components)

というわけで上記を参考にしつつ実装します。


htmlはこんな感じです。

```form.html
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <title>React Multi Input Form</title>
</head>
    <body>

    <div id="multi-input-form"></div>

    <!-- scripts -->
    <script src="path/to/react.js"></script>
    <script src="path/to/react-dom.js"></script>
    <script src="path/to/browser.min.js"></script>

    </body>
</html>
```


そしてReactはこんな感じです。


```form.js
var FormApp = React.createClass({
    getInitialState: function () {
        // ネストされたオブジェクトを用意
        return {
            data: {
                name: '',
                email: ''
            }
        };
    },

    handleChange: function (event) {
        // ネストされたオブジェクトのdataまでアクセスしておく
        var data = this.state.data;

        // eventが発火したname属性名ごとに値を処理
        switch (event.target.name) {
            case 'name':
                data.name = event.target.value;
                break;
            case 'email':
                data.email = event.target.value;
                break;
        }

        // 状態を更新
        this.setState({
            data: data
        });
    },

    handleSubmit: function () {
        console.log(this.state.data.name);
        console.log(this.state.data.email);
    },

    render: function () {
        return (
            <form action="javascript:void(0)" onSubmit={this.handleSubmit}>
                {/* Name */}
                <label htmlFor="name">お名前</label>
                <input type="text" name="name" value={this.state.name} onChange={this.handleChange} />

                {/* Email */}
                <label htmlFor="email">メールアドレス</label>
                <input type="email" name="email" value={this.state.email} onChange={this.handleChange} />

                 {/* Submit Button */}
                <button type="submit">送信</button>
            </form>
        );
    }

});

ReactDOM.render(
    <FormApp />,
    document.getElementById('multi-input-form')
);
```



冒頭で面倒だと言いましたが、案外シンプルな気がしてきました。笑
特筆すべき点は、、、特になさそうですね。


動作はこんな感じです。


![react-form.gif](https://qiita-image-store.s3.amazonaws.com/0/124495/9f1e3bb6-d5ed-eec8-7605-cfdeb54ae4b2.gif)



# 所感

Reactの記事を色々調べていて思ったのですが、Reactの触りしかない記事ともうほんとに前線でReactやってみますみたいな記事の二極化が結構進んでいるように感じました。

Javascriptの前提知識が不足気味な私には生意気な物言いですが、もう少し過程にフォーカスした記事があるといいなーと思いました。

同じReactでも人によってはESなんちゃらだったりstrictモードだったり、ビルド環境が違ってたり、、フロントエンドのぬかるみにはまっている自分には何が何やらです(*_*)

毎度Reactネタは悲壮感漂っている気がしますが、Reactをかくのは楽しいです。(｀･ω･´)ゞ


# 参考
* [Stack Overflow - Best practice for ReactJS form components](http://stackoverflow.com/questions/26626454/best-practice-for-reactjs-form-components)・・・今回１番参考にした記事かもしれません。
* [Stack Overflow - How do I edit multiple input controlled components in React?](http://stackoverflow.com/questions/35965275/how-do-i-edit-multiple-input-controlled-components-in-react)
* [ES2015 以降で React 書くなら form 部品での onChange で setState するのもう全部これでいいんじゃないかなあ](http://bps-tomoya.hateblo.jp/entry/2016/05/25/154401)
* [React.jsでFormを扱う](http://qiita.com/koba04/items/40cc217ab925ef651113)

