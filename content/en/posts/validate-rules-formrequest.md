---
title: Distributing Validation Rules within the FormRequest's rules Method
slug: validate-rules-formrequest
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: validate-rules-formrequest
---



In situations where there are multiple forms, you might think, "I want to consolidate the form request classes into one and branch within the rules method." I had this thought while creating a REST API.

It seems others have thought the same, and I found the solution before diving into the references.

[Multiple Forms, Multiple Requests?](https://laracasts.com/discuss/channels/general-discussion/multiple-forms-multiple-requests)

Let's tinker within the FormRequest's rules method.

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

The parts labeled hogehoge and mogemoge are values passed in the request. (I'm not sure how to best describe this, sorry.)

You can retrieve the request values with something like $request->hoge, but that $request is manipulated by the framework's implementation and becomes $this. (Saying "becomes" might be misleading, but since I haven't looked into the underlying implementation, please excuse my lack of vocabulary...)

The final `return []` is for when null is passed in the request. Without this, it will cause an error when null is encountered.


# Thoughts
Nothing in particular.