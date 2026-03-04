---
title: "FormRequestのrulesメソッド内でバリデーションルールを振り分ける"
slug: "validate-rules-formrequest"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
draft: false
---

複数のフォームがある状況で、「フォームリクエストのクラスは一つに絞って、rulesメソッド内で分岐したい」なんて思う日があるかもしれません。（私はRest APIつくっているときにありました。）

案外同じことを考えている人がいたようで、リファレンス漁るよりも先に結果が出ました。

[Multiple Forms, Multiple Requests?](https://laracasts.com/discuss/channels/general-discussion/multiple-forms-multiple-requests)


FormRequestのrulesメソッド内でゴニョゴニョします。


```HogeRequest.php
/**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
  		if ($this->hogehoge) {
  			return [
  				'alias_name' => 'max:50|required|unique:users',
  			];
  		}

  		if ($this->mogemoge) {
  			return [
  				'self_introduction' => 'max:200'
  			];
  		}

  		// デフォルト
  		return [];
    }
```

hogehoge、mogemogeのところはそれぞれリクエストに渡される値です。（なんといえばいいのでしょうか汗）

$request->hogeって感じでリクエストの値を取得できますが、その$requestがフレームワークの実装でごにょごにょされて、$thisに代わったといった感じです。（代わったというのはかなり語弊がある気がしますが、裏側の実装を見ていないのでボキャ貧でごめんなさいということにしてください。。。）

１番最後の`return []`はnullがリクエストで渡ってきた時のためです。
これがないとnullの時エラーになります。


# 所感
特にありませぬ。

