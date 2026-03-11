---
title: React Tag Autocompleteを使ってサジェスト付きタグ機能を実装する
slug: react-tag-autocomplete-implementation
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - React
translation_key: react-tag-autocomplete-implementation
---


タグ機能をフルスクラッチで開発しようと思っていたのですが、便利なReact Componentが沢山見つかったのでそちらを利用してみました。

# 環境
* React
* Babel
* Browserify
* [React Tag Autocomplete](https://github.com/i-like-robots/react-tags)
* npm

# 導入
npmでReact Tag Autocompleteを導入します。

`npm install --save react-tag-autocomplete`


インクルードの仕方は色々あるかと思いますが、今回の環境ではrequireを使います。

```hoge.js
var ReactTags = require('react-tag-autocomplete');
```

これで準備OKです。

# 実装

```hoge.html
// 色々省略
<div id="react-tag-autocomplete"></div>
```

[github](https://github.com/i-like-robots/react-tags)にUsageがありますが、ちょっと加工してapiをたたいてデータを取得してきたケースを想定してみます。（ここではsuperagentを使っています。）

apiではこんな感じのJsonレスポンスを返しています。

```
[{"id":1,"name":"プログラミング"},{"id":2,"name":"家事"},{"id":3,"name":"自宅警備"},{"id":4,"name":"早寝早起き"},{"id":5,"name":"三日坊主"}]
```

res.body.skillsのデバッグ
![スクリーンショット 2016-09-28 3.04.10.png](/assets/images/posts/react-tag-autocomplete-implementation/819ec751-a2d9-df73-d353-76847710c4b7.png)


```hoge.js
var ReactTags = require('react-tag-autocomplete');

var App = React.createClass({
  getInitialState: function () {
    return {
      tags: [],
      suggestions: []
    }
  },

  componentDidMount: function () {
    request
      .get('/api/v1/user/config')
      .end(function(err, res){
        if (err) {
          alert('通信エラーです。リロードしてください。');
        }
        this.setState({
          suggestions: res.body.skills
        });
      }.bind(this));
  },

  handleDelete: function (i) {
    var tags = this.state.tags
    tags.splice(i, 1)
    this.setState({ tags: tags })
  },![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/173c6de9-b87a-6200-65ed-506e181f565e.gif)
![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/a3372702-2a85-9b80-0b53-ede2c9c3c486.gif)


  handleAddition: function (tag) {
    var tags = this.state.tags
    tags.push(tag)
    this.setState({ tags: tags })
  },

  render: function () {
    return (
      <ReactTags
        tags={this.state.tags}
        suggestions={this.state.suggestions}
        handleDelete={this.handleDelete}
        handleAddition={this.handleAddition} />
    )
  }
})

ReactDOM.render(
  <App />,
  document.getElementById('react-tag-autocomplete')
);
```

動作確認（余計なものが映っていますが・・）
![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/c0c75e77-d212-4cd4-5e1b-c59a8a412c61.gif)

cssは設定していないのですごくダサいですねｗ

他のオプションやcssのクラス名など[github](https://github.com/i-like-robots/react-tags)に丁寧に明記してあります。

# 所感
便利な時代だなぁヽ(´ー｀)ノ

