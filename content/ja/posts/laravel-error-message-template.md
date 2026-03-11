---
title: Laravelでエラーメッセージを共通のテンプレートで対応する方法
slug: laravel-error-message-template
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: laravel-error-message-template
---


# 概要
Laravelのエラーページを共通のテンプレートで対応する方法について説明します。

# エラーハンドリングをカスタマイズ
`app/Exceptions/Handler.php`で`renderHttpException`メソッドをオーバーライドします。

- [Laravel API - Handler.php](https://github.com/laravel/framework/blob/5.4/src/Illuminate/Foundation/Exceptions/Handler.php)

```php:Handler.php
<?php

namespace App\Exceptions;

use Exception;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;

class Handler extends ExceptionHandler
{
    /**
     * A list of the exception types that should not be reported.
     *
     * @var array
     */
    protected $dontReport = [
        \Illuminate\Auth\AuthenticationException::class,
        \Illuminate\Auth\Access\AuthorizationException::class,
        \Symfony\Component\HttpKernel\Exception\HttpException::class,
        \Illuminate\Database\Eloquent\ModelNotFoundException::class,
        \Illuminate\Session\TokenMismatchException::class,
        \Illuminate\Validation\ValidationException::class,
    ];

    /**
     * Report or log an exception.
     *
     * This is a great spot to send exceptions to Sentry, Bugsnag, etc.
     *
     * @param  \Exception  $exception
     * @return void
     */
    public function report(Exception $exception)
    {
        parent::report($exception);
    }

    /**
     * Render an exception into an HTTP response.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Exception  $exception
     * @return \Illuminate\Http\Response
     */
    public function render($request, Exception $exception)
    {
        return parent::render($request, $exception);
    }

    /**
     * Override default method - render the given HttpException.
     *
     * @param  \Symfony\Component\HttpKernel\Exception\HttpException  $e
     * @return \Symfony\Component\HttpFoundation\Response
     */
    protected function renderHttpException(\Symfony\Component\HttpKernel\Exception\HttpException $e)
    {
        $status = $e->getStatusCode();
        $errorMessages = $this->handleErrorMessages($status);

        view()->replaceNamespace('errors', [
            resource_path('views/errors'),
            __DIR__.'/views',
        ]);
        if (view()->exists("errors.index")) {
            return response()->view("errors.index", ['errorMessages' => $errorMessages], $status); // viewを指定する
        } else {
            return $this->convertExceptionToResponse($e);
        }
    }

    /**
     * Handle error messages.
     *
     * @param  int $status
     * @return array $errorMessages
     */
    private function handleErrorMessages($status)
    {
        $errorMessages['status'] = $status;

        switch ($status) {
            case '401':
                return $errorMessages['message'] = 'Unauthorized';
                break;

            case '403':
                return $errorMessages['message'] = 'forbidden';
                break;

            case '404':
                $errorMessages['message'] = 'Not Found';
                break;

            case '500':
                $errorMessages['message'] = 'Internal Server Error';
                break;

            case '503':
                $errorMessages['message'] = 'Service Unavailable';
                break;
        }

        return $errorMessages;
    }
}
```

もっとキレイな書き方がある気がしますが、とりあえずこれで。

後は任意のviewファイルで変数を受け取って出力するだけです。

# 参考
- [Laravel5: エラーページを共通化〜どんなステータスコードでもどんと来い！](http://qiita.com/M_Ishikawa/items/1f0d72fc93286109464e)

