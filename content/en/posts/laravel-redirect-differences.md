---
title: Differences Between redirect('hoge') and redirect()->to('hoge') in Laravel
description: 'An in-depth look at Differences Between redirect(''hoge'') and redirect()->to(''hoge'') in Laravel, covering key concepts and practical insights.'
slug: laravel-redirect-differences
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-redirect-differences
---

It's a trivial matter, but I was curious, so I looked into it.

```php
public function getIndex()
{
  return redirect()->to('hoge');
}
```

I've been using this one for some reason,

```php
public function getIndex()
{
  return redirect('hoge');
}
```

But both work without any issues, so I investigated the implementation of the redirect helper.

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

It seems that when the argument is empty, it calls the instance. It was also mentioned in the [documentation](https://readouble.com/laravel/5.1/ja/responses.html) lol.

The API of the called instance can be found [here](https://github.com/laravel/framework/blob/5.1/src/Illuminate/Routing/Redirector.php#L113).

The implementation of the to method is as follows:

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

If you just want a simple redirect, use `redirect('hoge')`. If you want to carry flash data or redirect to a controller method, use `redirect()` which returns an instance with an empty argument.

# Note
I thought I should check the implementation of the code I've been unconsciously writing since I started using Laravel φ(..)