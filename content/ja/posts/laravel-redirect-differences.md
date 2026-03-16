---
title: Laravelのredirect('hoge')とredirect()->to('hoge')の違い
description: 'Laravelのredirect(''hoge'')とredirect()->to(''hoge'')の違い'
slug: laravel-redirect-differences
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: laravel-redirect-differences
---


些細な事ですが、気になったので調べてみました。

```php
public function getIndex()
{
  return redirect()->to('hoge');
}
```

今まで何となくこっちを使っていましたが、


```php
public function getIndex()
{
  return redirect('hoge');
}
```

こっちでも問題なく動作するのでredirectヘルパーの実装について調べてみました。


# redirectヘルパーの実装

```Illuminate/Foundation/helpers.php
if (!function_exists('redirect')) {
    /**
     * Get an instance of the redirector.
     *
     * @param  string|null  $to
     * @param  int     $status
     * @param  array   $headers
     * @param  bool    $secure
     * @return \Illuminate\Routing\Redirector|\Illuminate\Http\RedirectResponse
     */
    function redirect($to = null, $status = 302, $headers = [], $secure = null)
    {
        if (is_null($to)) {
            return app('redirect');
        }

        return app('redirect')->to($to, $status, $headers, $secure);
    }
}
```

どうやら引数が空だとインスタンスを呼び出してくれるみたいです。
[ドキュメント](https://readouble.com/laravel/5.1/ja/responses.html)にもそう書いてありましたｗ

呼び出されるインスタンスのapiは[ここ](https://github.com/laravel/framework/blob/5.1/src/Illuminate/Routing/Redirector.php#L113)
　
toメソッドの実装は以下の通り。

```Illuminate/Routing/Redirector.php
/**
     * Create a new redirect response to the given path.
     *
     * @param  string  $path
     * @param  int     $status
     * @param  array   $headers
     * @param  bool    $secure
     * @return \Illuminate\Http\RedirectResponse
     */
    public function to($path, $status = 302, $headers = [], $secure = null)
    {
        $path = $this->generator->to($path, [], $secure);
        return $this->createRedirect($path, $status, $headers);
    }
```

# 結論
`redirect('hoge')`と`redirect()->to('hoge')`は同じ。

単純にリダイレクトだけなら`redirect('hoge')`、フラッシュデータを持たせたり、コントローラーのメソッドにリダイレクトさせたい時などは空でインスタンスを返す`redirect()`を使う。

#　所管
Laravel使い始めた頃から染み付いて無意識にかいているコードの実装は一度くらい確認しようと思ったφ(..)

