---
title: Encountering 'Uncaught (in promise) error' with redux-form Server-side Validation
description: 'A troubleshooting guide for Encountering ''Uncaught (in promise) error'' with redux-form Server-side Validation, explaining the root cause and how to resolve it.'
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
This post discusses an issue encountered when implementing server-side validation with redux-form. While manipulating promises and throwing redux-form's `SubmissionError`, an `Uncaught (in promise) error` occurred.

# Solution
The issue was simply due to a missing `return` statement.

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

  // Other parts omitted
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

  // Other parts omitted
}
```

# Thoughts
JavaScript is tough. I still don't fully understand promises (I only know they make callbacks easier...).

If you're creating an SPA with Laravel and React, please join us at [Lara Cafe](https://laracafe.connpass.com/)! (Help needed)

# References
[Redux Form -Submit Validation Example](http://redux-form.com/6.0.0-alpha.4/examples/submitValidation/)
[throw new SubmissionError() causing Uncaught (in promise) error](https://github.com/erikras/redux-form/issues/2269)
