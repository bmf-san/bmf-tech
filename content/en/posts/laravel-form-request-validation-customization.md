---
title: Customizing Values Validated by Laravel Form Requests
slug: laravel-form-request-validation-customization
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: Learn how to customize the values validated by Laravel form requests.
translation_key: laravel-form-request-validation-customization
---



# Overview
This post explains how to customize the values validated by Laravel form requests. It might be useful when you want to apply validation to route parameters, such as when the API endpoint is `/post/:id/delete`.

# Modifying the validationData Method
We will modify the `validationData` method found in the [Laravel API](https://github.com/laravel/framework/blob/5.2/src/Illuminate/Foundation/Http/FormRequest.php).

Here's an example of applying validation to the route parameter `id`.

```php
<?php

namespace App\Http\Requests\Api\v1\Category;

use Illuminate\Foundation\Http\FormRequest;

class DeleteCategoryRequest extends FormRequest
{
    const NOT_FOUND_CODE = 400;

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
        return [
            'id' => 'numeric',
        ];
    }

    /**
     * Get data to be validated from the request.
     *
     * @return array
     */
    protected function validationData()
    {
        return array_merge($this->request->all(), [
            'id' => $this->route('id'),
        ]);
    }

    /**
     * Get the error messages for the defined validation rules.
     *
     * @return array
     */
    public function messages()
    {
        return [
            'id.numeric' => 'This id is not number',
        ];
    }

    /**
     * Get the proper failed validation response for the request.
     *
     * @param array $errors
     *
     * @return \Symfony\Component\HttpFoundation\Response
     */
    public function response(array $errors)
    {
        $response['messages'] = $errors;

        return response()->json($response, (int) self::NOT_FOUND_CODE);
    }
}
```

# Thoughts
In the end, I haven't used this customization because it feels somewhat awkward.

# References
+ [How to validate Route Parameters in Laravel 5?](https://laracasts.com/discuss/channels/general-discussion/how-to-validate-route-parameters-in-laravel-5)
