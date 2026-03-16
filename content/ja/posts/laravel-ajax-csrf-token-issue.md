---
title: Laravelでajaxの際にCsrfTokenに引っかかった話
description: Laravelでajaxの際にCsrfTokenに引っかかった話
slug: laravel-ajax-csrf-token-issue
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - AJAX
  - Laravel
  - React
  - superagent
translation_key: laravel-ajax-csrf-token-issue
---


Laravel+React+SuperagentでAjaxを実装していたら、500エラーがでて、「いやいやまさかTokenじゃないっしょ〜　そんなの知ってるもん〜」と捻くれていたら、CsrfTokenが原因でした。


# 解決方法

ヘッダにTokenを含める方法もありますが、ハードコーディング感があるので、VerifyCsrfToken.phpで該当URLを除外する方法がスマートだと思うのでそちらを記載します。


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
        'api/*'
    ];
}
```

ワイルドカードも使えちゃいます。


# 所感
この手の記事が多々あるのでそれにまんまと引っかかった自分を殴り飛ばしたいです。。。

