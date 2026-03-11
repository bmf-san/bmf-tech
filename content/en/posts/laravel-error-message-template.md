---
title: How to Use a Common Template for Error Messages in Laravel
slug: laravel-error-message-template
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: Explains how to handle Laravel error pages with a common template.
translation_key: laravel-error-message-template
---



# Overview
This post explains how to handle Laravel error pages using a common template.

# Customizing Error Handling
Override the `renderHttpException` method in `app/Exceptions/Handler.php`.

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
            return response()->view("errors.index", ['errorMessages' => $errorMessages], $status); // specify the view
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

I feel there might be a cleaner way to write this, but for now, this will do.

After that, just receive the variables in any view file and output them.

# References
- [Laravel5: Unifying Error Pages - Bring on Any Status Code!](http://qiita.com/M_Ishikawa/items/1f0d72fc93286109464e)
