---
title: 'Handling Promises in Server-Side Validation with redux-form: ''Uncaught (in promise) error'''
slug: redux-form-promise-error
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
  - Redux
  - redux-form
  - Tips
translation_key: redux-form-promise-error
---

# Overview
This post discusses an issue encountered while implementing server-side validation with redux-form, where manipulating promises led to an 'Uncaught (in promise) error' when throwing a `SubmissionError`.

# Solution
The issue was simply the absence of a `return` statement.

Before Fix

```
class Categories extends Component {
  onSubmit(props) {
    const {createCategory, fetchCategories, reset} = this.props;

    createCategory(props).then((res) => {
      if (res.error) {
        console.log('error');
        throw new SubmissionError({name: 'User does not exist', _error: 'Login failed!'});
      } else {
        console.log('success');
        reset();
        fetchCategories();
      }
    });
  }

  // Other methods omitted
}
```

After Fix

```
class Categories extends Component {
  onSubmit(props) {
    const {createCategory, fetchCategories, reset} = this.props;

    return createCategory(props).then((res) => {
      if (res.error) {
        console.log('error');
        throw new SubmissionError({name: 'User does not exist', _error: 'Login failed!'});
      } else {
        console.log('success');
        reset();
        fetchCategories();
      }
    });
  }

  // Other methods omitted
}
```

# Thoughts
JavaScript is tough. I still don't fully understand promises (I only recognize them as a way to simplify callbacks...).

If you're building a SPA with Laravel and React, please join [Lara Cafe](https://laracafe.connpass.com/) (Help me!).

# References
[Redux Form - Submit Validation Example](http://redux-form.com/6.0.0-alpha.4/examples/submitValidation/)
[throw new SubmissionError() causing Uncaught (in promise) error](https://github.com/erikras/redux-form/issues/2269)