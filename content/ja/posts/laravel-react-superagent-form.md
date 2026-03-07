---
title: Laravel+React+Superagentでフォームを実装
slug: laravel-react-superagent-form
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
  - React
  - AJAX
  - superagent
translation_key: laravel-react-superagent-form
---


表題の通りです。
Laravel+React+SuperagentでAjaxなフォームを実装します。

AjaxライブラリとしてSuperagentを採用しているのは、jQueryから脱却したいのと、jQueryのAjaxよりも分かりやすかったからです。
プロミスとかいう難しい概念があるらしいですが、それは横に置いておいてもとりあえずは使えそうです。

Web標準の観点からするとFetchAPIがイケてるらしいのですが、各ブラウザベンダーの実装にばらつきがあるようなので避けました。

フロントエンドってつくづくカオスだなーとボヤキつつも話を進めていきたいと想います。


# やること
* LaravelでAPIを用意する
* FormRequestのレスポンスをJSONにする
* ReactでLaravelのAPIをSuperagentでたたく→getとpostの確認

# やらないこと
* ビルド環境のセットアップ
* reactやsuperagentのセットアップ


# Laravel側の実装　

順番テキトーですが、ご了承ください。。。

## ルーティング

```route.php
Route::group(['prefix' => 'api/v1'], function () {
  Route::get('/api/user', 'HogeController@index');
  Route::post('/api/user', 'HogeController@update');
});
```

## ビュー

色々省略しちゃいます。
こんな感じでコンポーネント召喚しますよーという体だけです。

```hoge.blade.php
<div id="form-component" class="mdl-cell mdl-cell--12-col"></div>
```

## コントローラー

実際はResouceControllerでAPIつくって、Restな感じに仕立てているのですが、詳しい作り方は省きます。

```HogeController.php
<?php

// NameSpace

class ConfigController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // Jsonを返すAPIを用意
        $users = \Auth::user();

        return \Response::json($users);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(ConfigRequest $request)
    {
        // Update処理例
        $user_name = \Auth::guard('users')->user()->name;
        $users = User::where('name', $user_name)->first();
        $users->fill(\Input::all())->save();


        // 配列つくってJsonにポイー
        $response['status'] = 'success';
        $response["message"] = ['入力に問題ありません！'];

        return \Response::json($response, '200');
    }
}

```
# Request(FormRequest)

```HogeRequest.php
<?php

namespace App\Http\Requests\User;

use App\Http\Requests\Request;

class HogeRequest extends Request
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
　　　　　　　　　　　　　　　　if ($this->form_type == 'name') {
          return [
            'user_name' => 'max:5|required',
            'email' => 'email|required'
          ];
        }

        // デフォルト（nullの時）
        return [];
    }

    //
    public function response(array $errors)
    {
		$response['status'] = 'error';
        $response['message'] = $errors;

        return \Response::json($response, 200);
    }
}
```

FormReqeustでエラーをJsonで返す方法ですが、Illuminate/Foundation/Http/FormRequestのresponseメソッドをオーバーライドしてあげるだけです。

それで使い方はいつものFormRequestと同じです。
エラーがあればJsonResponseでエラーメッセージを返してくれます。


[Laravel API - FormRequest](https://github.com/laravel/framework/blob/5.2/src/Illuminate/Foundation/Http/FormRequest.php)


## CSRF Tokenの例外設定
VerifyCsrfToken.phpで設定を忘れずに済ませておきます。

```VerifyCsrfToken.php
<?php

namespace App\Http\Middleware;

use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken as BaseVerifier;

class VerifyCsrfToken extends BaseVerifier
{
    /**
     * The URIs that should be excluded from CSRF verification.
     *
     * @var array
     */
    protected $except = [
        // ワイルドカード使えます。
        'api/*'
    ];
}
```


Laravel側はこれで完了です。


# React側の実装

```hoge.js

// strictモードは雰囲気だけです。。。
"use strict";

var request = require('superagent');

var ConfigNameForm = React.createClass({
  getInitialState: function () {
    return {
      // フォームの値
      data: {
        user_name: '',
        email: '',
      },

      // メッセージ
      message: {
      　// 入力エラーがない場合はコントローラーのレスポンスが、ある場合はフォームリクエストのレスポンスがそれぞれ代入されます。（もし単純にしたい場合はフォームリクエストのバリデーションをやめてコントローラーでバリデーションロジックを組んだ方がいいかもしれないです。）
        user_name: '',
        email: ''
      }
    }
  },

  // API-GET
  componentDidMount: function () {
    request
      .get('/api/user')
      .set('Content-Type', 'application/json')
      .end(function(err, res){
        if (err) {
          alert('通信エラーです。リロードしてください。');
        }
        this.setState({
          data: {
            user_name: res.body.user_name,
            email: res.body.email
          }
        });
      }.bind(this));
  },

  handleChange: function (event) {
    var data = this.state.data;

    switch(event.target.name) {
      case 'user_name':
        data.user_name = event.target.value;
        break;

      case 'email':
        data.email = event.target.value;
        break;
    }

    this.setState({
      data: data
    });
  },

  // API-POST
  handleSubmit: function () {
    request
      .post('/api/user')
      .set('Content-Type', 'application/json')
      .send({
        user_name: this.state.data.user_name,
        email: this.state.data.email
      })
      .end(function(err, res){
        if (res.ok) {
          var message = this.state.message;

          switch (res.body.status) {
            case 'success':
              // ここは野暮ったいですが、適当に調整してください。
              message.user_name = res.body.message;
              message.email = res.body.message;
              break;

            case 'error':
              message.user_name = res.body.message.user_name;
              message.email = res.body.message.email;
              break;
          }

          this.setState({
            message: message;
          });
        } else {
          alert('通信エラーです。もう一度お試しください。')
        }
      }.bind(this));
  },

  render: function () {
    // 野暮ったい。。.
    var msgOfName = false;

    if (this.state.message.name.length > 0) {
      var msgOfName = this.state.message.name.map(function (msg) {
        return (
          <p key={msg}>{msg}</p>
        );
      });
    }

    var msgOfEmail = false;
    if (this.state.message.email.length > 0) {
      var msgOfEmail = this.state.message.email.map(function (msg) {
        return (
          <p key={msg}>{msg}</p>
        );
      });
    }

    return (
      <div>
        {/* Message */}
        {msgOfName}
        {msgOfEmail}

        {/* Form */}
        <form action="javascript:void(0)" method="POST" onSubmit={this.handleSubmit}>
          {/* Name */}
          <label htmlFor="user_name">名前</label>
          <input type="text" name="user_name" id="user_name" value={this.state.data.user_name} onChange={this.handleChange} disabled />

          {/* Email */}
          <label htmlFor="email">メールアドレス　</label>
          <input type="text" name="email" id="email" value={this.state.data.email} onChange={this.handleChange} />
          <button type="submit">更新</button>
        </form>
      </div>
    );
  }
});

ReactDOM.render(
  <FormApp />,
  document.getElementById('form-compoent')
);

```
# 所感
結構雑につくったので手直しすべきところは多そうです。

アーキテクチャも大事ですが、モダンなJavaScriptの書き方はもっと勉強して柔軟にかけるようにすべきだと思いました。



# 参考
* [Laravel5.1.xでAPIを作る際に気になっていたことを調べました](http://qiita.com/zaburo/items/f0db54bd3ebd81a8ce68)
* [React.jsの地味だけど重要なkeyについて](http://qiita.com/koba04/items/a4d23245d246c53cd49d)・・・API叩いて返ってきたレスポンスを走査するのに重要でした。

