---
title: Implementing a Form with Laravel, React, and Superagent
slug: laravel-react-superagent-form
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - React
  - AJAX
  - Superagent
description: A guide to implementing an AJAX form using Laravel, React, and Superagent.
translation_key: laravel-react-superagent-form
---

As the title suggests, this post explains how to implement an AJAX form using Laravel, React, and Superagent.

Superagent is chosen as the AJAX library because I wanted to move away from jQuery, and I found it easier to understand compared to jQuery's AJAX. While there is a complex concept called Promises, you can still use Superagent without fully understanding it for now.

From a web standards perspective, Fetch API seems to be the modern choice, but due to inconsistencies in browser implementations, I decided to avoid it.

Frontend development can feel chaotic at times, but let's move forward.

# What We'll Do
* Set up an API in Laravel
* Convert FormRequest responses to JSON
* Use React to call Laravel's API with Superagent (verify GET and POST requests)

# What We Won't Do
* Setting up the build environment
* Setting up React or Superagent

# Laravel Implementation

The order is a bit random, so bear with me.

## Routing

```route.php
Route::group(['prefix' => 'api/v1'], function () {
  Route::get('/api/user', 'HogeController@index');
  Route::post('/api/user', 'HogeController@update');
});
```

## View

Skipping a lot here. This is just an example of how to summon the component.

```hoge.blade.php
<div id="form-component" class="mdl-cell mdl-cell--12-col"></div>
```

## Controller

In practice, I use a ResourceController to create a RESTful API, but I’ll skip the detailed implementation here.

```HogeController.php
<?php

// NameSpace

class ConfigController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // Prepare an API that returns JSON
        $users = \Auth::user();

        return \Response::json($users);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(ConfigRequest $request)
    {
        // Example update process
        $user_name = \Auth::guard('users')->user()->name;
        $users = User::where('name', $user_name)->first();
        $users->fill(\Input::all())->save();

        // Create an array and return it as JSON
        $response['status'] = 'success';
        $response["message"] = ['No issues with the input!'];

        return \Response::json($response, '200');
    }
}
```

# Request (FormRequest)

```HogeRequest.php
<?php

namespace App\Http\Requests\User;

use App\Http\Requests\Request;

class HogeRequest extends Request
{
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
        if ($this->form_type == 'name') {
          return [
            'user_name' => 'max:5|required',
            'email' => 'email|required'
          ];
        }

        // Default (when null)
        return [];
    }

    public function response(array $errors)
    {
        $response['status'] = 'error';
        $response['message'] = $errors;

        return \Response::json($response, 200);
    }
}
```

To return errors as JSON with FormRequest, you just need to override the `response` method in `Illuminate/Foundation/Http/FormRequest`.

The usage is the same as the usual FormRequest. If there are errors, it will return the error messages as a JSON response.

[Laravel API - FormRequest](https://github.com/laravel/framework/blob/5.2/src/Illuminate/Foundation/Http/FormRequest.php)

## CSRF Token Exception Settings

Don't forget to configure this in `VerifyCsrfToken.php`.

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
        // Wildcards can be used.
        'api/*'
    ];
}
```

This completes the Laravel side.

# React Implementation

```hoge.js

// Strict mode is just for the vibe...
"use strict";

var request = require('superagent');

var ConfigNameForm = React.createClass({
  getInitialState: function () {
    return {
      // Form values
      data: {
        user_name: '',
        email: '',
      },

      // Messages
      message: {
        // If there are no input errors, the controller response is assigned. If there are errors, the FormRequest response is assigned. (If you want to simplify, you might want to skip FormRequest validation and implement validation logic in the controller instead.)
        user_name: '',
        email: ''
      }
    }
  },

  // API-GET
  componentDidMount: function () {
    request
      .get('/api/user')
      .set('Content-Type', 'application/json')
      .end(function(err, res){
        if (err) {
          alert('Communication error. Please reload.');
        }
        this.setState({
          data: {
            user_name: res.body.user_name,
            email: res.body.email
          }
        });
      }.bind(this));
  },

  handleChange: function (event) {
    var data = this.state.data;

    switch(event.target.name) {
      case 'user_name':
        data.user_name = event.target.value;
        break;

      case 'email':
        data.email = event.target.value;
        break;
    }

    this.setState({
      data: data
    });
  },

  // API-POST
  handleSubmit: function () {
    request
      .post('/api/user')
      .set('Content-Type', 'application/json')
      .send({
        user_name: this.state.data.user_name,
        email: this.state.data.email
      })
      .end(function(err, res){
        if (res.ok) {
          var message = this.state.message;

          switch (res.body.status) {
            case 'success':
              // This part is clunky, so adjust as needed.
              message.user_name = res.body.message;
              message.email = res.body.message;
              break;

            case 'error':
              message.user_name = res.body.message.user_name;
              message.email = res.body.message.email;
              break;
          }

          this.setState({
            message: message;
          });
        } else {
          alert('Communication error. Please try again.');
        }
      }.bind(this));
  },

  render: function () {
    // Clunky...
    var msgOfName = false;

    if (this.state.message.name.length > 0) {
      var msgOfName = this.state.message.name.map(function (msg) {
        return (
          <p key={msg}>{msg}</p>
        );
      });
    }

    var msgOfEmail = false;
    if (this.state.message.email.length > 0) {
      var msgOfEmail = this.state.message.email.map(function (msg) {
        return (
          <p key={msg}>{msg}</p>
        );
      });
    }

    return (
      <div>
        {/* Message */}
        {msgOfName}
        {msgOfEmail}

        {/* Form */}
        <form action="javascript:void(0)" method="POST" onSubmit={this.handleSubmit}>
          {/* Name */}
          <label htmlFor="user_name">Name</label>
          <input type="text" name="user_name" id="user_name" value={this.state.data.user_name} onChange={this.handleChange} disabled />

          {/* Email */}
          <label htmlFor="email">Email</label>
          <input type="text" name="email" id="email" value={this.state.data.email} onChange={this.handleChange} />
          <button type="submit">Update</button>
        </form>
      </div>
    );
  }
});

ReactDOM.render(
  <FormApp />,
  document.getElementById('form-component')
);

```

# Thoughts
I made this quite roughly, so there are likely many areas that need improvement.

While architecture is important, I realized I need to study modern JavaScript practices more to write code more flexibly.

# References
* [Things I looked into when creating an API with Laravel 5.1.x](http://qiita.com/zaburo/items/f0db54bd3ebd81a8ce68)
* [About the subtle but important `key` in React.js](http://qiita.com/koba04/items/a4d23245d246c53cd49d) - This was important for iterating over responses from the API.