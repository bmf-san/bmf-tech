---
title: Difference Between redirect('hoge') and redirect()->to('hoge') in Laravel
slug: laravel-redirect-differences
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Applications
tags:
  - Laravel
description: A quick investigation into the differences between two redirect methods in Laravel.
translation_key: laravel-redirect-differences
---

A minor detail, but I was curious, so I decided to look into it.

```php
public function getIndex()
{
  return redirect()->to('hoge');
}
```

I've been using this method without much thought, but...

```php
public function getIndex()
{
  return redirect('hoge');
}
```

This one works just as well, so I decided to investigate the implementation of the `redirect` helper.

# Implementation of the redirect Helper

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

It seems that if the argument is null, it will return an instance. This is also mentioned in the [documentation](https://readouble.com/laravel/5.1/ja/responses.html) lol.

The API of the instance being called can be found [here](https://github.com/laravel/framework/blob/5.1/src/Illuminate/Routing/Redirector.php#L113).

The implementation of the `to` method is as follows:

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

# Conclusion
`redirect('hoge')` and `redirect()->to('hoge')` are the same.

If you just want a simple redirect, use `redirect('hoge')`. If you need to include flash data or redirect to a controller method, use `redirect()` to return an instance.

# Personal Note
Since I started using Laravel, I've been writing code unconsciously. This made me realize that I should take a moment to review the implementation of the code I use regularly. φ(..)