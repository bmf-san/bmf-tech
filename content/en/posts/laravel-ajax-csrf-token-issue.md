---
title: Encountering CsrfToken Issues with AJAX in Laravel
description: 'Resolve Laravel CSRF token issues in AJAX requests by configuring VerifyCsrfToken middleware to exclude API routes.'
slug: laravel-ajax-csrf-token-issue
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - AJAX
  - Laravel
  - React
  - superagent
translation_key: laravel-ajax-csrf-token-issue
---



When implementing AJAX with Laravel, React, and Superagent, I encountered a 500 error. Initially, I thought, "No way it's the Token, I know better than that," but it turned out the CsrfToken was indeed the cause.


# Solution

While you can include the Token in the header, which feels a bit hard-coded, I believe a smarter approach is to exclude the relevant URL in `VerifyCsrfToken.php`. Here's how you can do it:


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

You can even use wildcards.


# Thoughts
There are many articles like this, and I feel like kicking myself for falling into the same trap...