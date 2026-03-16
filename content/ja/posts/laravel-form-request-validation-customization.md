---
title: Laravelのフォームリクエストでバリデーションされる値をカスタマイズする
description: Laravelのフォームリクエストでバリデーションされる値をカスタマイズする
slug: laravel-form-request-validation-customization
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: laravel-form-request-validation-customization
---


# 概要
Laravelのフォームリクエストで、バリデーションされる値をカスタマイズする方法です。
APIのエンドポイントが`/post/:id/delete`の時に、ルートパラメーターにフォームリクエストのバリデーションをかけたい・・なんて時に有効かもしれません。

# validationDataメソッドをいじる
[Laravel API](https://github.com/laravel/framework/blob/5.2/src/Illuminate/Foundation/Http/FormRequest.php)にある`validationData`をいじります。

ルートパラメーターのidにバリデーションをかける例です。

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

# 所感
結局このカスタマイズは何となく気持ち悪くて使っていません（）

# 参考
+ [How to validate Route Parameters in Laravel 5?](https://laracasts.com/discuss/channels/general-discussion/how-to-validate-route-parameters-in-laravel-5)

