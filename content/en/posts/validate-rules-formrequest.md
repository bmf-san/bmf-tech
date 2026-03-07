---
title: Distributing Validation Rules in FormRequest's Rules Method
slug: validate-rules-formrequest
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: validate-rules-formrequest
---

In situations with multiple forms, you might think, "I want to limit the FormRequest class to one and branch within the rules method." (I encountered this while creating a Rest API.)

It seems that others have had the same thought, and I found a solution before digging through the references.

[Multiple Forms, Multiple Requests?](https://laracasts.com/discuss/channels/general-discussion/multiple-forms-multiple-requests)

Let's get into the rules method of FormRequest.

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

		// Default
		return [];
    }
```

The `hogehoge` and `mogemoge` are values passed to the request. (I'm not sure how to describe them...)

You can retrieve request values like `$request->hoge`, but that `$request` is manipulated by the framework's implementation and becomes `$this`. (Saying it has changed might be misleading, but since I haven't looked at the underlying implementation, I apologize for my limited vocabulary...)

The last `return []` is for when null is passed in the request. Without this, it would throw an error when null is encountered.

# Thoughts
I don't have any particular thoughts.