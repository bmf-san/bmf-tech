---
title: Implementing Forms with Laravel, React, and Superagent
description: A step-by-step guide on Implementing Forms with Laravel, React, and Superagent, with practical examples and configuration tips.
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
translation_key: laravel-react-superagent-form
---

As the title suggests, we will implement an Ajax form using Laravel, React, and Superagent.

The reason for choosing Superagent as the Ajax library is that I wanted to move away from jQuery, and I found it easier to understand than jQuery's Ajax. There seems to be a complicated concept called Promises, but we can set that aside for now and see if we can use it.

From a web standards perspective, the Fetch API is considered cool, but it seems to have inconsistencies in implementation across different browser vendors, so I decided to avoid it.

While I grumble about how chaotic the frontend can be, I would like to proceed with the discussion.

# What to Do
* Prepare an API with Laravel
* Set the FormRequest response to JSON
* Use Superagent to call Laravel's API with React → Check GET and POST

# What Not to Do
* Set up the build environment
* Set up React and Superagent

# Laravel Implementation

Please excuse the random order...

## Routing

```route.php
Route::group(['prefix' => 'api/v1'], function () {
  Route::get('/api/user', 'HogeController@index');
  Route::post('/api/user', 'HogeController@update');
});
```

## View

I will omit various details. This is just a demonstration of how to summon components.

```hoge.blade.php
<div id="form-component" class="mdl-cell mdl-cell--12-col"></div>
```

## Controller

In reality, I create the API with a ResourceController and structure it in a RESTful manner, but I will skip the detailed construction.

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
        // Example of update processing
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

    //
    public function response(array $errors)
    {
        $response['status'] = 'error';
        $response['message'] = $errors;

        return \Response::json($response, 200);
    }
}
```

To return errors as JSON with FormRequest, you just need to override the response method of Illuminate/Foundation/Http/FormRequest.

After that, the usage is the same as the usual FormRequest. If there are errors, it will return the error messages as a JsonResponse.

[Laravel API - FormRequest](https://github.com/laravel/framework/blob/5.2/src/Illuminate/Foundation/Http/FormRequest.php)

## CSRF Token Exception Settings
Don't forget to set this in VerifyCsrfToken.php.

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

That's it for the Laravel side.

# React Implementation

```hoge.js

// Strict mode is just for show...
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
        // If there are no input errors, the controller's response will be assigned; if there are, the form request's response will be assigned. (If you want to simplify, it might be better to stop the form request validation and build the validation logic in the controller.)
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
              // This is a bit clunky, but please adjust as needed.
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
    // A bit clunky...
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
          <label htmlFor="email">Email Address</label>
          <input type="text" name="email" id="email" value={this.state.data.email} onChange={this.handleChange} />
          <button type="submit">Update</button>
        </form>
      </div>
    );
  }
});

ReactDOM.render(
  <FormApp />, 
  document.getElementById('form-compoent')
);
```

# Thoughts
I made it quite roughly, so there seems to be a lot that needs fixing.

While architecture is important, I think I need to study modern JavaScript writing more to be able to write flexibly.

# References
* [Things I was curious about when creating an API with Laravel 5.1.x](http://qiita.com/zaburo/items/f0db54bd3ebd81a8ce68)
* [The important but subtle key in React.js](http://qiita.com/koba04/items/a4d23245d246c53cd49d) ... It was important for scanning the response returned from the API.